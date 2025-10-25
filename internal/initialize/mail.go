package initialize

import (
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/resend/resend-go/v2"
	"go.uber.org/zap"
)

func InitMailClient() {

	client := resend.NewClient(global.Config.Resend.ApiKey)
	global.Log.Info("Mail client established successfully",
		zap.String("provider", "resend"),
		zap.String("from_email", global.Config.Resend.From),
	)
	global.Mail = client

}
