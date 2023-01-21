package main

import (
  "os"
  "fmt"
  "time"
  "io/ioutil"
  "encoding/json"
  "github.com/spf13/cobra"
)

func main() {
  Execute()
}

/*
** Initialization
*/
var filename string
var start bool
var stop bool
var today bool
var report int

var rootCmd = &cobra.Command{
  Use:   "time",
  Short: "Time Commander",
  Long: "Time Commander | Generate time sheets in your terminal",
  Run: func(cmd *cobra.Command, args []string) {
      if start {
        startAction()
      } else if stop {
        stopAction()
      } else if today {
        todayAction()
      } else if report > 0 {
        reportAction()
      }
  },
}

func init() {
  dir, err := os.Getwd()
  handleError(err)

  filename = fmt.Sprintf("%v\\time.txt", dir)
  handleNoFile()

  rootCmd.PersistentFlags().BoolVarP(&start, "start", "a", false, "Set the starting time")
  rootCmd.PersistentFlags().BoolVarP(&stop, "stop", "z", false, "Set the ending time")
  rootCmd.PersistentFlags().BoolVarP(&today, "today", "t", false, "Display a report for today")
  rootCmd.PersistentFlags().IntVarP(&report, "report", "r", 7, "Display a report for the last X days")

  //initializeList()
}

func Execute() {
  rootCmd.Execute()
}

/*
** Actions
*/
func startAction () {
  fmt.Println("startAction")
  records := getRecords()

  if (len(records) > 0) {
    record := &records[len(records) - 1]

    now := time.Now()
    date := now.Format("2006-01-02")

    if (date != record.Date) {
      newSessions := make([]Session, 1)
      newSessions[0].Start = getCurrentTimestamp()
      newRecord := Record{
        Date: date,
        Sessions: newSessions,
      }
      records = append(records, newRecord)
    } else if len(record.Sessions) > 0 {
      session := &record.Sessions[len(record.Sessions) - 1]

      if session.End == 0  {
        fmt.Println("Session has already started")
      } else {
        newSession := Session{
          Start: getCurrentTimestamp(),
          End: 0,
        }
        record.Sessions = append(record.Sessions, newSession)
      }
    }
    setRecords(records)
  }
}

func stopAction () {
  fmt.Println("stopAction")
  records := getRecords()

  if (len(records) > 0) {
    record := &records[len(records) - 1]
      
    if len(record.Sessions) > 0 {
      session := &record.Sessions[len(record.Sessions) - 1]

      if session.End == 0  {
        session.End = getCurrentTimestamp()
        setRecords(records)
      } else {
        fmt.Println("Session is already stopped")
      }
    }
  }
}

func todayAction () {
  fmt.Println("todayAction")
}

func reportAction () {
  fmt.Println("reportAction")
}

/*
** Structures
*/
type TimeCommanderData struct {
  Records []Record `json:"records"`
}

type Record struct {
	Date     string    `json:"date"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

/*
** Files
*/
func getRecords () []Record {
  jsonFile, err := os.Open("time.json")
  handleError(err)

  defer jsonFile.Close()

  byteValue, err := ioutil.ReadAll(jsonFile)
  handleError(err)

  var data TimeCommanderData
  json.Unmarshal(byteValue, &data)
  return data.Records
}

func setRecords (records []Record) {
  data := TimeCommanderData{ Records: records }
  jsonData, err := json.MarshalIndent(data, "", "  ")
  handleError(err)
  
  _ = ioutil.WriteFile("time.json", jsonData, 0644)
}

/*
** Utils
*/
func getCurrentTimestamp () int64 {
  now := time.Now()
  return now.Unix()
}

func handleError(err error) {
  if err != nil {
      panic(err)
  }
}

func handleNoFile() {
  _, err := os.Open(filename)
  if err != nil {
      os.Create(filename)
  }
}
