package src

import "fmt"

func SessionHasBeenStarted() {
	fmt.Println("Session has been started")
}

func SessionHasAlreadyBeenStarted() {
	fmt.Println("Session has already been started")
}

func SessionHasBeenStopped() {
	fmt.Println("Session has been stopped")
}

func SessionIsAlreadyStopped() {
	fmt.Println("Session is already stopped")
}

func NoRecord() {
	fmt.Println("There is no record")
}

func NoRecordToday() {
	fmt.Println("There is no record today")
}
