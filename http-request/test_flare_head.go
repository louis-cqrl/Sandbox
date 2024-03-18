// This is a package comment.
package main

import (
	"fmt"
	"net/http"
)

func acceptRedirection(url string, client *http.Client) (bool, error) {
	res, err := client.Head(url) // 307 with this function
	if err != nil {
		return false, err
	}
	defer res.Body.Close() // Close the response body
	if res.StatusCode == http.StatusTemporaryRedirect || res.StatusCode == http.StatusPermanentRedirect || res.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}

func testWithHTTPHead(url string, client *http.Client) {

	res, err := client.Head(url) // 307 with this function
	if err != nil {
		fmt.Println(err)
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
	urlString := resp.Request.URL.String() // Assign the result to a variable
	fmt.Println("URL:", urlString)         // Use the variable
	defer resp.Body.Close()
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}

func main() {
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	url := "https://7-53-0-flare.agent.datadoghq.com/support/flare"
	testWithHTTPHead(url, client)
	fmt.Println("======================================================")
	testWithHTTPNewRequest(url)
	fmt.Println("======================================================")
	isRedirect, err := acceptRedirection(url, client)
	if err != nil || !isRedirect {
		fmt.Println("Error:", err)
	} else if isRedirect {
		fmt.Println("Redirected successfully")
	}
	fmt.Println("==============try with no version url=================")
	ulr2 := "https://app.datadoghq.com//support/flare"
	testWithHTTPHead(ulr2, client)
	fmt.Println("======================================================")
	testWithHTTPNewRequest(ulr2)
	fmt.Println("======================================================")
	isRedirect, err = acceptRedirection(ulr2, client)
	if err != nil || !isRedirect {
		fmt.Println("Error:", err)
	} else if isRedirect {
		fmt.Println("Redirected successfully")
	}
}
