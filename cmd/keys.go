package cmd

import ethermintKeys "github.com/evmos/ethermint/client/keys"

func init() {
	rootCmd.AddCommand(ethermintKeys.Commands(defaultNodeHome))
}
