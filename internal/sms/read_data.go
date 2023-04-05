package sms

import (
	"bufio"
	"fmt"
	countries "github.com/biter777/countries"
	"log"
	"os"
	"strings"
)

var providers = [...]string{"Topolo", "Rond", "Kildy"}

func readFile(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
			return nil, err
		}
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return rows, nil

}

func loadValidationAlpha2Map() map[string]string {
	Alpha2Map := make(map[string]string)
	all := countries.AllInfo()
	for _, country := range all {
		Alpha2Map[country.Alpha2] = country.Name
	}
	return Alpha2Map
}

func checkValueInArray(val string, arr []string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

var Alpha2Map = loadValidationAlpha2Map()

func checkSMSvalid(data SMSData) bool {

	if Alpha2Map[data.Country] == "" {
		return false
	} else if !checkValueInArray(data.Provider, providers[:]) {
		return false
	}
	return true
}

func GetSMSDataSlice(path string) []*SMSData {
	loadValidationAlpha2Map()
	rows, err := readFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var result []*SMSData
	for _, row := range rows {
		params := strings.Split(row, ";")
		if len(params) != 4 {
			continue
		}
		sms := SMSData{
			Country:      params[0],
			Bandwith:     params[1],
			ResponseTime: params[2],
			Provider:     params[3],
		}
		if !checkSMSvalid(sms) {
			continue
		}
		result = append(result, &sms)
	}
	return result
}
