package main

import (
	"fmt"
	"time"
)

func main() {
	timeFormat := "2006-01-02"
	t, _ := time.Parse(timeFormat, "2022-02-01")
	fmt.Println(t)
	duration := time.Since(t)
	// duration := time.Now().Sub(t)
	fmt.Printf("%f", duration.Hours()/24)
}
