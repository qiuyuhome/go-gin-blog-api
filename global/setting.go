package global

import (
	"github.com/qiuyuhome/go-gin-blog-api/pkg/logger"
	"github.com/qiuyuhome/go-gin-blog-api/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
