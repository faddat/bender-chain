package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/faddat/bender-chain/x/benderchain/types"
)

func CmdCreateAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-add [post] [title] [body] [ipfs]",
		Short: "Creates a new add",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPost := string(args[0])
			argsTitle := string(args[1])
			argsBody := string(args[2])
			argsIpfs := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAdd(clientCtx.GetFromAddress().String(), string(argsPost), string(argsTitle), string(argsBody), string(argsIpfs))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-add [id] [post] [title] [body] [ipfs]",
		Short: "Update a add",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsPost := string(args[1])
			argsTitle := string(args[2])
			argsBody := string(args[3])
			argsIpfs := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAdd(clientCtx.GetFromAddress().String(), id, string(argsPost), string(argsTitle), string(argsBody), string(argsIpfs))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-add [id] [post] [title] [body] [ipfs]",
		Short: "Delete a add by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAdd(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
