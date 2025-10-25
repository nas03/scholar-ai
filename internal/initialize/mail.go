package initialize

import (
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/resend/resend-go/v2"
)

func InitMailClient() {

	client := resend.NewClient(global.Config.Resend.ApiKey)
	global.Mail = client
}
