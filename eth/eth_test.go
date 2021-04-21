package eth

import (
	"testing"
)

func TestTest(t *testing.T) {
	contract, _ := NewContract(Config{
		Network:         "https://goerli.infura.io/v3/f81902bab6204bb88b03a1507225fe0a",
		ContractAddress: "0xC061F4a261A64F1d19aFa834ADFfd7182B684c45",
		PrivateKey:      "4a3257d745cd03a02293f774c781342bd063f4eafde6575c1f22560dbb8eeee5",
		GasLimit:        2718749,
		GasPrice:        5,
	})

	//contract.Test()

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
	t.Logf("accunt count: %v", count)
	for i := int64(0); i < count; i++ {
		address, id, err := contract.GetAccount(i)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		log.Infof("account: %v %v", address, id)
	}
}
