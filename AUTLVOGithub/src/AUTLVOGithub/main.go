package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

type XMLTable struct {
	XMLName xml.Name `xml:"TABLE"`
	Row []struct{
		Cell string `xml:"TD"`
	}`xml:"TR"`
}

func main()  {
	/*var p io.Writer
	client := http.Client{}
	a, _ := client.Get("https://github.com/google/go-github/archive/v28.1.1.zip")
	out, _ := os.Create("")
	_, _ = io.Copy(out, a.Body)
	client.
	t,_ := ioutil.ReadAll(a.Body)
	body := &Data{}
	html
	errXML := xml.Unmarshal(t, body)
	fmt.Println(errXML)
	fmt.Println(body)

	fmt.Println(string(t))
	fmt.Println(b)

	fileUrl := "https://github.com/google/go-github/archive/v28.1.1.zip"

	if err := DownloadFile("v28.1.1.zip", fileUrl); err != nil {
		panic(err)
	}*/
	str := "MSDB01|4021"
	strSplit := strings.Split(str,"|")
	fmt.Printf("Server name: %v -> Shop: %v",strSplit[0],strSplit[1])
}
