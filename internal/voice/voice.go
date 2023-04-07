package voice

type VoiceData struct {
	Country             string  `json:"country,omitempty"`
	Bandwidth           string  `json:"bandwidth,omitempty"`
	ResponseTime        string  `json:"response_time,omitempty"`
	Provider            string  `json:"provider,omitempty"`
	ConnectionStability float32 `json:"connection_stability,omitempty"`
	TTFB                int     `json:"ttfb,omitempty"`
	VoicePurity         int     `json:"voice_purity,omitempty"`
	MedianOfCallsTime   int     `json:"median_of_calls_time,omitempty"`
}
