package utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SmtpClient struct {
	clientAuth smtp.Auth
}

func (c *SmtpClient) GetAuthFromGoogle(account, password string) {
	c.getAuth("foodiePath", account, password, "smtp.gmail.com")
}

func (c *SmtpClient) getAuth(identity, account, password, smtpHost string) {
	c.clientAuth = smtp.PlainAuth(identity, account, password, smtpHost)
}

func (c *SmtpClient) SendEmailFromGoogle(subject, content, from string, to []string) {
	c.sendEmail(subject, content, from, to, "smtp.gmail.com", "587")
}

func (c *SmtpClient) sendEmail(subject, content, from string, to []string, smtpHost, smtpPort string) {
	// Authentication.
	if c.clientAuth == nil {
		fmt.Println("no auth found. please get auth before sending email. ")
		return
	}

	// Message.
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		content + "\r\n")

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, c.clientAuth, from, to, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func MuskAnEmail(rawEmail *string) string {
	if rawEmail == nil {
		return "" // no content in the string
	}
	res := strings.Clone(*rawEmail)
	atPos := strings.Index(res, "@")
	domainPos := strings.LastIndex(res, ".")
	if atPos == -1 || domainPos == -1 {
		return "" // for cases that are not email
	}
	return res[0:1] + strings.Repeat("*", len(res[1:atPos])) + "@" + strings.Repeat("*", len(res[atPos+1:domainPos])) + res[domainPos:]
}

func Btoi(mybool bool) int {
	if mybool {
		return 1
	} else {
		return 0
	}
}
