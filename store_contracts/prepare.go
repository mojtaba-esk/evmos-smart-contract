package store_contracts

import (
	"encoding/json"
	"io/ioutil"
)

// This function receives the path of a compiled contract in JSON format
// Retrieve the data from the JSON file and binds it in a metadata variable
// To get prepared for deployment
func BindMetaData(contractJsonFilePath string) error {

	bytes, err := ioutil.ReadFile(contractJsonFilePath)
	if err != nil {
		return err
	}

	var contract CompiledContract

	err = json.Unmarshal(bytes, &contract)
	if err != nil {
		return err
	}

	StoreMetaData.ABI = contract.ABI
	StoreMetaData.Bin = contract.Bin

	StoreABI = StoreMetaData.ABI
	StoreBin = StoreMetaData.Bin

	return nil
}
