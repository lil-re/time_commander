package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
** Files
 */
func GetFileData() []Record {
	jsonFile, err := os.Open(Filename)
	HandleError(err)

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	HandleError(err)

	var data TimeCommanderData
	json.Unmarshal(byteValue, &data)
	return data.Records
}

func SetFileData(records []Record) {
	data := TimeCommanderData{Records: records}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	HandleError(err)

	_ = ioutil.WriteFile(Filename, jsonData, 0644)
}
