// This is a package comment.
package main

import (
	"fmt"
	"net/http"
)

func testWithHTTPHead(url string) {
	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Close the response body
	contentlength := res.ContentLength
	// Print the response status code
	fmt.Println("Response Status:", res.Status)
	fmt.Printf("ContentLength:%v\n", contentlength)
}

func testWithHTTPNewRequest(url string) {
	// Create a new HEAD request
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// Send the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}

func main() {
	url := "https://app.datadoghq.com/support/flare/0?api_key=abcdef"
	testWithHTTPHead(url)
	fmt.Println("======================================================")
	testWithHTTPNewRequest(url)
}
