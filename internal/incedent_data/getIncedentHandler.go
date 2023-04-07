package incidentData

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetIncident(path string) ([]IncidentData, error) {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
		log.Fatal("error has occured, when http-get response sended on ", path)
		return []IncidentData{}, errors.New("error has occured, when http-get response sended on ")
	}
	if resp.StatusCode != 200 {
		log.Fatal("Status code is not 200, error is occured")
		var list []IncidentData
		return list, errors.New("Status code is not 200, error is occured")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return getIncidentStruct(body), nil

}

func getIncidentStruct(body []byte) []IncidentData {
	var list []IncidentData
	err := json.Unmarshal(body, &list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
	return list
}
