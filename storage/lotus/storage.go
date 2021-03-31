package lotus

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	types2 "github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
	"ipfc/storage/types"
	"os"
	"time"
)

var log = logging.Logger("lotus")

type DealConfig struct {
	Miners      []string `yaml:"miners"`
	FundAddress string   `yaml:"fund_address"`
	Copies      int      `yaml:"copies"`
}

type Storage struct {
	node          api.FullNode
	closer        jsonrpc.ClientCloser
	dealConf      DealConfig
	minerAsks     map[string]*storagemarket.StorageAsk
	minerSelector types.MinerSelector
	FundAddress   address.Address
}

func NewStorage(apiAddr, token string, dealConf DealConfig) (*Storage, error) {
	node, closer, err := GetFullNodeAPIUsingCredentials(context.Background(), apiAddr, token)
	if err != nil {
		return nil, err
	}
	stor := &Storage{
		node:     node,
		closer:   closer,
		dealConf: dealConf,
	}
	err = stor.init()
	return stor, err
}

func (s *Storage) init() error {
	addr, err := address.NewFromString(s.dealConf.FundAddress)
	if err != nil {
		log.Errorf("tailed to new address: %v", err)
		return err
	}
	s.FundAddress = addr
	minerAsks := make(map[string]*storagemarket.StorageAsk)
	for _, miner := range s.dealConf.Miners {
		addr, err := address.NewFromString(miner)
		if err != nil {
			log.Errorf("tailed to new address: %v", err)
			return err
		}
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			mi, err := s.node.StateMinerInfo(ctx, addr, types2.EmptyTSK)
			if err != nil {
				log.Errorf("failed to get peerID for miner: %w", err)
			}
			ask, err := s.node.ClientQueryAsk(ctx, *mi.PeerId, addr)
			if err != nil {
				log.Errorf("tailed to new address: %v", err)
				return
			}
			minerAsks[miner] = ask
			log.Infof("miner: %+v", *ask)
		}()
	}
	s.minerAsks = minerAsks
	s.minerSelector = NewRandomMinerSelector(minerAsks, s.dealConf.Copies)
	return nil
}

func (s *Storage) AddFile(ctx context.Context, filePath string) (cid string, err error) {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		log.Errorf("failed to get file stat: %v, filePath=%v", err, filePath)
		return "", err
	}
	miners, err := s.minerSelector.GetMiners(ctx, &types.FileInfo{Path: filePath, Size: uint64(fileStat.Size())})
	if err != nil {
		log.Errorf("no available miner found: %v", err)
		return "", err
	}
	if len(miners) == 0 {
		log.Errorf("no available miner found")
		return "", xerrors.New("no available miner found")
	}
	if len(miners) < s.dealConf.Copies {
		log.Warnf("need %v miners, but %v satisfy", s.dealConf.Copies, len(miners))
	}
	ref := api.FileRef{Path: filePath}
	res, err := s.node.ClientImport(ctx, ref)
	if err != nil {
		log.Errorf("failed to import file: %v", err)
		return "", err
	}
	cid = res.Root.String()
	//TODO: 数据过期后的处理
	for _, miner := range miners {
		ask := s.minerAsks[miner]
		param := &api.StartDealParams{
			Data: &storagemarket.DataRef{
				TransferType: storagemarket.TTGraphsync,
				Root:         res.Root,
			},
			Wallet:            s.FundAddress,
			Miner:             ask.Miner,
			EpochPrice:        ask.Price,
			DealStartEpoch:    0, // 当前高度
			MinBlocksDuration: uint64(build.MinDealDuration),
			FastRetrieval:     false,
		}
		_, err := s.node.ClientStartDeal(ctx, param)
		if err != nil {
			log.Errorf("failed to start deal: %v, miner=%v", err, miner)
		}
	}
	return
}

func (s *Storage) RetrieveFile(ctx context.Context, cid, outputPath string) error {
	return nil
}

func (s *Storage) addFileToMiner(ctx context.Context, cid string) error {

	return nil
}

var _ types.Storage = &Storage{}
