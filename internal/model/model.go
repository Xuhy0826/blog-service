package model

import (
	"blog-service/pkg/setting"
	"blog-service/pkg/tracer"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

	//use gorm plugin，trace sql
	_ = db.Use(&tracer.OpentracingPlugin{})
	//span := opentracing.StartSpan("gormTracing")
	//defer span.Finish()
	////把生成的Root Span写入到Context上下文，获取一个子Context
	////通常在Web项目中，Root Span由中间件生成
	//ctx := opentracing.ContextWithSpan(context.Background(), span)
	//session := db.WithContext(ctx)

	return db, nil
}

//创建前
func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate")
	m.CreatedOn = time.Now()
	//tx.Model(m).Updates(map[string]interface{}{"created_by": m.CreatedBy, "created_on": nowTime})
	return
}

//更新后
func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	log.Println("BeforeUpdate")
	return
}

//删除后
func (m *Model) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println("BeforeDelete")
	m.IsDel = 1
	m.DeletedOn = time.Now()
	return
}
