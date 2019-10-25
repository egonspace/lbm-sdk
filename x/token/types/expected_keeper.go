package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	supply "github.com/cosmos/cosmos-sdk/x/supply/exported"
	iam "github.com/link-chain/link/x/iam/exported"
)

type IamKeeper interface {
	GetPermissions(sdk.Context, sdk.AccAddress) []iam.PermissionI
	InheritPermission(sdk.Context, sdk.AccAddress, sdk.AccAddress)
	GrantPermission(sdk.Context, sdk.AccAddress, iam.PermissionI)
	RevokePermission(sdk.Context, sdk.AccAddress, iam.PermissionI)
	HasPermission(sdk.Context, sdk.AccAddress, iam.PermissionI) bool
	WithPrefix(string) iam.IamKeeper
}

type SupplyKeeper interface {
	GetModuleAddress(string) sdk.AccAddress
	MintCoins(sdk.Context, string, sdk.Coins) sdk.Error
	BurnCoins(sdk.Context, string, sdk.Coins) sdk.Error
	GetSupply(sdk.Context) supply.SupplyI
	SendCoinsFromModuleToAccount(sdk.Context, string, sdk.AccAddress, sdk.Coins) sdk.Error
	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) sdk.Error
}
