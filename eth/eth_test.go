package eth

import (
	"ipfc/utils/encoding"
	"testing"
)

var conf = Config{
	//Network:         "https://goerli.infura.io/v3/f81902bab6204bb88b03a1507225fe0a",
	//Network:         "https://data-seed-prebsc-1-s2.binance.org:8545/",
	Network:         "https://data-seed-prebsc-1-s1.binance.org:8545/",
	ContractAddress: "0xc5E80ab6109472E8FF21841b3e96F125592C6e37",
	PrivateKey:      "080ebaab2013d69ac721ff891060f5fdb60cf1a2e353db77f3bca7e5e01de0ff",
	GasLimit:        3141592,
}

func TestTest(t *testing.T) {
	contract, _ := NewContract(conf)

	contract.Test()

	symbol, err := contract.GetSymbol()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("symbol: %v", symbol)

	supply, err := contract.TotalSupply()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("total supply: %v", supply.String())

	decimals, err := contract.GetDecimals()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("decimals: %v", decimals)

	count, err := contract.GetAccountCount()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("account count: %v", count)
	for i := int64(0); i < count; i++ {
		address, id, err := contract.GetAccount(i)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		log.Infof("account: %v %v", address, id)
	}
}

func Test2(t *testing.T) {
	contract, _ := NewContract(conf)
	idx := []int{0}
	data := encoding.EncodeIndex(idx)
	tx, err := contract.Approves(data, len(idx))
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	log.Infof("tx: %v", tx.Hash().String())
}
