package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	//"github.com/yasir2000/go-datastructures/threadsafe/err"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no commnad is provided")
		os.Exit(2)
	}
	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDERS CLI", "the message for the user")
		if *msgFlag == "msg" {
			err := greetCmd.Parse(os.Args[2:])
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("hello and welcome: %s", *msgFlag)
		} else {
			fmt.Printf("Wrong flag, please use flag --msg")
		}
	case "help":
		fmt.Println("some help message")
	default:
		fmt.Printf("unkown command: %s\n", cmd)
	}
}
