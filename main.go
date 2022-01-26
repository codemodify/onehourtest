package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	loggingFormatter "goframework.io/logging-formatters-timerfc3339nano"
	loggingFilter "goframework.io/logging-mixers-logtolevel"
	loggingSpec "goframework.io/logging-spec"
	loggingStyle "goframework.io/logging-styles-classic"

	loggingStorage2 "goframework.io/logging-persisters-console"
	loggingStorage1 "goframework.io/logging-persisters-file"

	httpHelpers "goframework.io/networking-http"
)

func main() {
	appConfig := getConfig()

	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	// setup logger and config file
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	logConfig := loggingSpec.NewPipe([]loggingSpec.Logger{
		loggingFilter.NewLogger(loggingSpec.LevelDebug),
		loggingFormatter.NewLogger(),
		loggingStorage1.NewLogger(),
		loggingStorage2.NewLogger(),
	})

	logger := loggingStyle.NewLogger(logConfig)

	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	// GET http://petstore-demo-endpoint.execute-api.com/petstore/pets
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	headers := map[string]string{}
	timeout := 10 * time.Second
	httpResp, rawData, err := httpHelpers.DoHTTPRequest(
		http.MethodGet,
		headers,
		appConfig.PetsURL,
		nil,
		timeout,
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		logger.Error(fmt.Errorf("Invalid HTTP Response"))
		os.Exit(1)
	}

	logger.Debugf("RAW-DATA: %s", string(rawData))

	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	// POST http://petstore-demo-endpoint.execute-api.com/petstore/pets
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	newPet := Pet{
		// ID:    9999,
		Type:  "bird",
		Price: 2.1,
	}

	headers = map[string]string{
		"Content-type": "application/json",
	}
	httpResp, rawData, err = httpHelpers.DoHTTPRequest(
		http.MethodPost,
		headers,
		appConfig.PetsURL,
		[]byte(PetToJSONString(newPet)),
		timeout,
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		logger.Error(fmt.Errorf("Invalid HTTP Response"))
		os.Exit(1)
	}

	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	// GET http://petstore-demo-endpoint.execute-api.com/petstore/pets/9999
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	headers = map[string]string{}
	httpResp, rawData, err = httpHelpers.DoHTTPRequest(
		http.MethodGet,
		headers,
		fmt.Sprintf("%s/%d", appConfig.PetsURL, newPet.ID),
		nil,
		timeout,
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		logger.Error(fmt.Errorf("Invalid HTTP Response"))
		os.Exit(1)
	}

	logger.Debugf("RAW-DATA: %s", string(rawData))

	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
	// DELETE http://petstore-demo-endpoint.execute-api.com/petstore/pets/9999
	// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

	headers = map[string]string{}
	httpResp, rawData, err = httpHelpers.DoHTTPRequest(
		http.MethodDelete,
		headers,
		fmt.Sprintf("%s/%d", appConfig.PetsURL, newPet.ID),
		nil,
		timeout,
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		logger.Error(fmt.Errorf("Invalid HTTP Response"))
		os.Exit(1)
	}

	logger.Debugf("RAW-DATA: %s", string(rawData))
}
