package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (keeper *Keeper) BeginBlocker(ctx context.Context) error {
	rm := []string{}
	vals, err := keeper.pendingApplyChanges.Iterate(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to iterate pending apply changes: %w", err)
	}

	for ; vals.Valid(); vals.Next() {
		k, err := vals.Key()
		if err != nil {
			return fmt.Errorf("failed to get key: %w", err)
		}
		v, err := vals.Value()
		if err != nil {
			return fmt.Errorf("failed to get value: %w", err)
		}

		rm = append(rm, string(k))

		fmt.Println("-- BeginBlocker pendingApplyChanges", k, v)

		// convert k to sdk.ValAddress
		addr, err := sdk.ValAddressFromBech32(string(k))
		if err != nil {
			return fmt.Errorf("failed to convert key to sdk.ValAddress: %w", err)
		}

		vals, err := keeper.poaKeeper.SetPOAPower(ctx, addr.String(), int64(v))
		fmt.Printf("BeginBlocker injectedData operator power %v | %v\n", vals, err)
	}

	// iter rm and remove them
	for _, k := range rm {
		err := keeper.pendingApplyChanges.Remove(ctx, k)
		if err != nil {
			return fmt.Errorf("failed to remove key: %w", err)
		}
	}

	return nil
}

func (keeper *Keeper) EndBlocker(ctx context.Context) error {
	return nil
}
