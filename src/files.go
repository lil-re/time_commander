package src


import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)


/*
** Files
*/
func GetFileData () []Record {
	jsonFile, err := os.Open(Filename)
	HandleError(err)
  
	defer jsonFile.Close()
  
	byteValue, err := ioutil.ReadAll(jsonFile)
	HandleError(err)
  
	var data TimeCommanderData
	json.Unmarshal(byteValue, &data)
	fmt.Printf("%v \n", data)
	return data.Records
}

  func SetFileData (records []Record) {
	data := TimeCommanderData{ Records: records }
	jsonData, err := json.MarshalIndent(data, "", "  ")
	HandleError(err)
	
	_ = ioutil.WriteFile(Filename, jsonData, 0644)
}
  