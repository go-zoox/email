package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Config struct {
	Host string
	Port int
	User string
	Pass string
}

type Email struct {
	From        string
	To          []string
	Subject     string
	Text        []byte
	HTML        []byte
	Attachments []string
	Cc          []string
	Bcc         []string
}

type Client struct {
	cfg *Config
	// pool *email.Pool
	server string
	auth   smtp.Auth
}

func New(cfg *Config) (*Client, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("host is required")
	}

	if cfg.Port == 0 {
		return nil, fmt.Errorf("port is required")
	}

	if cfg.User == "" {
		return nil, fmt.Errorf("user is required")
	}

	if cfg.Pass == "" {
		return nil, fmt.Errorf("pass is required")
	}

	server := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	auth := smtp.PlainAuth("", cfg.User, cfg.Pass, cfg.Host)

	// pool, err := email.NewPool(
	// 	server,
	// 	4,
	// 	auth,
	// 	// &tls.Config{ServerName: cfg.Host},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return &Client{
		cfg: cfg,
		// pool: pool,
		server: server,
		auth:   auth,
	}, nil
}

func (c *Client) Send(mail *Email) error {
	e := email.NewEmail()
	e.From = mail.From
	e.To = mail.To
	e.Subject = mail.Subject
	e.Cc = mail.Cc
	e.Bcc = mail.Bcc

	if mail.Attachments != nil {
		for _, attachment := range mail.Attachments {
			e.AttachFile(attachment)
		}
	}

	if mail.HTML != nil {
		e.HTML = mail.HTML
	} else if mail.Text != nil {
		e.Text = mail.Text
	}

	if e.From == "" {
		e.From = c.cfg.User
	}

	// pool not works with tls
	//	issue: https://github.com/jordan-wright/email/issues/77
	// return c.pool.Send(e, 10*time.Second)

	return e.SendWithTLS(c.server, c.auth, &tls.Config{ServerName: c.cfg.Host})
}
