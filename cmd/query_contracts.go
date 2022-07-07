package cmd

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"

	"github.com/mojtaba-esk/deploy-smart-contract/contract"
)

var queryContractsCmd = &cobra.Command{
	Use:   "query [contract_name] [contract_address] [method_to_call]",
	Short: "Query a contract to the evmos chain",
	Long:  `This command receives a contract address and queries it`,
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {

		clientCtx, err := client.GetClientTxContext(cmd)
		if err != nil {
			return err
		}

		contractJsonFilePath := getCompiledContractPath(cmd, args[0])

		instance, err := getContractInstance(contractJsonFilePath, args[1], clientCtx.NodeURI)
		if err != nil {
			return err
		}

		var out []interface{}
		if err := instance.Call(nil, &out, args[2]); err != nil {
			return err
		}

		// out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		fmt.Printf("output: %v\n", out)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(queryContractsCmd)
}

// This function receives the address of a contract and
// a nodeURI then returns the instance of the contract
func getContractInstance(contractJsonFilePath string, contractAddress string, nodeURI string) (*bind.BoundContract, error) {

	ec, err := ethclient.Dial(nodeURI)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)
	return contract.Instance(contractJsonFilePath, address, ec)
}
