package support

type SupportData struct {
	Topic         string `json:"topic,omitempty"`
	ActiveTickets int    `json:"active_tickets"`
}
