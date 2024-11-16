package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func getOperators(contract *bind.BoundContract, opts *bind.CallOpts, height uint64) ([][]common.Address, error) {
	// quorumNumbers := eigentypes.QuorumNums{0} // TODO: what is this?

	// cast call 0xf5059a5d33d5853360d16c683c16e67980206f36 "getOperators()(address[])"

	var out []interface{}
	// err := contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, height) // original
	err := contract.Call(opts, &out, "getOperators") // TODO: the cast works, but not the go bind call. need to play around with
	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)
	return out0, err
}

func main() {
	ctx := context.Background()

	address := common.HexToAddress("0xf5059a5d33d5853360d16c683c16e67980206f36") // set in state, gather on setup

	opts := &bind.CallOpts{Context: ctx}

	r := strings.NewReader(manager.ManagerMetaData.ABI)
	contractABI, err := abi.JSON(r)
	if err != nil {
		fmt.Printf("Error reading abi %v\n", err)
		return nil, fmt.Errorf("error reading abi: %w", err)
	}

	b := bind.NewBoundContract(address, contractABI, nil, nil, nil)

	v, err := getOperators(b, opts, ethBlockHeight)

}
