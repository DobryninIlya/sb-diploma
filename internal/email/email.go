package email

type EmailData struct {
	Country      string `json:"country,omitempty"`
	Provider     string `json:"provider,omitempty"`
	DeliveryTime int    `json:"delivery_time,omitempty"`
}
