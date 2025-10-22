package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Mdb *gorm.DB
	Log *zap.Logger
)
