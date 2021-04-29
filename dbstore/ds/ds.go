package ds

import (
	"fmt"
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
		AutoMigrate(&model.Miner{}, &model.File{}, &model.MinerFile{},
			&model.DistributionLog{}, &model.WindowPost{})
}

func (s *DbStore) CreateMiner(miner *model.Miner) error {
	ret := s.db.Create(miner)
	if ret.Error != nil {
		log.Errorf("failed to create miner: %v", ret.Error)
	}
	return ret.Error
}

func (s *DbStore) UpdateMiner(miner *model.Miner) error {
	mm := map[string]interface{}{}
	mm["role"] = miner.Role
	mm["last_active_at"] = miner.LastActiveAt

	ret := s.db.Model(&model.Miner{}).Where("id=?", miner.Id).Updates(mm)
	if ret.Error != nil {
		log.Errorf("%v", ret.Error.Error())
		return ret.Error
	}
	return nil
}

func (s *DbStore) GetMiner(id string) (miner model.Miner, err error) {
	ret := s.db.Model(&model.Miner{}).Where("id=?", id).First(&miner)
	if ret.Error != nil {
		log.Errorf("failed to get miner: %v", ret.Error)
	}
	return miner, ret.Error
}

func (s *DbStore) MinerExist(id string) bool {
	var miner model.Miner
	ret := s.db.Model(&model.Miner{}).Where("id=?", id).First(&miner)
	if ret.Error != nil {
		return false
	}
	return true
}

func (s *DbStore) GetMiners(role, state, limit, offset int) (list []*model.Miner, err error) {
	sql := s.db.Model(&model.Miner{})
	if role != -1 {
		sql = sql.Where("role=?", role)
	}
	if state != -1 {
		sql = sql.Where("state=?", state)
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

func (s *DbStore) GetAllMiners(role, state int) (miners []*model.Miner, err error) {
	limit := 1000
	for i := 0; ; i++ {
		list, _ := s.GetMiners(role, state, limit, limit*i)
		if len(list) > 0 {
			miners = append(miners, list...)
		}
		if len(list) < limit {
			break
		}
	}
	return miners, nil
}

func (s *DbStore) GetMaxMinerIndex() (int, error) {
	var miner model.Miner
	ret := s.db.Model(&model.Miner{}).Order("idx DESC").First(&miner)
	if ret.Error != nil {
		if ret.Error != gorm.ErrRecordNotFound {
			log.Errorf("failed to get miner: %v", ret.Error)
			return 0, ret.Error
		}
		return -1, nil
	}
	return miner.Idx, ret.Error
}

func (s *DbStore) FileExist(cid string) bool {
	file := model.File{}
	ret := s.db.Model(&model.File{}).Where("cid=?", cid).First(&file)
	if ret.Error == nil {
		return true
	}
	return false
}

func (s *DbStore) CreateFile(file *model.File) (err error) {
	ret := s.db.Create(file)
	if ret.Error != nil {
		log.Errorf("failed to create file: %v", ret.Error)
	}
	return ret.Error
}

func (s *DbStore) CreateMinerFiles(files []model.MinerFile) (err error) {
	for _, v := range files {
		ret := s.db.Model(&model.MinerFile{}).Create(&v)
		if ret.Error != nil {
			log.Errorf("failed to create miner file: %v", ret.Error)
			return ret.Error
		}
	}
	return nil
}

func (s *DbStore) GetMinerFiles(mineId string) (minerFiles []*MinerFile, err error) {
	//SELECT mf.miner_id, mf.file_cid, f.size from miner_file mf inner join file f on  mf.miner_id = '12D3KooWRee6wn2LtpN1WYsj57avMFDYWbJ2479JnvQG5pm6RztS' and mf.file_cid = f.cid  ;
	sqlFmt := "SELECT mf.miner_id, mf.file_cid, f.size from miner_file mf inner join file f on  mf.miner_id = '%v' and mf.file_cid = f.cid"
	sqlStr := fmt.Sprintf(sqlFmt, mineId)
	ret := s.db.Raw(sqlStr).Scan(&minerFiles)
	if ret.Error != nil {
		log.Errorf("%v", ret.Error.Error())
	}
	return minerFiles, ret.Error
}

func (s *DbStore) CreateDistributionLog(item *model.DistributionLog) error {
	ret := s.db.Create(item)
	if ret.Error != nil {
		log.Errorf("failed to create approveLog: %v", ret.Error)
	}
	return ret.Error
}

func (s *DbStore) GetLastDistributionLog() (*model.DistributionLog, error) {
	var item model.DistributionLog
	ret := s.db.Model(&model.DistributionLog{}).Order("id DESC").First(&item)
	if ret.Error != nil {
		log.Errorf("failed to get approve: %v", ret.Error)
		return nil, ret.Error
	}
	return &item, nil
}

func (s *DbStore) SaveWindowPost(wndPost *model.WindowPost) error {
	ret := s.db.Create(wndPost)
	if ret.Error != nil {
		log.Errorf("failed to save wndPost: %v", ret.Error)
	}
	return ret.Error
}
