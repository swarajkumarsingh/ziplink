package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func sendGetRequest(targetURL string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Check if the response status code is 308 (Permanent Redirect).
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 redirects")
			}

			if req.Response != nil && req.Response.StatusCode == http.StatusPermanentRedirect {
				// Handle status code 308 (Permanent Redirect) here.
				return fmt.Errorf("stopped at 308 Permanent Redirect")
			}

			return nil
		},
	}

	startTime := time.Now()

	resp, err := client.Get(targetURL)
	if err != nil {
		if !strings.Contains(err.Error(), "308") {
			fmt.Printf("Error sending GET request: %v\n", err.Error())
		}
	}
	defer resp.Body.Close()

	// Calculate and log the response time.
	elapsed := time.Since(startTime)
	fmt.Printf("Request sent with status code: %d, Final URL: %s, Response Time: %s\n", resp.StatusCode, resp.Request.URL, elapsed)
}

func GET() {
	// Define the target URL of your Go backend.
	targetURL := "http://localhost:8080/AAAAAAB"

	// Define the number of concurrent requests.
	concurrentRequests := 1

	// Define the total number of requests to send.
	totalRequests := 10000

	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go sendGetRequest(targetURL, &wg)

		if i%concurrentRequests == 0 {
			// Limit the number of concurrent requests.
			wg.Wait()
		}
	}

	wg.Wait()

	elapsed := time.Since(startTime)
	fmt.Printf("Load test completed. Total time: %s\n", elapsed)
}

func POST() {
	// Define the target URL of your Go backend for POST requests.
	targetURL := "http://localhost:8080/create-url"

	// Define the number of concurrent POST requests.
	concurrentRequests := 100

	// Define the total number of POST requests to send.
	totalRequests := 10000

	var wg sync.WaitGroup
	responseTimes := make(chan time.Duration, totalRequests)

	startTime := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go sendPostRequest(targetURL, &wg, responseTimes)

		if i%concurrentRequests == 0 {
			// Limit the number of concurrent requests.
			wg.Wait()
		}
	}

	wg.Wait()
	close(responseTimes)

	// Calculate the average response time for POST requests.
	var totalResponseTime time.Duration
	count := 0

	for responseTime := range responseTimes {
		totalResponseTime += responseTime
		count++
	}

	if count > 0 {
		averageResponseTime := totalResponseTime / time.Duration(count)
		fmt.Printf("Average Response Time for %d POST requests: %s\n", count, averageResponseTime)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Load test completed. Total time: %s\n", elapsed)
}

func sendPostRequest(targetURL string, wg *sync.WaitGroup, responseTimes chan time.Duration) {
	defer wg.Done()

	startTime := time.Now()

	client := &http.Client{}

	// Define the payload for your POST request.
	payload := []byte(`{"url": "https://instagram.com/"}`) // Replace with your POST data

	resp, err := client.Post(targetURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Error sending POST request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(startTime)
	responseTimes <- elapsed
	fmt.Printf("Request sent with status code: %d, Response Time: %s\n", resp.StatusCode, elapsed)
}

func main() {
	GET()
	// POST()
}

