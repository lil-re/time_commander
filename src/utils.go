package src

import (
	"os"
	"time"
	"io/ioutil"
	"encoding/json"
)

/*
** Utils
*/
func GetCurrentTimestamp () int64 {
	now := time.Now()
	return now.Unix()
}	

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleNoFile(filename string) {
	_, err := os.Open(filename)

	if err != nil {
		os.Create(filename)
		data := TimeCommanderData{ Records: make([]Record, 0) }
		jsonData, err := json.MarshalIndent(data, "", "  ")
		HandleError(err)
		
		_ = ioutil.WriteFile(Filename, jsonData, 0644)
	}
}
