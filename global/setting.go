package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JwtS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	Tracer          opentracing.Tracer
)
