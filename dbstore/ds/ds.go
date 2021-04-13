package ds

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/jinzhu/gorm"
	"ipfc/dbstore/model"
)

var log = logging.Logger("ds")

type DbStore struct {
	db *gorm.DB
}

func NewDbStore(db *gorm.DB) *DbStore {
	ds := &DbStore{
		db: db,
	}
	ds.init()
	return ds
}

func (s *DbStore) init() {
	s.db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&model.Miner{}, &model.File{}, &model.MinerFile{})
}

func (s *DbStore) CreateMiner(miner *model.Miner) error {
	ret := s.db.Create(miner)
	if ret.Error != nil {
		log.Errorf("failed to create miner: %v", ret.Error)
	}
	return ret.Error
}

func (s *DbStore) GetMiner(id string) (miner model.Miner, err error) {
	ret := s.db.Model(&model.Miner{}).Where("id=?", id).First(&miner)
	if ret.Error != nil {
		log.Errorf("failed to get miner: %v", ret.Error)
	}
	return miner, ret.Error
}

func (s *DbStore) GetMiners(role, state, limit, offset int) (list []*model.Miner, err error) {
	sql := s.db
	if role != -1 {
		sql = s.db.Where("role", role)
	}
	if state != -1 {
		sql = sql.Where("state", state)
	}
	if limit != 0 {
		sql = sql.Limit(limit).Offset(offset)
	}
	ret := sql.Find(&list)
	if ret.Error != nil {
		log.Errorf("failed to get miners: %v", ret.Error)
	}
	return list, ret.Error
}

func (s *DbStore) GetAllMiners(role, state int) (list []*model.Miner, err error) {
	limit := 100
	var miners []*model.Miner
	for i := 0; ; i++ {
		list, _ := s.GetMiners(role, state, limit, limit*i)
		if len(list) <= 0 {
			break
		}
		miners = append(miners, list...)
	}
	return miners, nil
}
