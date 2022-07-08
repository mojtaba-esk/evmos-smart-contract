package contract

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

type CompiledContract struct {
	ABI  string `json:"abi"`
	Bin  string `json:"bin"`
	Name string `json:"contractName"`
}

const (
	METHOD_NAME       = "name"
	METHOD_BALANCE_OF = "balanceOf"
	METHOD_SYMBOL     = "symbol"
	METHOD_DECIMALS   = "decimals"
	METHOD_TRANSFER   = "transfer"
)

// This function parses the given json value for contract method's parameters
// It extracts the params in the same given order and converts the simple types
// such as numbers and strings
// Example: {"params":[51, "test msg"]}
func ParseJsonParams(paramsJson string) ([]interface{}, error) {

	var params map[string]interface{}
	if err := json.Unmarshal([]byte(paramsJson), &params); err != nil {
		return nil, fmt.Errorf("unmarshaling params failed: %v", err)
	}

	if len(params) == 0 {
		return nil, fmt.Errorf("unmarshaling params: empty parameters list")
	}

	var paramsToPass []interface{}
	// Let's range over it to let user to enter any key name they want for params
	for i := range params {
		paramsToPass = params[i].([]interface{})
		break
	}

	// Fix the numbers to make it compatible with int ptr
	for i, v := range paramsToPass {
		if reflect.TypeOf(v).Kind() == reflect.Float64 {
			numParam := &big.Int{}
			numParam.SetInt64(int64(v.(float64)))
			paramsToPass[i] = numParam
		}
	}

	return paramsToPass, nil

}
