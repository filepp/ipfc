package eth

import (
	"math/big"
	"testing"
)

func TestTest(t *testing.T) {
	contract, _ := NewContract(Config{
		Network:         "https://ropsten.infura.io/v3/f81902bab6204bb88b03a1507225fe0a",
		ContractAddress: "0x16e2f72518099b64bbD592AB0E3943aE9c41b806",
		PrivateKey: "6b352ad766b818d15e77ae879ccdd31c0f41aa312d0cc20ca1037b690d6a2531",
		GasLimit:  10000000,
	})

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

	tx, err := contract.Approve("0x7e1608cEe279C33067B1dFf480390c91b2765DBb", big.NewInt(100000000000))
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Logf("total supply: %v", tx.Hash().String())
}
