package global

import (
	"github.com/nas03/scholar-ai/backend/pkg/setting"
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Mdb    *gorm.DB
	Log    *zap.Logger
	Mail   *mail.Client
)
