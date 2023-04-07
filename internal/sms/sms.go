package sms

type SMSData struct {
	Country      string `json:"country,omitempty"`
	Bandwith     string `json:"bandwith,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
	Provider     string `json:"provider,omitempty"`
}

func (s *SMSData) SetCountry(new string) {
	s.Country = new
}
func (s *SMSData) GetCountry() string {
	return s.Country
}
func (s *SMSData) GetProvider() string {
	return s.Provider
}
