package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/grokify/go-awslambda"
)

var (
	ErrImageConversion = errors.New("Cannot convert provided image")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{}
	r, err := awslambda.NewReaderMultipart(request)
	if err != nil {
		return res, err
	}
	part, err := r.NextPart()
	if err != nil {
		return res, err
	}
	_, err = io.ReadAll(part)
	if err != nil {
		fmt.Println("io.ReadAll error")
		return res, err
	}
	filePath := part.FileName()

	flags := aic_package.DefaultFlags()

	flags.Dimensions = []int{50, 25}
	flags.Colored = true
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		return res, ErrImageConversion
	}

	res.Body = fmt.Sprintf("%v", asciiArt)
	res.StatusCode = 200

	return res, nil
}

func main() {
	lambda.Start(handler)
}
