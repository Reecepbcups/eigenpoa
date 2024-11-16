package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	poakeeper "github.com/strangelove-ventures/poa/keeper"

	apiv1 "github.com/reecepbcups/eigenpoa/api/avs/v1"
	"github.com/reecepbcups/eigenpoa/x/avs/keeper/manager"
	"github.com/reecepbcups/eigenpoa/x/avs/types"
)

type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	Eth *eth.InstrumentedClient

	poaKeeper *poakeeper.Keeper

	m *manager.Manager // the PoA <> AVS manager

	pendingApplyChanges collections.Map[string, uint64] // sdk.ValAddress -> uint64

	// state management
	Schema   collections.Schema
	Params   collections.Item[types.Params]
	OrmDB    apiv1.StateStore
	valStore baseapp.ValidatorStore

	authority string
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	valStore baseapp.ValidatorStore,
	poaKeeper *poakeeper.Keeper,
	logger log.Logger,
	authority string,
) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)

	sb := collections.NewSchemaBuilder(storeService)

	if authority == "" {
		authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	}

	db, err := ormdb.NewModuleDB(&types.ORMModuleSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := apiv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	// Setup ETH client
	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("ethHttp", reg) // TODO: POA?
	eth, err := eth.NewInstrumentedClient("http://127.0.0.1:8545", rpcCallsCollector)
	if err != nil {
		panic(err)
	}

	// TODO: I have to use the incredible-squaring-avs/ repo only i think? (probably fine with how I have it now though)
	// l, err := logging.NewZapLogger(logging.Development)
	// if err != nil {
	// 	panic(err)
	// }
	// ethAvs, err := chainio.BuildAvsReader(common.HexToAddress(c.AVSRegistryCoordinatorAddress), common.HexToAddress(c.OperatorStateRetrieverAddress), eth, l)
	// if err != nil {
	// panic(err)
	// }

	k := Keeper{
		cdc:      cdc,
		logger:   logger,
		valStore: valStore,

		poaKeeper: poaKeeper,

		Eth: eth,

		Params:              collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:               store,
		pendingApplyChanges: collections.NewMap(sb, types.PendingApplyChangesKey, "pending_apply", collections.StringKey, collections.Uint64Value),

		authority: authority,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	// TODO: set via params
	k.m, err = manager.NewManager(common.HexToAddress("0xf5059a5d33d5853360d16c683c16e67980206f36"), k.Eth)
	if err != nil {
		fmt.Printf("Error creating manager %v\n", err)
		panic(err)
	}

	return k
}

func (k Keeper) Logger() log.Logger {
	return k.logger
}

func (k Keeper) GetPOAKeeper() *poakeeper.Keeper {
	return k.poaKeeper
}

// InitGenesis initializes the module's state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {

	if err := data.Params.Validate(); err != nil {
		return err
	}

	return k.Params.Set(ctx, data.Params)
}

// ExportGenesis exports the module's state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params: params,
	}
}

func (k *Keeper) EthOperatorToCosmosValoper(ctx context.Context, e common.Address) (sdk.ValAddress, error) {
	opts := &bind.CallOpts{Context: context.Background()}

	valoper, err := k.m.GetOperatorCosmosAddress(opts, e)
	if err != nil {
		fmt.Printf("Error fetching operator cosmos address %v\n", err)
		return nil, fmt.Errorf("error fetching operator cosmos address: %w", err)
	}

	addr, err := sdk.ValAddressFromBech32(valoper) // TODO: use addresscodec instead
	if err != nil {
		fmt.Printf("Error converting valoper to bytes %v\n", err)
		return nil, fmt.Errorf("error converting valoper to bytes: %w", err)
	}

	return addr, nil
}

// Gets ETH Operators validator addresses mapped to the SDK
func (k *Keeper) GetOperators(ctx context.Context, ethBlockHeight uint64) ([]sdk.ValAddress, error) {
	opts := &bind.CallOpts{Context: ctx}

	// TODO: set in state params (if none, ignore and log)
	// TODO: is really the HelloWorldServiceManager for now

	ethOperators, err := k.m.GetOperators(opts)
	if err != nil {
		fmt.Printf("Error fetching ETH operators %v\n", err)
		return nil, fmt.Errorf("error fetching ETH operators: %w", err)
	}

	operators := make([]sdk.ValAddress, 0, len(ethOperators))
	for _, eOp := range ethOperators {
		valoper, err := k.EthOperatorToCosmosValoper(ctx, eOp)
		if err != nil {
			fmt.Printf("Error fetching operator cosmos address %v\n", err)
			return nil, fmt.Errorf("error fetching operator cosmos address: %w", err)
		}

		fmt.Printf(" - ETH: %v, Cosmos: %s\n", eOp, valoper.String())

		operators = append(operators, valoper)
	}
	return operators, nil
}

// This request does not match the sdk's checked PreBlocker interface (why??) so we have to manually impl this in the app PreBlocker.
func (k *Keeper) PreBlock(ctx context.Context, req *abci.RequestFinalizeBlock) error {
	if req == nil {
		return nil
	}

	injectedData := getInjectedData(req.Txs)
	if injectedData != nil {
		if len(injectedData.Operators) == 0 {
			fmt.Println("PreBlock no operators found, skipping")
			return nil
		}

		fmt.Println("PreBlock injectedData", injectedData) // TODO:

		for _, valoper := range injectedData.Operators { // TODO: are these eth or valoper
			fmt.Printf("PreBlock injectedData operator %v\n", valoper.String())

			// TODO: set the power on the next beginblock, the Comet/SDK does not like it here
			// TODO: this may have actually just been a npe, try moving that logic back here and see what happens (from the beginblock)
			err := k.pendingApplyChanges.Set(ctx, valoper.String(), 2_000_000)
			if err != nil {
				fmt.Printf("Error setting pendingApplyChanges %v\n", err)
				return fmt.Errorf("error setting pendingApplyChanges: %w", err)
			}

		}

	}
	return nil
}

func getInjectedData(txs [][]byte) *VoteExtension {
	if len(txs) != 0 {
		var injectedData VoteExtension
		// err := a.keeper.cdc.Unmarshal(txs[0], &injectedData) // TODO: ?
		err := json.Unmarshal(txs[0], &injectedData)
		if err == nil {
			return &injectedData
		}
	}
	return nil
}
