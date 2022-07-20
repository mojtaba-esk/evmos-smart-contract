package contract_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/go-bip39"

	// evmosapp "github.com/evmos/evmos/v6/app"
	etherminthd "github.com/evmos/ethermint/crypto/hd"
	evmoscfg "github.com/evmos/evmos/v6/cmd/config"
	"github.com/mojtaba-esk/evmos-smart-contract/cmd"
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

	config := sdk.GetConfig()
	evmoscfg.SetBech32Prefixes(config)
	version.Name = cmd.AppName
	config.Seal()

	cmd.GetClientContext() // We need it here to have the codecs configured properly for tests as well
}

func TestMain(m *testing.M) {
	// Disabling logs
	log.SetOutput(ioutil.Discard)

	os.Exit(m.Run())
}

func generateNewAccount(name string) (keyring.Info, error) {

	const mnemonicEntropySize = 256

	kb := keyring.NewInMemory(etherminthd.EthSecp256k1Option())

	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return nil, err
	}

	hdPath := hd.CreateHDPath(60, 0, 0).String()
	info, err := kb.NewAccount(name, mnemonic, "", hdPath, etherminthd.EthSecp256k1)
	return info, err
}
