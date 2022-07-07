package cmd

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	cosmosFlags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	sdkCrypto "github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	evmoskr "github.com/evmos/evmos/v6/crypto/keyring"

	"github.com/spf13/cobra"

	store "github.com/mojtaba-esk/deploy-smart-contract/store_contracts"
)

const FlagCompiledContractsPath = "compiled-contracts-path"

var deployContractsCmd = &cobra.Command{
	Use:   "deploy [contract_name]",
	Short: "Deploy a contract to the evmos chain",
	Long:  `This command receives a contract name and deploys it on the evmos chain`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		clientCtx, err := client.GetClientTxContext(cmd)
		if err != nil {
			return err
		}

		keyName := clientCtx.GetFromName()
		privateKey, err := getPrivateKey(cmd, keyName)
		if err != nil {
			return err
		}

		compiledContractsDir, _ := cmd.Flags().GetString(FlagCompiledContractsPath)
		contractName := args[0]
		contractJsonFilePath := filepath.Join(compiledContractsDir, contractName+".json")

		address, tx, _, err := deployContract(contractJsonFilePath, privateKey, clientCtx.NodeURI)
		if err != nil {
			return err
		}

		fmt.Println("Contract Address: ", address.Hex())
		fmt.Println("TX Hash: ", tx.Hash().Hex())

		return nil
	},
}

// This function receives the path to the compiled contract in JSON, a private key
// and a nodeURI then it deploys the smart contract to the chain via the given evmos node
func deployContract(contractJsonFilePath string, privateKey *ecdsa.PrivateKey, nodeURI string) (common.Address, *ethTypes.Transaction, *store.Store, error) {

	ec, err := ethclient.Dial(nodeURI)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, nil, fmt.Errorf("invalid key")
	}

	fromAddress := ethCrypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	// Retrieving the contract data and bind them
	if err := store.BindMetaData(contractJsonFilePath); err != nil {
		return common.Address{}, nil, nil, err
	}
	return store.DeployStore(auth, ec)
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

func init() {

	rootCmd.AddCommand(deployContractsCmd)
	deployContractsCmd.Flags().String(cosmosFlags.FlagFrom, "", "account address to sign the deployment tx")
	deployContractsCmd.Flags().String(FlagCompiledContractsPath, filepath.Join(GetRootPath(), "contracts", "compiled_contracts"), "The path to the compiled contracts in json format")
}
