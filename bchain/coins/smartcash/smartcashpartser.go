package smartcash

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// network constants
const (
	MainnetMagic wire.BitcoinNet = 0x5ca1ab1e
	TestnetMagic wire.BitcoinNet = 0xcffcbeea
)

// parser parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{63}
	MainNetParams.ScriptHashAddrID = []byte{18}
	MainNetParams.Bech32HRPSegwit = "sc"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{65}
	TestNetParams.ScriptHashAddrID = []byte{21}
	TestNetParams.Bech32HRPSegwit = "sct"
}

// SmartCashParser handle
type SmartCashParser struct {
	*btc.BitcoinLikeParser
}

// NewSmartCashParser returns new SmartCashParser instance
func NewSmartCashParser(params *chaincfg.Params, c *btc.Configuration) *SmartCashParser {
	return &SmartCashParser{BitcoinLikeParser: btc.NewBitcoinLikeParser(params, c)}
}

// GetChainParams contains network parameters for the main SmartCash network
// and the SmartCash Testnet network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
