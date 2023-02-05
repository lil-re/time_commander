package main

import (
  "github.com/lil-re/time_commander/src"
)

func main() {
  src.InitializeCommands()
  src.Commands.Execute()
}
