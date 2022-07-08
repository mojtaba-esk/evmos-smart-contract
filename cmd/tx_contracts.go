package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	cosmosFlags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mojtaba-esk/evmos-smart-contract/contract"

	"github.com/spf13/cobra"
)

var txContractsCmd = &cobra.Command{
	Use:   "tx [contract_name] [contract_address] [method_to_call] [method_arguments_json]",
	Short: "Call a method of a contract that modifies the state",
	Args:  cobra.RangeArgs(3, 4),
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

		ec, err := ethclient.Dial(clientCtx.NodeURI)
		if err != nil {
			return err
		}

		contractJsonFilePath := getCompiledContractPath(cmd, args[0])
		contractAddress := args[1]
		address := common.HexToAddress(contractAddress)
		instance, err := contract.Instance(contractJsonFilePath, address, ec)
		if err != nil {
			return err
		}

		/*--------*/

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("invalid key")
		}

		fromAddress := ethCrypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return err
		}

		chainID, err := ec.ChainID(context.Background())
		if err != nil {
			return err
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			return err
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = big.NewInt(1000000)

		/*--------*/

		var tx *ethTypes.Transaction
		method := args[2]

		if len(args) == 4 {
			methodParams := args[3]
			tx, err = instance.Transact(auth, method, methodParams)
		} else {
			tx, err = instance.Transact(auth, method)
		}

		if err != nil {
			return err
		}

		fmt.Println("TX Hash: ", tx.Hash().String())

		return nil
	},
}

func init() {

	rootCmd.AddCommand(txContractsCmd)
	txContractsCmd.Flags().String(cosmosFlags.FlagFrom, "", "account address to sign the deployment tx")
}

// // This function receives the path to the compiled contract in JSON, a private key, a nodeURI
// // and a method then it deploys the smart contract to the chain via the given evmos node
// func txContract(contractJsonFilePath string, privateKey *ecdsa.PrivateKey, nodeURI, method string, params string) (common.Address, *ethTypes.Transaction, *bind.BoundContract, error) {

// 	ec, err := ethclient.Dial(nodeURI)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return common.Address{}, nil, nil, fmt.Errorf("invalid key")
// 	}

// 	fromAddress := ethCrypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := ec.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	chainID, err := ec.ChainID(context.Background())
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)      // in wei
// 	auth.GasLimit = uint64(3000000) // in units
// 	auth.GasPrice = big.NewInt(1000000)

// 	return store.Deploy(contractJsonFilePath, auth, ec, initMsg)
// }
