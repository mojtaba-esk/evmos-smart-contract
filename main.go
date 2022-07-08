package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mojtaba-esk/evmos-smart-contract/cmd"
)

var (
	Bech32Prefix         = "evmos"
	Bech32PrefixAccAddr  = Bech32Prefix
	Bech32PrefixAccPub   = Bech32Prefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

func main() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	cfg.Seal()

	cmd.Execute()
}
