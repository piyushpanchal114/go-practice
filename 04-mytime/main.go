package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study...")
	presentTime := time.Now()
	fmt.Println("the present time", presentTime)

	fmt.Println("Formatted time", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.February, 19, 10, 30, 00, 0, time.Local)
	fmt.Println("created Date", createdDate)
	fmt.Println("Formatted Created Date", createdDate.Format("01-02-2006 Monday 15:04:05"))
}
