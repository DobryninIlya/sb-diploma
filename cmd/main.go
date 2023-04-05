package main

import (
	"fmt"
	"main/internal/sms"
	"path/filepath"
)

func main() {
	result := sms.GetSMSDataSlice(filepath.Join("internal", "data", "sms.data"))
	fmt.Println(result)
}
