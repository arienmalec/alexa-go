# Go Request/Response Helpers for Alexa Skill Services
### Install

```console
go get github.com/arienmalec/alexa-go
```

### Usage

#### Reponse

A minimal AWS Lambda implementing "Hello, World" as an Alexa skill in Go.

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