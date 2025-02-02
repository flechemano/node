package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/fungible/types"
)

// SetZetaDepositAndCallContract set zetaDepositAndCallContract in the store
func (k Keeper) SetSystemContract(ctx sdk.Context, sytemContract types.SystemContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemContractKey))
	b := k.cdc.MustMarshal(&sytemContract)
	store.Set([]byte{0}, b)
}

// GetZetaDepositAndCallContract returns zetaDepositAndCallContract
func (k Keeper) GetSystemContract(ctx sdk.Context) (val types.SystemContract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemContractKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveZetaDepositAndCallContract removes zetaDepositAndCallContract from the store
func (k Keeper) RemoveSystemContract(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemContractKey))
	store.Delete([]byte{0})
}
