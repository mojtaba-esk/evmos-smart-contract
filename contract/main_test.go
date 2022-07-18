package contract_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mojtaba-esk/evmos-smart-contract/contract"
)

const (
	dataDirFileContainerPath = "/tmp/evmos-test-data-dir" // The path to the data dir in /tmp is written into this file on build
	TestKey                  = "mykey"
	TestNodeURI              = "http://localhost:8545"
)

var TestDataDir = ""

func init() {

	// Get the data dir for the test running node in the tmp dir
	bytes, err := ioutil.ReadFile(dataDirFileContainerPath)
	if err != nil {
		panic(fmt.Errorf("dataDirFileContainerPath: %v", err))
	}

	TestDataDir = strings.Trim(string(bytes), "\r\n\t ")
	if TestDataDir == "" {
		panic(fmt.Errorf("dataDirFileContainerPath: the path is empty"))
	}
}

// Disabling logs
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestParseJsonParams(t *testing.T) {

	tt := []struct {
		name       string
		input      string
		wantErr    bool
		wantOutput []interface{}
	}{
		{
			name:       "Empty input",
			input:      "",
			wantErr:    false,
			wantOutput: []interface{}{},
		},
		{
			name:       "Empty parameters",
			input:      `{"params":[]}`,
			wantErr:    false,
			wantOutput: []interface{}{},
		},
		{
			name:       "Malformed Json parameters",
			input:      `{"params":[12,]}`,
			wantErr:    true,
			wantOutput: nil,
		},
		{
			name:    "Numeric parameters",
			input:   `{"params":[0,1,2,3,4,5,6,7,8,9,10]}`,
			wantErr: false,
			wantOutput: []interface{}{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(4),
				big.NewInt(5),
				big.NewInt(6),
				big.NewInt(7),
				big.NewInt(8),
				big.NewInt(9),
				big.NewInt(10),
			},
		},
		{
			name:    "Float point numbers",
			input:   `{"params":[0.5,1.2,2.000014,3.01111,4.999999]}`,
			wantErr: false,
			wantOutput: []interface{}{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
				big.NewInt(4),
			},
		},
		{
			name:    "Mix string, numeric params",
			input:   `{"params":[0,1,2,3,"test","Ciao"]}`,
			wantErr: false,
			wantOutput: []interface{}{
				big.NewInt(0),
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
				"test",
				"Ciao",
			},
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			gotOutput, err := contract.ParseJsonParams(tc.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("ParseJsonParams() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(gotOutput, tc.wantOutput) {
				t.Errorf("ParseJsonParams() = %v, want %v", gotOutput, tc.wantOutput)
			}
		})
	}
}
