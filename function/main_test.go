package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {

		bodyBuffer := new(bytes.Buffer)
		mw := multipart.NewWriter(bodyBuffer)

		part, err := mw.CreateFormFile("file", "test_image.png")
		if err != nil {
			t.Fatal(err)
		}
		tempFile, err := os.Open("test_image.png")
		if err != nil {
			t.Fatal(err)
		}
		tempFile.Close()
		io.Copy(part, tempFile)

		headers := map[string]string{
			"Content-Type": mw.FormDataContentType(),
		}
		mw.Close()

		req := events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: bodyBuffer.String(), Headers: headers}

		resp, err := handler(req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp.Body)
	})
}
