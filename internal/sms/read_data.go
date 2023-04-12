package sms

import (
	"bufio"
	"log"
	assert "main/internal/assertations"
	"os"
	"strings"
)

const CSVRowLen = 4

func CheckSMSvalid(data SMSData) bool {

	if assert.Alpha2Map[data.Country] == "" {
		return false
	} else if !assert.CheckValueInArray(data.Provider, assert.Providers[:]) {
		return false
	}
	return true
}

func readFile(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("file does not exist")
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

func GetSMSDataSlice(path string) ([]SMSData, error) {
	rows, err := readFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var result []SMSData
	for _, row := range rows {
		params := strings.Split(row, ";")
		if len(params) != CSVRowLen {
			continue
		}
		sms := SMSData{
			Country:      params[0],
			Bandwith:     params[1],
			ResponseTime: params[2],
			Provider:     params[3],
		}
		if !CheckSMSvalid(sms) {
			continue
		}
		result = append(result, sms)
	}
	return result, nil
}
