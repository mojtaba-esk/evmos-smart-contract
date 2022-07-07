package store_contracts

type CompiledContract struct {
	ABI  string `json:"abi"`
	Bin  string `json:"bin"`
	Name string `json:"contractName"`
}
