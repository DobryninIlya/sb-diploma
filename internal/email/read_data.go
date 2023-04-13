package email

import (
	"bufio"
	"errors"
	"log"
	assert "main/internal/assertations"
	"os"
	"strconv"
	"strings"
)

const CSVRowLen = 3

func checkEmailValid(data EmailData) bool {
	if assert.Alpha2Map[data.Country] == "" {
		return false
	} else if !assert.CheckValueInArray(data.Provider, assert.EmailProviders[:]) {
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

func getEmailData(params []string) (EmailData, error) {
	DeliveryTime, err := strconv.Atoi(params[2])
	if err != nil {
		return EmailData{}, errors.New("bad data format")
	}
	return EmailData{
		Country:      params[0],
		Provider:     params[1],
		DeliveryTime: DeliveryTime,
	}, nil

}

func GetEmailDataSlice(path string) ([]EmailData, error) {
	rows, err := readFile(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var result []EmailData
	for _, row := range rows {
		params := strings.Split(row, ";")
		if len(params) != CSVRowLen {
			continue
		}
		email, err := getEmailData(params)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		if !checkEmailValid(email) {
			continue
		}
		result = append(result, email)
	}
	return result, nil
}
