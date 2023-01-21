package main

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

func main() {
  // now := time.Now()
  // unix := now.Unix()
  // format := now.Format("02 Jan 2006")
  // fmt.Printf("Coucou %v", unix)
  // fmt.Printf("Coucou %v", format)
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
}

func stopAction () {
  fmt.Println("stopAction")
}

func todayAction () {
  fmt.Println("todayAction")
}

func reportAction () {
  fmt.Println("reportAction")
}

/*
** Utils
*/
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
