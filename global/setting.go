package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JwtS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
