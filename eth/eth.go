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
	"math/rand"
	"sync"
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
	Nonce      *big.Int
	mutex      sync.RWMutex
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
		Nonce:      big.NewInt(rand.Int63() % 1000),
	}, nil
}

func (c *Contract) getNonce() *big.Int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	ret := big.NewInt(0).Set(c.Nonce)
	c.Nonce.Add(c.Nonce, big.NewInt(1))
	return ret
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

func (c *Contract) GetAccountCount() (int64, error) {
	ret, err := c.token.GetAccountCount(nil)
	if err != nil {
		log.Errorf("failed to get Decimals, %v", err)
		return 0, err
	}
	return ret.Int64(), nil
}

func (c *Contract) GetAccount(index int64) (string, string, error) {
	address, err := c.token.Accounts(nil, big.NewInt(index))
	if err != nil {
		log.Errorf("failed to get Accounts, %v", err)
		return "", "", err
	}

	peerId, err := c.token.Nodes(nil, big.NewInt(index))
	if err != nil {
		log.Errorf("failed to get Nodes, %v", err)
		return "", "", err
	}

	return address.String(), peerId, nil
}

func (c *Contract) Test() {
	opt, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainId)
	if err != nil {
		log.Errorf("failed to NewKeyedTransactorWithChainID, %v", err)
		return
	}
	opt.GasLimit = c.conf.GasLimit

	addr := common.HexToAddress("0x902Dda3CB9281e44B974f825980a875E056682A3")
	tx, err := c.token.CreateMiner(opt, addr, "12D3KooWP7We1ubNk7U7iQNUihZeDgtcVy8fRQPhVBQs6pjHawm4")
	if err != nil {
		log.Errorf("failed to create miner, %v", err)
		return
	}
	log.Infof("tx: %v", tx.Hash().String())

	//opt.Nonce = c.getNonce()
	tx, err = c.token.SetAllower(opt, addr)
	if err != nil {
		log.Errorf("failed to SetAllower, %v", err)
		return
	}
	log.Infof("tx: %v", tx.Hash().String())

	addr = common.HexToAddress("0xc904025aB325669E1A2C24e217AA02BD3717F68b")
	tx, err = c.token.CreateMiner(opt, addr, "12D3KooWGBCq1wPyj8dinsi6YHMbCkvrjn9nPNX82MxXbgW2MPEv")
	if err != nil {
		log.Errorf("failed to create miner, %v", err)
		return
	}
	log.Infof("tx: %v", tx.Hash().String())
}

func (c *Contract) Approves(data []uint8, len int) (*types.Transaction, error) {
	opt, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainId)
	if err != nil {
		log.Errorf("failed to NewKeyedTransactorWithChainID, %v", err)
		return nil, err
	}
	opt.GasLimit = c.conf.GasLimit

	tx, err := c.token.ApproveWithArray(opt, data, big.NewInt(int64(len)))
	if err != nil {
		log.Errorf("failed to Approval, %v", err)
		return nil, err
	}
	return tx, nil
}
