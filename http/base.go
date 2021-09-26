package http

import (
	"crypto/tls"
	"fmt"
	"gitlab-cli-interface/http/structs"
	"io"
	"net/http"
	"os"
)

var (
	onlyReadAccessToken = "zysUgfeijUX_-zNWGxeZ"
)

func AllIssues() {
	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transportConfig}

	request, err := http.NewRequest("GET", "https://localhost:9010/api/v4/issues", nil)
	request.Header.Add("PRIVATE-TOKEN", onlyReadAccessToken)

	if err != nil {
		fmt.Println("The http client was unable to be constructed")
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Response body reader could not be closed. Please check for mem leaks.")
		}
	}(response.Body)

	rawBytes, err := io.ReadAll(response.Body)
	issues, err := structs.UnmarshalIssuesJSON(rawBytes)
	for _, issue := range issues {
		outStr := issue.GenericCSVRow()
		fmt.Println(outStr)
	}
	outFile, _ := os.Create("output.csv")
	defer outFile.Close()
	issues.MarshalCSV(outFile)
}
