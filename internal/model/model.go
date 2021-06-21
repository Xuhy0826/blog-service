package model

import (
	"blog-service/pkg/setting"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID         uint32    `gorm:"primary_key" json:"id"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	DeletedOn  time.Time `json:"deleted_on"`
	IsDel      uint8     `json:"is_del"`
}

func NewDBEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbSetting.Host,
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.DBName,
		dbSetting.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

//创建后
func (m *Model) BeforeSave(tx *gorm.DB) (err error) {
	m.CreatedOn = time.Now()
	//tx.Model(m).Updates(map[string]interface{}{"created_by": m.CreatedBy, "created_on": nowTime})
	return
}

//更新后
func (m *Model) AfterUpdate(tx *gorm.DB) (err error) {
	nowTime := time.Now().Unix()
	tx.Model(m).Updates(map[string]interface{}{"modified_by": m.ModifiedBy, "modified_on": nowTime})
	return
}

//删除后
func (m *Model) AfterDelete(tx *gorm.DB) (err error) {
	nowTime := time.Now().Unix()
	tx.Model(m).Updates(map[string]interface{}{"is_del": 1, "deleted_on": nowTime})
	return
}
