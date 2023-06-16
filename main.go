package main

import (
	"fmt"
)

var (
	gitHash   = "unknown"
	goVersion = "unknown"
	version   = "1.0.0"
	buildTime = "unknown"
)

func showVersion() {
	fmt.Printf("Git Commit Hash: %s \n", gitHash)
	fmt.Printf("GoLang Version: %s \n", goVersion)
	fmt.Printf("Agent Version: %s \n", version)
	fmt.Printf("Build TimeStamp: %s \n", buildTime)
}

func showBanner() {
	// Font Name: ANSI Regular
	banner := `
 ██      ██ ████████ ██      ██ ██ ██       ██ ████     ██ ██   ██
░██     ░██░██░░░░░ ░██     ░██░██░██      ░██░██░██   ░██░██  ██ 
░██     ░██░██      ░██     ░██░██░██      ░██░██░░██  ░██░██ ██  
░░██    ██ ░███████ ░██████████░██░██      ░██░██ ░░██ ░██░████   
 ░░██  ██  ░██░░░░  ░██░░░░░░██░██░██      ░██░██  ░░██░██░██░██  
  ░░████   ░██      ░██     ░██░██░██      ░██░██   ░░████░██░░██ 
   ░░██    ░████████░██     ░██░██░████████░██░██    ░░███░██ ░░██
    ░░     ░░░░░░░░ ░░      ░░ ░░ ░░░░░░░░ ░░ ░░      ░░░ ░░   ░░
`
	fmt.Print(banner)
}

func hold() {
}

func run() {
	showBanner()
	showVersion()
	hold()
}

func main() {
	run()
}
