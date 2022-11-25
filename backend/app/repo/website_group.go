package repo

import (
	"github.com/1Panel-dev/1Panel/backend/app/model"
	"github.com/1Panel-dev/1Panel/backend/global"
	"gorm.io/gorm/clause"
)

type WebSiteGroupRepo struct {
}

func (w WebSiteGroupRepo) Page(page, size int, opts ...DBOption) (int64, []model.WebSiteGroup, error) {
	var groups []model.WebSiteGroup
	db := getDb(opts...).Model(&model.WebSiteGroup{})
	count := int64(0)
	db = db.Count(&count)
	err := db.Limit(size).Offset(size * (page - 1)).Order("`default` desc").Find(&groups).Error
	return count, groups, err
}

func (w WebSiteGroupRepo) GetBy(opts ...DBOption) ([]model.WebSiteGroup, error) {
	var groups []model.WebSiteGroup
	db := getDb(opts...).Model(&model.WebSiteGroup{})
	if err := db.Order("`default` desc").Find(&groups).Error; err != nil {
		return groups, err
	}
	return groups, nil
}

func (w WebSiteGroupRepo) Create(app *model.WebSiteGroup) error {
	return getDb().Omit(clause.Associations).Create(app).Error
}

func (w WebSiteGroupRepo) Save(app *model.WebSiteGroup) error {
	return getDb().Omit(clause.Associations).Save(app).Error
}

func (w WebSiteGroupRepo) DeleteBy(opts ...DBOption) error {
	return getDb(opts...).Delete(&model.WebSiteGroup{}).Error
}

func (w WebSiteGroupRepo) CancelDefault() error {
	return global.DB.Model(&model.WebSiteGroup{}).Where("`default` = 1").Updates(map[string]interface{}{"default": 0}).Error
}