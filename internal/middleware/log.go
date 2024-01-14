package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
)

func LogRequestAndResponseBody(ctx echo.Context, requestBodyBytes, responseBodyBytes []byte) {
	requestBodyJson := bytesToJson(requestBodyBytes)
	if requestBodyJson != nil {
		fmt.Printf("request body: %s\n", string(requestBodyJson))
	}

	responseBodyJson := bytesToJson(responseBodyBytes)
	if responseBodyJson != nil {
		fmt.Printf("reponse body: %s\n", string(responseBodyJson))
	}
}

func bytesToJson(bytes []byte) []byte {
	var bytesMap map[string]string
	err := json.Unmarshal(bytes, &bytesMap)
	if err != nil {
		return nil
	}

	bytesJson, err := json.Marshal(bytesMap)
	if err != nil {
		return nil
	}

	return bytesJson
}
