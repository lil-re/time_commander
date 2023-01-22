package main

import (
  "./src"
)

func main() {
  src.InitializeCommands()
  src.Commands.Execute()
}
