package src

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
** Commands
 */
var Filename string
var Start bool
var Stop bool
var Today bool
var Report int

var Commands = &cobra.Command{
	Use:   "time",
	Short: "Time Commander",
	Long:  "Time Commander | Generate time sheets in your terminal",
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

func InitializeCommands() {
	dir, err := os.Getwd()
	HandleError(err)

	Filename = fmt.Sprintf("%v/time_commander.json", dir)
	HandleNoFile(Filename)

	Commands.PersistentFlags().BoolVarP(&Start, "start", "s", false, "Start a new session")
	Commands.PersistentFlags().BoolVarP(&Stop, "finish", "f", false, "Finish the current session")
	Commands.PersistentFlags().BoolVarP(&Today, "today", "t", false, "Display a report for today")
	Commands.PersistentFlags().IntVarP(&Report, "report", "r", 7, "Display a report for the last X days")
}
