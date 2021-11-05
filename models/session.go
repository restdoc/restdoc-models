package Models

type Session struct {
	Id            string                  `json:"id"`
	User          string                  `json:"user"`
	Email         string                  `json:"email"`
	Valid         string                  `json:"valid"`
	Admin         string                  `json:"admin"`
	Login         string                  `json:"login"`
	DomainId      string                  `json:"domain_id"`
	Locale        string                  `json:"locale"`
	Labels        map[string]Label        `json:"labels"`
	ContactLabels map[string]ContactLabel `json:"contact_labels"`
}

type SendSession struct {
	DomainId          uint64 `json:"domain_id"`
	UserId            uint64 `json:"user_id"`
	SmtpServer        string `json:"smtp_server"`
	Type              uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	ForeignSmtpServer string `json:"foreign_smtp_server"`
	PrivateKey        string `json:"private_key"`
	Amount            uint32 `json:"amount"`
}
