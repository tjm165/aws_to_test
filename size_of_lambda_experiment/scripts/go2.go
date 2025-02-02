package main

import (
        "fmt"
        "github.com/aws/aws-lambda-go/lambda"
        "bytes"
)

type Request struct {
        ID float64 `json:"id"`
        Value string `json:"value"`
}

type Response struct {
        Body string `json:"body"`
        StatusCode int `json:"statusCode"`
}

type MyDict struct {
	m string
	tommys []string
}

func HandleRequest(request Request) (Response, error) {
	my_dict := MyDict{
		m:      "Hi this is a dict inside Go2!",
		tommys: []string{

                },
	}

	var message bytes.Buffer
	message.WriteString("{\"message\": \"")
	message.WriteString("Hi this is Go2! ")
	message.WriteString(my_dict.m)
	message.WriteString("\"}")

        return Response {
                Body: fmt.Sprintf(message.String()),
                StatusCode: 200,
        }, nil
}

func main() {
        lambda.Start(HandleRequest)
}