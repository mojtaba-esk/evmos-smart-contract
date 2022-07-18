package contract_test

import (
	"fmt"
	"testing"

	"github.com/mojtaba-esk/evmos-smart-contract/cmd"
	"github.com/mojtaba-esk/evmos-smart-contract/contract"
)

func TestDeploy(t *testing.T) {

	tt := []struct {
		name         string
		contractName string
		wantErr      bool
		wantOutput   string
	}{
		{
			name:         "Contract without parameters",
			contractName: "Counter1",
			wantErr:      false,
			wantOutput:   "",
		},
	}

	TestDataDir = "/home/moji/.evmosd"
	privateKey, err := cmd.GetPrivateKey(TestKey, "test", TestDataDir, "", cmd.AppName, nil)
	if err != nil {
		t.Errorf("Cannot get the private key: %v", err)
		return
	}
	fmt.Printf("privateKey: %v\n", privateKey)

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			jsonFilePath, err := GetContractJsonFilePath(tc.contractName)
			if err != nil {
				t.Errorf("Error getting contract json file path: %v", err)
			}

			address, tx, _, err := contract.Deploy(jsonFilePath, privateKey, TestNodeURI, "")
			fmt.Println("Contract Address: ", address.Hex())
			fmt.Println("TX Hash: ", tx.Hash().Hex())

			if (err != nil) != tc.wantErr {
				t.Errorf("Deploy() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// if !reflect.DeepEqual(gotOutput, tc.wantOutput) {
			// 	t.Errorf("Deploy() = %v, want %v", gotOutput, tc.wantOutput)
			// }
		})
	}
}
