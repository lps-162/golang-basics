package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("Go is running in Mac OS")
	case "windows":
		fmt.Println("Go is running in Windows")
	case "linux":
		fmt.Println("Go is running in Linux")
	default:
		fmt.Println("Running in some other platform")
	}
	today := time.Now().Weekday()

	fmt.Println(timeNow)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today ...")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Its far away dude")
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning !!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon !!")
	default:
		fmt.Println("Good Evening !!")
	}
}
