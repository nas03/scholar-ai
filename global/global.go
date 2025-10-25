package global

import (
	"github.com/nas03/scholar-ai/backend/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Mdb    *gorm.DB
	Log    *zap.Logger
	Mail   *resend.Client
	Redis  *redis.Client
)
