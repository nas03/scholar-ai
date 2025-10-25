package helper

import (
	"context"
	"fmt"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/resend/resend-go/v2"
)

type IMailHelper interface {
	SendMail(ctx context.Context, to, subject, body string) (string, error)
}

type MailHelper struct {
	client *resend.Client
}

func NewMailHelper() IMailHelper {
	return &MailHelper{
		client: global.Mail,
	}
}

func (h *MailHelper) SendMail(ctx context.Context, to, subject, html string) (string, error) {
	params := &resend.SendEmailRequest{
		From:    global.Config.Resend.From,
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}

	sent, err := h.client.Emails.SendWithContext(ctx, params)

	if err != nil {
		return "", fmt.Errorf("failed to send email to '%s': %w", to, err)
	}

	return sent.Id, nil
}
