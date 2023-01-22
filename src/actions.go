package src

import (
	"fmt"
	"time"
)

/*
** Actions
*/
func StartAction () {
	records := GetFileData()
  
	if (len(records) > 0) {
	  record := &records[len(records) - 1]
  
	  now := time.Now()
	  date := now.Format("2006-01-02")
  
	  if (date != record.Date) {
		newSessions := make([]Session, 1)
		newSessions[0].Start = GetCurrentTimestamp()
		newRecord := Record{
		  Date: date,
		  Sessions: newSessions,
		}
		records = append(records, newRecord)
		fmt.Println("Session has been started")
	  } else if len(record.Sessions) > 0 {
		session := &record.Sessions[len(record.Sessions) - 1]
  
		if session.End == 0  {
		  fmt.Println("Session has already been started")
		} else {
		  newSession := Session{
			Start: GetCurrentTimestamp(),
			End: 0,
		  }
		  record.Sessions = append(record.Sessions, newSession)
		  fmt.Println("Session has been started")
		}
	  }
	  SetFileData(records)
	}
}
  
func StopAction () {
	records := GetFileData()
  
	if (len(records) > 0) {
	  record := &records[len(records) - 1]
		
	  if len(record.Sessions) > 0 {
		session := &record.Sessions[len(record.Sessions) - 1]
  
		if session.End == 0  {
		  session.End = GetCurrentTimestamp()
		  SetFileData(records)
		  fmt.Println("Session has been stopped")
		} else {
		  fmt.Println("Session is already stopped")
		}
	  }
	}
}
  
func TodayAction () {
	records := GetFileData()
  
	if (len(records) > 0) {
	  record := records[len(records) - 1]
  
	  now := time.Now()
	  date := now.Format("2006-01-02")
  
	  if (date == record.Date) {
		floatDuration := 0.0
  
		for i := 0; i < len(record.Sessions); i++ {
		  session := record.Sessions[i]
		  start := time.Unix(session.Start, 0)
		  end := time.Unix(session.End, 0)
		  floatDuration = floatDuration + end.Sub(start).Seconds()
		}
  
		textDuration := fmt.Sprintf("%vs", floatDuration)
		parsedDuration, _ := time.ParseDuration(textDuration)
		fmt.Printf("Today => %v\n", parsedDuration)
	  } else {
		fmt.Println("There is no Record today")
	  }
	}
}
  
func ReportAction () {
	records := GetFileData()
	recordCounter := len(records)
	recordLimit := recordCounter - Report - 1
  
	if (recordCounter > 0) {
  
	  for i := recordCounter - 1; i > recordLimit; i-- {
		record := records[i]
		floatDuration := 0.0
  
		for j := 0; j < len(record.Sessions); j++ {
		  session := record.Sessions[j]
		  start := time.Unix(session.Start, 0)
		  end := time.Unix(session.End, 0)
		  floatDuration = floatDuration + end.Sub(start).Seconds()
		}
		textDuration := fmt.Sprintf("%vs", floatDuration)
		parsedDuration, _ := time.ParseDuration(textDuration)
		fmt.Printf("%v => %v\n", record.Date, parsedDuration)
	  }
	} else {
	  fmt.Println("There is no Record")
	}
}