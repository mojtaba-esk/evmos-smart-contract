package cmd

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"

	cosmosFlags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	sdkCrypto "github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	ethermintKeys "github.com/evmos/ethermint/client/keys"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	evmoskr "github.com/evmos/evmos/v6/crypto/keyring"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ethermintKeys.Commands(defaultNodeHome))
}

// This function receives a keyName and retrieves its private key
// Please note that this is designed for `ethsecp256k1` algorithm
func getPrivateKey(cmd *cobra.Command, keyName string) (*ecdsa.PrivateKey, error) {

	inBuf := bufio.NewReader(cmd.InOrStdin())
	keyringBackend, _ := cmd.Flags().GetString(cosmosFlags.FlagKeyringBackend)
	keyringDir, _ := cmd.Flags().GetString(cosmosFlags.FlagKeyringDir)

	kr, err := keyring.New(
		AppName,
		keyringBackend,
		keyringDir,
		inBuf,
		evmoskr.Option(),
	)
	if err != nil {
		return nil, err
	}

	decryptPassword := ""
	conf := true

	switch keyringBackend {
	case keyring.BackendFile:
		decryptPassword, err = input.GetPassword(
			"**WARNING this is unsafe, use it only for test**\nEnter key password:",
			inBuf)
	case keyring.BackendOS:
		conf, err = input.GetConfirmation(
			"**WARNING this is unsafe, use it only for test**\nEnter key password:",
			inBuf, cmd.ErrOrStderr())
	}
	if err != nil || !conf {
		return nil, err
	}

	// Exports private key from keybase using password
	armor, err := kr.ExportPrivKeyArmor(keyName, decryptPassword)
	if err != nil {
		return nil, err
	}

	privKey, algo, err := sdkCrypto.UnarmorDecryptPrivKey(armor, decryptPassword)
	if err != nil {
		return nil, err
	}

	if algo != ethsecp256k1.KeyType {
		return nil, fmt.Errorf("invalid key algorithm, got %s, expected %s", algo, ethsecp256k1.KeyType)
	}

	// Converts key to Ethermint secp256k1 implementation
	ethPrivKey, ok := privKey.(*ethsecp256k1.PrivKey)
	if !ok {
		return nil, fmt.Errorf("invalid private key type %T, expected %T", privKey, &ethsecp256k1.PrivKey{})
	}

	return ethPrivKey.ToECDSA()
}
