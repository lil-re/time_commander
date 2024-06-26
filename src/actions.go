package src

import (
	"fmt"
	"time"
)

/*
** Actions
 */
func StartAction() {
	records := GetFileData()
	now := time.Now()
	date := now.Format("2006-01-02")

	if len(records) == 0 {
		createNewRecord(&records, date)
	} else {
		record := &records[len(records)-1]

		if date != record.Date {
			createNewRecord(&records, date)
		} else if len(record.Sessions) > 0 {
			session := &record.Sessions[len(record.Sessions)-1]

			if session.End == 0 {
				SessionHasAlreadyBeenStarted()
			} else {
				createNewSession(record)
			}
		}
	}
	SetFileData(records)
}

func StopAction() {
	records := GetFileData()

	if len(records) > 0 {
		record := &records[len(records)-1]

		if len(record.Sessions) > 0 {
			session := &record.Sessions[len(record.Sessions)-1]

			if session.End == 0 {
				session.End = GetCurrentTimestamp()
				SetFileData(records)
				SessionHasBeenStopped()
				TodayAction()
			} else {
				SessionIsAlreadyStopped()
			}
		}
	} else {
		NoRecord()
	}
}

func TodayAction() {
	records := GetFileData()

	if len(records) > 0 {
		record := records[len(records)-1]
		now := time.Now()
		date := now.Format("2006-01-02")

		if date == record.Date {
			printRecord(record, "Today")
		} else {
			NoRecordToday()
		}
	}
}

func ReportAction() {
	records := GetFileData()
	recordCounter := len(records)
	recordLimit := -1

	if Report <= recordCounter {
		recordLimit = recordCounter - Report - 1
	}

	if recordCounter > 0 {
		for i := recordCounter - 1; i > recordLimit; i-- {
			printRecord(records[i], records[i].Date)
		}
	} else {
		NoRecord()
	}
}

func createNewRecord(records *[]Record, date string) {
	newSessions := make([]Session, 1)
	newSessions[0].Start = GetCurrentTimestamp()
	newRecord := Record{
		Date:     date,
		Sessions: newSessions,
	}
	*records = append(*records, newRecord)
	SessionHasBeenStarted()
}

func createNewSession(record *Record) {
	newSession := Session{
		Start: GetCurrentTimestamp(),
		End:   0,
	}
	record.Sessions = append(record.Sessions, newSession)
	SessionHasBeenStarted()
}

func getSessionDuration(session Session) float64 {
	start := time.Unix(session.Start, 0)
	var end time.Time

	if session.End == 0 {
		end = time.Unix(GetCurrentTimestamp(), 0)
	} else {
		end = time.Unix(session.End, 0)
	}

	return end.Sub(start).Seconds()
}

func printRecord(record Record, date string) {
	floatDuration := 0.0

	for j := 0; j < len(record.Sessions); j++ {
		floatDuration += getSessionDuration(record.Sessions[j])
	}

	textDuration := fmt.Sprintf("%vs", floatDuration)
	parsedDuration, _ := time.ParseDuration(textDuration)
	fmt.Printf("%v => %v\n", date, parsedDuration)
}
