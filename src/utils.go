package src

import (
	"os"
	"time"
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
	}
}
