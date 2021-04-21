package eth

import (
	"ipfc/utils/encoding"
	"testing"
)

var conf = Config{
	Network:         "https://goerli.infura.io/v3/f81902bab6204bb88b03a1507225fe0a",
	ContractAddress: "0x7b63F4B3d97f509bE4fCbb5a85e2d0672b9Fda4a",
	PrivateKey:      "4a3257d745cd03a02293f774c781342bd063f4eafde6575c1f22560dbb8eeee5",
	GasLimit:        5718749,
	GasPrice:        10,
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
