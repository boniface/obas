package domain

import "time"

type EmailMessage struct {
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

type MailApi struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Sender string `json:"sender"`
}

type MailConfig struct {
	Id    string    `json:"id"`
	Key   string    `json:"key"`
	Value string    `json:"value"`
	Host  string    `json:"host"`
	Port  string    `json:"port"`
	State string    `json:"state"`
	Date  time.Time `json:"date"`
}

type MailEvents struct {
	MAIL             string `json:"mail"`
	MAIL_SENT        string `json:"mail_sent"`
	MAIL_SENT_FAILED string `json:"mail_sent_failed"`
	PASSWORD_RESET   string `json:"password_reset"`
	NO_PERSON_FOUND  string `json:"no_person_found"`
}

type MessageResponse struct {
	StatusCode int    `json:"statusCode"`
	Headers    string `json:"headers"`
	Body       string
}

type SmtpConfig struct {
	Id       string `json:"id"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}
