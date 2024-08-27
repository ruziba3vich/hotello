package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"math/rand"

	gomail "gopkg.in/gomail.v2"
)

func (s *UsersStorage) generateVerificationCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(90000) + 10000
}

func (s *UsersStorage) sendVerificationCode(ctx context.Context, email string) error {
	code := s.generateVerificationCode()

	m := gomail.NewMessage()
	m.SetHeader("From", "hotello@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/plain", fmt.Sprintf("Your verification code is: %d", code))

	d := gomail.NewDialer(s.config.GetSmptpHost(), s.config.GetSmptpPort(), s.config.GetSmptpUser(), s.config.GetSmptpPass())

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return s.redisService.StoreEmailAndCode(ctx, email, code)
}

func (s *UsersStorage) verifyEmail(ctx context.Context, email string, code int) error {
	c, err := s.redisService.GetCodeByEmail(ctx, email)
	if err != nil {
		return err
	}
	if c != code {
		return errors.New("invalide code")
	}
	return nil
}
