package smartcash

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/SmartCash/blockbook/bchain"
	"github.com/SmartCash/blockbook/bchain/coins/btc"
)

// SmartCashRPC is an interface to JSON-RPC bitcoind service.
type SmartCashRPC struct {
	*btc.BitcoinRPC
}

// NewSmartCashRPC returns new SmartCashRPC instance.
func NewSmartCashRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &SmartCashRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV2{}
	s.ChainConfig.SupportsEstimateFee = false

	return s, nil
}

// Initialize initializes SmartCashRPC instance.
func (b *SmartCashRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewSmartCashParser(params, b.ChainConfig)

	b.Testnet = false
	b.Network = "livenet"

	glog.Info("rpc: block chain ", params.Name)

	return nil
}
