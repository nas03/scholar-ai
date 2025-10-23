package global

import (
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Mdb  *gorm.DB
	Log  *zap.Logger
	Mail *mail.Client
)
