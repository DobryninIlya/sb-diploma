package assertations

import (
	"github.com/biter777/countries"
)

var Alpha2Map = loadValidationAlpha2Map()
var Providers = [...]string{"Topolo", "Rond", "Kildy"}
var VoiceProviders = [...]string{"TransparentCalls", "E-Voice", "JustPhone"}
var EmailProviders = [...]string{"Gmail", "Gmail", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail",
	"Yandex", "Mail.ru"}

func loadValidationAlpha2Map() map[string]string {
	Alpha2Map := make(map[string]string)
	all := countries.AllInfo()
	for _, country := range all {
		Alpha2Map[country.Alpha2] = country.Name
	}
	return Alpha2Map
}

func CheckValueInArray(val string, arr []string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
