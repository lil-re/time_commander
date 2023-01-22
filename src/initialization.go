package src

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

/*
** Initialization
*/
var Filename string
var Start bool
var Stop bool
var Today bool
var Report int

var Command = &cobra.Command{
  Use:   "time",
  Short: "Time Commander",
  Long: "Time Commander | Generate time sheets in your terminal",
  Run: func(cmd *cobra.Command, args []string) {
      if Start {
        StartAction()
      } else if Stop {
        StopAction()
      } else if Today {
        TodayAction()
      } else if Report > 0 {
        ReportAction()
      }
  },
}

func Initialize() {
  dir, err := os.Getwd()
  HandleError(err)

  Filename = fmt.Sprintf("%v\\time_commander.json", dir)
  HandleNoFile(Filename)

  Command.PersistentFlags().BoolVarP(&Start, "start", "a", false, "Set the starting time")
  Command.PersistentFlags().BoolVarP(&Stop, "stop", "z", false, "Set the ending time")
  Command.PersistentFlags().BoolVarP(&Today, "today", "t", false, "Display a report for today")
  Command.PersistentFlags().IntVarP(&Report, "report", "r", 7, "Display a report for the last X days")
}
