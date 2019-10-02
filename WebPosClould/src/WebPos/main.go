package main

import (
	"Retry"
	"fmt"
	"time"
)

func main()  {
	fmt.Println(time.Now().Format(Retry.FORMAT_TIME),"================Process Start================")
	defer fmt.Println(time.Now().Format(Retry.FORMAT_TIME),"================Process End================")
	Retry.Retry()
}
