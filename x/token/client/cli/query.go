package cli

import (
	"github.com/link-chain/link/client"
	"github.com/link-chain/link/x/token/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the token module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		GetTokenCmd(cdc),
		GetTokensCmd(cdc),
		GetPermsCmd(cdc),
	)

	return cmd
}

func GetTokenCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "symbol [symbol]",
		Short: "Query symbol",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewCLIContext().WithCodec(cdc)
			tokenGetter := types.NewTokenRetriever(cliCtx)

			symbol := args[0]
			if err := tokenGetter.EnsureExists(symbol); err != nil {
				return err
			}

			token, err := tokenGetter.GetToken(symbol)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(token)
		},
	}

	return client.GetCommands(cmd)[0]
}

func GetTokensCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "symbols",
		Short: "Query all symbol",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewCLIContext().WithCodec(cdc)
			tokenGetter := types.NewTokenRetriever(cliCtx)

			tokens, err := tokenGetter.GetAllTokens()
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(tokens)
		},
	}

	return client.GetCommands(cmd)[0]
}

func GetPermsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perm [addr]",
		Short: "Get Permission of the Account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewCLIContext().WithCodec(cdc)
			permGetter := types.NewAccountPermissionRetriever(cliCtx)

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			pms, err := permGetter.GetAccountPermission(addr)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(pms)
		},
	}

	return client.GetCommands(cmd)[0]
}
