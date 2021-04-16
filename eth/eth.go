package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	logging "github.com/ipfs/go-log/v2"
	"ipfc/eth/contract"
	"math/big"
)

var log = logging.Logger("eth")

type Config struct {
	Network         string `yaml:"network"`
	ContractAddress string `yaml:"contract_address"`
	PrivateKey      string `yaml:"private_key"`
	GasLimit        uint64 `yaml:"gas_limit"`
}

type Contract struct {
	conf       Config
	token      *contract.FCToken
	chainId    *big.Int
	privateKey *ecdsa.PrivateKey
}

func NewContract(conf Config) (*Contract, error) {
	client, err := ethclient.Dial(conf.Network)
	if err != nil {
		log.Errorf("Failed to connect to eth: %v", err)
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("Failed to get chainId: %v", err)
		return nil, err
	}
	token, err := contract.NewFCToken(common.HexToAddress(conf.ContractAddress), client)
	if err != nil {
		log.Errorf("Failed to instantiate a Token contract: %v", err)
		return nil, err
	}
	key, err := crypto.HexToECDSA(conf.PrivateKey)
	if err != nil {
		log.Errorf("Failed to decode private key: %v", err)
		return nil, err
	}

	return &Contract{
		conf:       conf,
		token:      token,
		chainId:    chainId,
		privateKey: key,
	}, nil
}

func (c *Contract) TotalSupply() (*big.Int, error) {
	s, err := c.token.TotalSupply(nil)
	if err != nil {
		log.Errorf("failed to get TotalSupply, %v", err)
		return nil, err
	}

	return s, nil
}

func (c *Contract) GetSymbol() (string, error) {
	s, err := c.token.Symbol(nil)
	if err != nil {
		log.Errorf("failed to get TotalSupply, %v", err)
		return "", err
	}
	return s, nil
}

func (c *Contract) GetDecimals() (uint8, error) {
	s, err := c.token.Decimals(nil)
	if err != nil {
		log.Errorf("failed to get Decimals, %v", err)
		return 0, err
	}
	return s, nil
}

func (c *Contract) Approve(address string, value *big.Int) (*types.Transaction, error) {
	opt, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainId)
	if err != nil {
		log.Errorf("failed to NewKeyedTransactorWithChainID, %v", err)
		return nil, err
	}
	opt.GasLimit = c.conf.GasLimit
	tx, err := c.token.Approval(opt, common.HexToAddress(address), value)
	if err != nil {
		log.Errorf("failed to Approval, %v", err)
		return nil, err
	}
	return tx, nil
}

func (c *Contract) Approves(addresses []string, value []*big.Int) (*types.Transaction, error) {
	opt, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainId)
	if err != nil {
		log.Errorf("failed to NewKeyedTransactorWithChainID, %v", err)
		return nil, err
	}
	opt.GasLimit = c.conf.GasLimit
	addrs := make([]common.Address, len(addresses))
	for i, addr := range addresses {
		addrs[i] = common.HexToAddress(addr)
	}
	tx, err := c.token.ApproveWthArray(opt, addrs, value)
	if err != nil {
		log.Errorf("failed to Approval, %v", err)
		return nil, err
	}
	return tx, nil
}
