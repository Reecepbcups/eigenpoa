package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	eigentypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	abci "github.com/cometbft/cometbft/abci/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	apiv1 "github.com/reecepbcups/eigenpoa/api/avs/v1"
	"github.com/reecepbcups/eigenpoa/x/avs/types"
)

type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	Eth    *eth.InstrumentedClient
	EthAvs *chainio.AvsReader

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

	// TODO: nil -> &rpccalls.NewCollector("poa", ) ?
	eth, err := eth.NewInstrumentedClient("http://127.0.0.1:8545", nil)
	if err != nil {
		panic(err)
	}

	k := Keeper{
		cdc:      cdc,
		logger:   logger,
		valStore: valStore,

		Eth:    eth,
		EthAvs: nil,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:  store,

		authority: authority,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

func (k Keeper) Logger() log.Logger {
	return k.logger
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

func (k *Keeper) GetOperators(ctx context.Context, ethBlockHeight uint64) ([][]byte, error) {
	quorumNumbers := eigentypes.QuorumNums{0} // TODO: what is this?

	operatorStakes, err := k.EthAvs.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{Context: ctx}, quorumNumbers, uint32(ethBlockHeight))
	if err != nil {
		fmt.Printf("Error fetching operator stake %v\n", err)
		return nil, fmt.Errorf("error fetching operator stake: %w", err)
	}

	if len(operatorStakes) == 0 {
		return nil, fmt.Errorf("no operators found")
	}

	operators := make([][]byte, 0, len(operatorStakes))
	for _, operator := range operatorStakes {
		operators = append(operators, operator[0].Operator.Bytes())
	}
	return operators, nil
}

// This request does not match the sdk's checked PreBlocker interface (why??) so we have to manually impl this in the app PreBlocker.
func (k *Keeper) PreBlock(_ context.Context, req *abci.RequestFinalizeBlock) error {
	injectedData := getInjectedData(req.Txs)
	if injectedData != nil {
		fmt.Println("PreBlock injectedData", injectedData, "TODO: put the POA logic here :D") // TODO:
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
