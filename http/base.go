package http

import (
	"crypto/tls"
	"fmt"
	config2 "gitlab-cli-interface/config"
	"gitlab-cli-interface/http/structs"
	"io"
	"net/http"
	"os"
)

var (
	onlyReadAccessToken string
)

func AllIssues() {
	// Get the global config object.
	configuration := config2.GetConfig()
	onlyReadAccessToken = configuration.GitlabConnection.AccessTokens.ReadOnly

	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: configuration.GitlabConnection.AllowInsecure,
		},
	}
	client := &http.Client{Transport: transportConfig}

	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%v/api/v4/issues", configuration.GitlabConnection.Base),
		nil,
	)
	request.Header.Add("PRIVATE-TOKEN", onlyReadAccessToken)

	if err != nil {
		fmt.Println("The http client was unable to be constructed")
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("---ERROR---\n%v\n-----------", err.Error())
		return
	}

	if response.StatusCode != 200 {
		fmt.Printf("An invalid response was received from GitLab\n%v", response.Status)
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
