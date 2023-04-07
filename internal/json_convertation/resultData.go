package json_convertation

import (
	assert "main/internal/assertations"
	"main/internal/billing"
	"main/internal/email"
	incidentData "main/internal/incedent_data"
	"main/internal/mms"
	"main/internal/sms"
	"main/internal/support"
	"main/internal/voice"
	"path/filepath"
	"time"
)

const speedSupport = 18

type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voice.VoiceData              `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []incidentData.IncidentData    `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data,omitempty""`
	Error  string     `json:"error,omitempty"`
}
type ResultTErr struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}

const gate = "http://127.0.0.1:8383"

type msgData interface {
	SetCountry(new string)
	GetCountry() string
	GetProvider() string
}

func sortMsgByDelivery(list []email.EmailData) []email.EmailData {
	// Bubble sort :)
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if list[j].DeliveryTime > list[j+1].DeliveryTime {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func sortMsgByProvider(list []msgData) []msgData {
	// Bubble sort :)
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if list[j].GetProvider() > list[j+1].GetProvider() {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func sortMsgByCountry(list []msgData) []msgData {
	// Bubble sort v.2 :)
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := (length - 1); j > i; j-- {
			if list[j].GetCountry() < list[j-1].GetCountry() {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return list
}

func replaceCodeOnNameSMS(list []sms.SMSData) []sms.SMSData {
	result := make([]sms.SMSData, 0, len(list))
	for _, data := range list {
		data.Country = assert.Alpha2Map[data.Country]
		result = append(result, data)
	}
	return result
}
func replaceCodeOnNameMMS(list []mms.MMSData) []mms.MMSData {
	result := make([]mms.MMSData, 0, len(list))
	for _, data := range list {
		data.Country = assert.Alpha2Map[data.Country]
		result = append(result, data)
	}
	return result
}

// В этой функции реализована подготовка к сортировке двух слайсов []SMSData и их последующая
// соритровка в универсальной функции
func getSMSdata(list []sms.SMSData) [][]sms.SMSData {
	resultSMS := replaceCodeOnNameSMS(list)
	//var genericMessages []msgData
	genericMessages := make([]msgData, 0, len(resultSMS))
	for i, _ := range resultSMS {
		genericMessages = append(genericMessages, &resultSMS[i])
	}
	sortedSMSByProvider := sortMsgByProvider(genericMessages)
	sliceSMS1 := make([]sms.SMSData, len(sortedSMSByProvider))
	for i, message := range sortedSMSByProvider { // Приводим весь слайс []msgData к []SMSData
		smsMessage, ok := message.(*sms.SMSData)
		if !ok {
		}
		sliceSMS1[i] = *smsMessage
	}
	sliceSMS2 := make([]sms.SMSData, len(sortedSMSByProvider))
	sortedSMSByCountry := sortMsgByCountry(genericMessages)
	for i, message := range sortedSMSByCountry { // Приводим весь слайс []msgData к []SMSData
		smsMessage, ok := message.(*sms.SMSData)
		if !ok {
		}
		sliceSMS2[i] = *smsMessage
	}
	return [][]sms.SMSData{
		sliceSMS1, // second slice
		sliceSMS2, //first slice
	}

}

// В этой функции реализована подготовка к сортировке двух слайсов []MMSData и их последующая
// соритровка в универсальной функции
func getMMSdata(list []mms.MMSData) [][]mms.MMSData {
	resultMMS := replaceCodeOnNameMMS(list)
	//var genericMessages []msgData
	genericMessages := make([]msgData, 0, len(resultMMS))
	for i, _ := range resultMMS {
		genericMessages = append(genericMessages, &resultMMS[i])
	}
	sortedSMSByProvider := sortMsgByProvider(genericMessages)
	sliceMMS1 := make([]mms.MMSData, len(sortedSMSByProvider))
	for i, message := range sortedSMSByProvider { // Приводим весь слайс []msgData к []SMSData
		smsMessage, ok := message.(*mms.MMSData)
		if !ok {
		}
		sliceMMS1[i] = *smsMessage
	}
	sliceMMS2 := make([]mms.MMSData, len(sortedSMSByProvider))
	sortedSMSByCountry := sortMsgByCountry(genericMessages)
	for i, message := range sortedSMSByCountry { // Приводим весь слайс []msgData к []SMSData
		smsMessage, ok := message.(*mms.MMSData)
		if !ok {
		}
		sliceMMS2[i] = *smsMessage
	}
	return [][]mms.MMSData{
		sliceMMS1, // second slice
		sliceMMS2, //first slice
	}

}

func getMinMax3email(list []email.EmailData) (result [][]email.EmailData) {
	resultSort := sortMsgByDelivery(list)
	min3 := make([]email.EmailData, 3)
	max3 := make([]email.EmailData, 3)
	min3 = resultSort[:3]
	max3 = resultSort[len(resultSort)-3:]
	result = [][]email.EmailData{
		min3,
		max3,
	}
	return result
}

func getEmailData(list []email.EmailData) map[string][][]email.EmailData {
	//func getEmailData(list []email.EmailData) map[string][][]email.EmailData {
	result := make(map[string][][]email.EmailData)
	CountryProvider := make(map[string][]email.EmailData)
	for _, data := range list {
		CountryProvider[data.Country] = append(CountryProvider[data.Country], data)
	}
	//CountryProviderBest3 := make(map[string][]string)
	for country, provider := range CountryProvider {
		sorted := getMinMax3email(provider)
		result[country] = sorted
	}
	return result
}

func getSupportData(list []support.SupportData) []int {
	result := make([]int, 2)
	hour := time.Now().Hour()
	if hour < 9 {
		result[0] = 1
	} else if hour > 16 {
		result[0] = 3
	} else {
		result[0] = 2
	}
	waitTime := func() int {
		result := 0
		for _, data := range list {
			result += data.ActiveTickets
		}
		return result * speedSupport
	}()
	result[1] = waitTime
	return result

}

func getIncidentData(list []incidentData.IncidentData) []incidentData.IncidentData {
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if list[j+1].Status == "active" && list[j].Status != "active" {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func GetResultData() (ResultSetT, error) {
	resultSMS, errSMS := sms.GetSMSDataSlice(filepath.Join("internal", "data", "sms.data"))
	if errSMS != nil {
		return ResultSetT{}, errSMS
	}
	resultMMS, errMMS := mms.GetMMS(gate + "/mms")
	if errMMS != nil {
		return ResultSetT{}, errSMS
	}
	resultVoice, errVoice := voice.GetVoiceDataSlice(filepath.Join("internal", "data", "voice.data"))
	if errVoice != nil {
		return ResultSetT{}, errSMS
	}
	resultEmail, errEmail := email.GetEmailDataSlice(filepath.Join("internal", "data", "email.data"))
	if errEmail != nil {
		return ResultSetT{}, errSMS
	}
	resultBilling, errBilling := billing.GetBillingData(filepath.Join("internal", "data", "billing.data"))
	if errBilling != nil {
		return ResultSetT{}, errSMS
	}
	resultSupport, errSupport := support.GetSupport(gate + "/support")
	if errSupport != nil {
		return ResultSetT{}, errSMS
	}
	resultIncident, errIncident := incidentData.GetIncident(gate + "/accendent")
	if errIncident != nil {
		return ResultSetT{}, errSMS
	}
	dataSms := getSMSdata(resultSMS)
	dataMMS := getMMSdata(resultMMS)
	dataEmail := getEmailData(resultEmail)
	dataSupport := getSupportData(resultSupport)
	dataIncidents := getIncidentData(resultIncident)
	//return ResultSetT{}, nil
	return ResultSetT{
		SMS:       dataSms,
		MMS:       dataMMS,
		VoiceCall: resultVoice,
		Email:     dataEmail,
		Billing:   resultBilling,
		Support:   dataSupport,
		Incidents: dataIncidents,
	}, nil

}
