package mms

type MMSData struct {
	Country      string `json:"country,omitempty"`
	Provider     string `json:"provider,omitempty"`
	Bandwidth    string `json:"bandwidth,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
}

func (s *MMSData) SetCountry(new string) {
	s.Country = new
}
func (s *MMSData) GetCountry() string {
	return s.Country
}
func (s *MMSData) GetProvider() string {
	return s.Provider
}
