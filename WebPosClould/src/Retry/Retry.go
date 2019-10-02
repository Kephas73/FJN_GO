package Retry

import (
	"fmt"
	"math/rand"
	"time"
)

func Count(index int) int {
	return index
}

func RandomInt32() int {
	return rand.Intn(INPUT) + rand.Intn(INPUT)
}

func Retry()  {
	retryCount := 0
	for  {
		index := func() int{
			return Count(RandomInt32())
		}()

		if index == rand.Intn(OUTPUT) {
			fmt.Print(time.Now().Format(FORMAT_TIME))
			fmt.Printf("[%v] = [%v]\n",index,rand.Intn(OUTPUT))
			fmt.Println(time.Now().Format(FORMAT_TIME),"Processing success!")
			fmt.Println("-------------------")
			break
		} else {
			fmt.Print(time.Now().Format(FORMAT_TIME))
			fmt.Printf("[%v] <> [%v]\n",index,rand.Intn(OUTPUT))
			fmt.Println(time.Now().Format(FORMAT_TIME),"Processing failed!")
			fmt.Println("-------------------")
			if retryCount < MAX_RETRY_DEFAULT {
				retryCount++
				time.Sleep(time.Second * TIME_WATTING)
				fmt.Print(time.Now().Format(FORMAT_TIME))
				fmt.Printf("Processing retry %v\n",retryCount)
				time.Sleep(time.Second)
			} else {
				break
			}
		}
	}
}
