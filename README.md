# Go Request/Response and Helpers for Alexa Skill Services

### Install

```console
go get github.com/arienmalec/alexa-go
```

### Usage

#### Response

The `alexa.Response` struct implements the AWS Alexa Skill response, and contains a helper for simple speach responses.

The following is a minimal AWS Lambda implementing "Hello, World" as an Alexa skill in Go.

```go
package main

import (
	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the lambda hander
func Handler() (alexa.Response, error) {
	return alexa.NewSimpleResponse("Saying Hello", "Hello, World"), nil
}

func main() {
	lambda.Start(Handler)
}
```

### Request

The `alexa.Request` struct implements the AWS Alexa Skill request, and contains some constants for locales and intents.

The following is a Lambda delivering localized content to users and handling multiple intents.

```go
package main

import (
	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "hello":
		response = handleHello(request)
	case alexa.HelpIntent:
		response = handleHelp()
	}

	return response
}

func handleHello(request alexa.Request) alexa.Response {
	title := "Saying Hello"
	var text string
	switch request.Body.Locale {
	case alexa.LocaleAustralianEnglish:
		text = "G'day mate!"
	case alexa.LocaleGerman:
		text = "Hallo Welt"
	case alexa.LocaleJapanese:
		text = "こんにちは世界"
	default:
		text = "Hello, World"
	}
	return alexa.NewSimpleResponse(title, text)
}

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse("Help for Hello", "To receive a greeting, ask hello to say hello")
}

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
}

func main() {
	lambda.Start(Handler)
}
```

### Credits

Request/Response struct layout influenced by `https://github.com/mikeflynn/go-alexa` which was written before Go was an AWS Lambda native language.

