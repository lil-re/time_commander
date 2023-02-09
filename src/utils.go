package src

import (
	"time"
)

/*
** Utils
 */
func GetCurrentTimestamp() int64 {
	now := time.Now()
	return now.Unix()
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
