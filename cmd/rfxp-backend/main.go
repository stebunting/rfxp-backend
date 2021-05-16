package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stebunting/rfxp-backend/router"
)

func main() {
	lambda.Start(router.HandleLambdaEvent)
}
