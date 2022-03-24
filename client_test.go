package email

import (
	"strconv"
	"testing"

	"github.com/go-zoox/dotenv"
)

func TestEmail(t *testing.T) {
	dotenv.Load()
	Port, _ := strconv.Atoi(dotenv.Get("SMTP_PORT"))

	client, err := New(&Config{
		Host: dotenv.Get("SMTP_HOST"),
		Port: Port,
		User: dotenv.Get("SMTP_USER"),
		Pass: dotenv.Get("SMTP_PASS"),
	})
	if err != nil {
		t.Error(err)
	}

	err = client.Send(&Email{
		To:      []string{dotenv.Get("SMTP_TO_1")},
		Subject: "Test Email By Go Email",
		Text:    []byte("This is a test email sent by Go Email"),
	})
	if err != nil {
		t.Error(err)
	}

	err = client.Send(&Email{
		To:      []string{dotenv.Get("SMTP_TO_2")},
		Subject: "Test Email By Go Email",
		Text:    []byte("This is a test email sent by Go Email"),
	})
	if err != nil {
		t.Error(err)
	}
}
