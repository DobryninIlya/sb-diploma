package billing

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const ByteMaskLen = 6

func readFile(path string) ([]byte, error) {
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
	data := make([]byte, 6)
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}
	return data, nil
}

func GetBillingData(path string) (BillingData, error) {
	buff, err := readFile(path)
	if err != nil {
		log.Fatal(err)
		return BillingData{}, err
	}
	if len(buff) != ByteMaskLen {
		log.Printf("Bad byte mask")
		return BillingData{}, errors.New("Bad byte mask")
	}
	return BillingData{
		CreateCustomer: buff[5] == '1',
		Purchase:       buff[4] == '1',
		Payout:         buff[3] == '1',
		Recurring:      buff[2] == '1',
		FraudControl:   buff[1] == '1',
		CheckoutPage:   buff[0] == '1',
	}, nil
}
