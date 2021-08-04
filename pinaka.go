package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/jayateertha043/pinaka/pkg/httpclient"
)

func main() {
	//
	URL := flag.String("url", "", "Input Url(https://www.example.com)")
	THREADS := flag.Int("t", 100, "Enter amount of threads")
	MAXREQUEST := flag.Uint64("maxrequest", 100000, "Enter max requests to make")
	TIMEOUT := flag.Int("timeout", 3, "Enter request timeout in seconds")
	POST := flag.Bool("post", false, "To send post request")
	DATA := flag.String("data", "", "To send custom post data)")
	PROXY := flag.String("proxy", "", "Use custom proxy [http://ip:port or https://ip:port]")
	headersF := flag.String("headers", "", "To use Custom Headers headers.json file")
	flag.Parse()
	printBanner()
	if *URL == "" {
		fmt.Println("No URL supplied ,exiting...")
		return
	}
	if *THREADS < 1 {
		fmt.Println("Please supply valid number of threads, exiting...")
		return
	}
	if *MAXREQUEST < 1 {
		fmt.Println("Please supply valid number of max requests, exiting...")
		return
	}

	headers := make(map[string]string)
	*headersF = strings.TrimSpace(*headersF)
	if *headersF != "" {
		jsonFile, err := os.Open(*headersF)
		if err != nil {
			fmt.Println("Unable to find " + *headersF)
			return
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("Unable to read")
			return
		}
		err = json.Unmarshal(byteValue, &headers)
		if err != nil {
			fmt.Println("Json format invalid in headers.json")
			return
		}
	}

	sem := make(chan bool, *THREADS)

	u, err := url.Parse(*URL)
	if err != nil {
		if !strings.Contains(u.Hostname(), "http://") || !strings.Contains(u.Hostname(), "https://") {
			fmt.Println("URL is invalid")
			return
		}
		fmt.Println("URL is invalid")
		return
	}
	var wg sync.WaitGroup
	var requests_made uint64 = 0
	for requests_made < *MAXREQUEST {
		requests_made++
		sem <- true
		wg.Add(1)
		go func() {
			defer func() {
				<-sem
				wg.Done()
			}()
			if *POST {
				httpclient.PostRequest(*URL, *DATA, headers, *TIMEOUT, *PROXY)

			} else {
				httpclient.GetRequest(*URL, headers, *TIMEOUT, *PROXY)

			}
			progress := fmt.Sprintf("\rRequests sent: %v/%v", requests_made, *MAXREQUEST)
			fmt.Print(progress)
		}()
	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	wg.Wait()
}

func printBanner() {
	banner := `      ___                     ___           ___           ___           ___     
	/  /\      ___          /__/\         /  /\         /__/|         /  /\    
   /  /::\    /  /\         \  \:\       /  /::\       |  |:|        /  /::\   
  /  /:/\:\  /  /:/          \  \:\     /  /:/\:\      |  |:|       /  /:/\:\  
 /  /:/~/:/ /__/::\      _____\__\:\   /  /:/~/::\   __|  |:|      /  /:/~/::\ 
/__/:/ /:/  \__\/\:\__  /__/::::::::\ /__/:/ /:/\:\ /__/\_|:|____ /__/:/ /:/\:\
\  \:\/:/      \  \:\/\ \  \:\~~\~~\/ \  \:\/:/__\/ \  \:\/:::::/ \  \:\/:/__\/
 \  \::/        \__\::/  \  \:\  ~~~   \  \::/       \  \::/~~~~   \  \::/     
  \  \:\        /__/:/    \  \:\        \  \:\        \  \:\        \  \:\     
   \  \:\       \__\/      \  \:\        \  \:\        \  \:\        \  \:\    
	\__\/                   \__\/         \__\/         \__\/         \__\/    `
	fmt.Println(banner)
}
