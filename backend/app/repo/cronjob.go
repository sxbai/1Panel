package repo

import (
	"time"

	"github.com/1Panel-dev/1Panel/app/model"
	"github.com/1Panel-dev/1Panel/constant"
	"github.com/1Panel-dev/1Panel/global"
	"gorm.io/gorm"
)

type CronjobRepo struct{}

type ICronjobRepo interface {
	Get(opts ...DBOption) (model.Cronjob, error)
	Page(limit, offset int, opts ...DBOption) (int64, []model.Cronjob, error)
	Create(cronjob *model.Cronjob) error
	WithByDate(startTime, endTime time.Time) DBOption
	WithByJobID(id int) DBOption
	Save(id uint, cronjob model.Cronjob) error
	Update(id uint, vars map[string]interface{}) error
	Delete(opts ...DBOption) error
	StartRecords(cronjobID uint, targetPath string) model.JobRecords
	EndRecords(record model.JobRecords, status, message, records string)
}

func NewICronjobService() ICronjobRepo {
	return &CronjobRepo{}
}

func (u *CronjobRepo) Get(opts ...DBOption) (model.Cronjob, error) {
	var cronjob model.Cronjob
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(&cronjob).Error
	return cronjob, err
}

func (u *CronjobRepo) Page(page, size int, opts ...DBOption) (int64, []model.Cronjob, error) {
	var users []model.Cronjob
	db := global.DB.Model(&model.Cronjob{})
	for _, opt := range opts {
		db = opt(db)
	}
	count := int64(0)
	db = db.Count(&count)
	err := db.Limit(size).Offset(size * (page - 1)).Find(&users).Error
	return count, users, err
}

func (u *CronjobRepo) PageRecords(page, size int, opts ...DBOption) (int64, []model.JobRecords, error) {
	var users []model.JobRecords
	db := global.DB.Model(&model.JobRecords{})
	for _, opt := range opts {
		db = opt(db)
	}
	count := int64(0)
	db = db.Count(&count)
	err := db.Limit(size).Offset(size * (page - 1)).Find(&users).Error
	return count, users, err
}

func (u *CronjobRepo) Create(cronjob *model.Cronjob) error {
	return global.DB.Create(cronjob).Error
}

func (c *CronjobRepo) WithByDate(startTime, endTime time.Time) DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("start_time > ? AND start_time < ?", startTime, endTime)
	}
}
func (c *CronjobRepo) WithByJobID(id int) DBOption {
	return func(g *gorm.DB) *gorm.DB {
		return g.Where("cronjob_id = ?", id)
	}
}

func (u *CronjobRepo) StartRecords(cronjobID uint, targetPath string) model.JobRecords {
	var record model.JobRecords
	record.StartTime = time.Now()
	record.CronjobID = cronjobID
	record.Status = constant.StatusRunning
	if err := global.DB.Create(&record).Error; err != nil {
		global.LOG.Errorf("create record status failed, err: %v", err)
	}
	return record
}
func (u *CronjobRepo) EndRecords(record model.JobRecords, status, message, records string) {
	errMap := make(map[string]interface{})
	errMap["records"] = records
	errMap["status"] = status
	errMap["message"] = message
	errMap["interval"] = time.Since(record.StartTime).Milliseconds()
	if err := global.DB.Model(&model.JobRecords{}).Where("id = ?", record.ID).Updates(errMap).Error; err != nil {
		global.LOG.Errorf("update record status failed, err: %v", err)
	}
}

func (u *CronjobRepo) Save(id uint, cronjob model.Cronjob) error {
	return global.DB.Model(&model.Cronjob{}).Where("id = ?", id).Save(&cronjob).Error
}
func (u *CronjobRepo) Update(id uint, vars map[string]interface{}) error {
	return global.DB.Model(&model.Cronjob{}).Where("id = ?", id).Updates(vars).Error
}

func (u *CronjobRepo) Delete(opts ...DBOption) error {
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	return db.Delete(&model.Cronjob{}).Error
}