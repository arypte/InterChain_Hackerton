package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"recipes/x/recipes/types"
)

var _ = strconv.Itoa(0)

func CmdCreateRecipe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-recipe [data] [meta]",
		Short: "Broadcast message createRecipe",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argData := args[0]
			argMeta := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRecipe(
				clientCtx.GetFromAddress().String(),
				argData,
				argMeta,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
