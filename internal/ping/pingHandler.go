package ping

import (
	"encoding/json"
	"fmt"
	conv "main/internal/json_convertation"
	"net/http"
)

func New() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ping handler")
		structura, err := conv.GetResultData()
		if err != nil {
			answer := conv.ResultTErr{
				Status: false,
				Error:  err.Error(),
			}
			result, _ := json.Marshal(answer)
			w.Write(result)
			return
		}

		answer := conv.ResultT{
			Status: true,
			Data:   structura,
		}

		result, _ := json.Marshal(answer)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(result)
	}
}
