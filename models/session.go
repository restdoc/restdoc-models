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
