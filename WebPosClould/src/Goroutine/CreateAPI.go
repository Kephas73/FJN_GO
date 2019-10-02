package Goroutine

import "time"

func CreateAPI() []RequestAPI {
	var requestAPI []RequestAPI
	for i := 0; i < 100; i++ {
		item := RequestAPI{
			NAME,
			NUM,
		}
		requestAPI = append(requestAPI, item)
	}
	return requestAPI
}

func SendAPI(data RequestAPI) RequestAPI {
	time.Sleep(10000*5)
	return data
}

// Create chanel
func Chanel() []RequestAPI {
	var requestAPI []RequestAPI
	data := CreateAPI()
	ch := make(chan RequestAPI, len(data))
	for _,v := range data{
		go func(item RequestAPI) {
			k := SendAPI(item)
			ch <- k
		}(v)
	}

	for range data{
		item := <- ch
		requestAPI = append(requestAPI, item)
	}
	return requestAPI
}
