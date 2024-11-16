package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
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

	apiv1 "github.com/reecepbcups/eigenpoa/api/avs/v1"
	"github.com/reecepbcups/eigenpoa/x/avs/keeper/manager"
	"github.com/reecepbcups/eigenpoa/x/avs/types"
)

type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	Eth    *eth.InstrumentedClient
	EthAvs *chainio.AvsReader // not used

	// state management
	Schema   collections.Schema
	Params   collections.Item[types.Params]
	OrmDB    apiv1.StateStore
	valStore baseapp.ValidatorStore

	authority string
}

type Config struct {
	ChainReader     sdkavsregistry.ChainReader
	ServiceBindings *chainio.AvsManagersBindings
}

func NewConfig() *Config {
	return &Config{
		// ChainReader: sdkavsregistry.NewChainReader(),
	}

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

	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("ethHttp", reg) // TODO: POA?
	eth, err := eth.NewInstrumentedClient("http://127.0.0.1:8545", rpcCallsCollector)
	if err != nil {
		panic(err)
	}

	// I have to use the incredible-squaring-avs/ repo only i think?
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

		Eth: eth,
		// EthAvs: ethAvs,

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

	address := common.HexToAddress("0xf5059a5d33d5853360d16c683c16e67980206f36") // set in state, gather on setup
	// instance, err := stake_registry.NewStakeRegistry(address, k.Eth)
	// if err != nil {
	// 	panic(err)
	// }

	m, err := manager.NewManager(address, k.Eth)
	if err != nil {
		fmt.Printf("Error creating manager %v\n", err)
		return nil, fmt.Errorf("error creating manager: %w", err)
	}

	// TODO: I need to get all... how plz
	// weight, err := instance.GetOperatorWeightAtBlock(, common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"), uint32(ethBlockHeight))
	// if err != nil {
	// 	fmt.Printf("Error fetching operator weight %v\n", err)
	// 	return nil, fmt.Errorf("error fetching operator weight: %w", err)
	// }

	// instance.GetOperators

	// read from operatorStateRetriever

	opts := &bind.CallOpts{Context: ctx}
	// registryCoordinatorAddr := ""

	// read from escaped sequence manager.ManagerMetaData.ABI
	// r := strings.NewReader(manager.ManagerMetaData.ABI)
	// contractABI, err := abi.JSON(r)
	// if err != nil {
	// 	fmt.Printf("Error reading abi %v\n", err)
	// 	return nil, fmt.Errorf("error reading abi: %w", err)
	// }

	ethOperators, err := m.GetOperators(opts)
	fmt.Println("ETH OPERATORS", ethOperators, err)
	if err != nil {
		fmt.Printf("Error fetching operators %v\n", err)
		return nil, fmt.Errorf("error fetching operators: %w", err)
	}

	// address common.Address, abi abi.ABI, caller ContractCaller, transactor ContractTransactor, filterer ContractFilterer
	// b := bind.NewBoundContract(address, contractABI, m., m.ManagerTransactor, m.ManagerFilterer)

	// v, err := getOperators(b, opts, ethBlockHeight)

	// k.EthAvs.GetOperatorsStakeInQuorumsAtBlock // ideally we use this but too confusing how to setup for me rn. (EthAvs config)

	operatorStakes := []string{}
	// operatorStakes, err := k.EthAvs.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{Context: ctx}, quorumNumbers, uint32(ethBlockHeight))
	// operatorStakes, err := k.EthAvs.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{Context: ctx}, quorumNumbers, uint32(ethBlockHeight))
	// if err != nil {
	// 	fmt.Printf("Error fetching operator stake %v\n", err)
	// 	return nil, fmt.Errorf("error fetching operator stake: %w", err)
	// }

	// if len(operatorStakes) == 0 {
	// 	return nil, fmt.Errorf("no operators found")
	// }

	operators := make([][]byte, 0, len(operatorStakes))
	for _, eOp := range ethOperators {
		// operators = append(operators, operator[0].Operator.Bytes())
		valoper, err := m.GetOperatorCosmosAddress(opts, eOp)
		if err != nil {
			fmt.Printf("Error fetching operator cosmos address %v\n", err)
			return nil, fmt.Errorf("error fetching operator cosmos address: %w", err)
		}

		// convert valoper to bytes
		addr, err := sdk.ValAddressFromBech32(valoper) // TODO: convert to addresscodec
		if err != nil {
			fmt.Printf("Error converting valoper to bytes %v\n", err)
			return nil, fmt.Errorf("error converting valoper to bytes: %w", err)
		}

		operators = append(operators, addr)
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
