package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/iov-one/iovnsd/x/configuration/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// group config queries under a sub-command
	configQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	// add queries
	configQueryCmd.AddCommand(
		flags.GetCommands(
			getCmdQueryConfig(queryRoute, cdc),
		)...,
	)
	// return cmd list
	return configQueryCmd
}

// getCmdQueryConfig returns the command to get the configuration
func getCmdQueryConfig(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "query",
		Short: "query configuration",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			path := fmt.Sprintf("custom/%s/%s", route, types.QuerierRoute)
			resp, _, err := cliCtx.Query(path)
			if err != nil {
				return err
			}
			var jsonResp types.Config // TODO make this a separate type
			cdc.MustUnmarshalJSON(resp, &jsonResp)
			return cliCtx.PrintOutput(jsonResp)
		},
	}
}