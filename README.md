# AWS Lambda Go ASCII Image Converter

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)](https://golang.org/doc/go1.24)
[![AWS Lambda](https://img.shields.io/badge/AWS-Lambda-orange.svg)](https://aws.amazon.com/lambda/)
[![Serverless](https://img.shields.io/badge/Serverless-Framework-red.svg)](https://serverless.com/)

A serverless AWS Lambda function written in Go that converts uploaded images into ASCII art. The function accepts image uploads via HTTP POST requests and returns the converted ASCII representation.

## Features

- 🖼️ Convert images to ASCII art
- 🎨 Colored ASCII output support
- ☁️ Serverless deployment on AWS Lambda
- 📤 Multipart file upload handling
- ⚡ Fast Go-based processing
- 🔧 Configurable ASCII dimensions (50x25 default)

## Architecture

The application uses:
- **AWS Lambda** for serverless compute
- **API Gateway** for HTTP endpoints
- **CloudFormation** for infrastructure as code
- **Go 1.24.2** with custom runtime
- **[ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter)** library for image processing

## Prerequisites

- AWS CLI configured with appropriate permissions
- Go 1.24.2 or later
- Bash shell for deployment scripts

## Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/EvgeniiKlepilin/aws-lambda-go-ascii-image-converter.git
cd aws-lambda-go-ascii-image-converter
```

### 2. Run Tests

```bash
./0-run-tests.sh
```

### 3. Create S3 Bucket for Artifacts

```bash
./1-create-bucket.sh
```

### 4. Deploy to AWS

```bash
./2-deploy.sh
```

### 5. Test the Function

```bash
./3-invoke.sh
```

### 6. Cleanup Resources

```bash
./4-cleanup.sh
```

## Usage

### API Endpoint

Once deployed, the function accepts POST requests to the `/upload` endpoint:

```bash
curl -X POST \
  https://your-api-gateway-url/dev/upload \
  -H "Content-Type: multipart/form-data" \
  -F "file=@path/to/your/image.png"
```

### Response Format

The function returns the ASCII art as plain text in the response body:

```
                                                  
                                                  
                      ....           ...          
                   :=++++++=:    :=++++++=:       
         ......  :++++===++++= :++++===+++++.     
      ..::::::: :+++=.    ::. -+++=.    :++++     
     .....::::  +++=   =======+++=       =+++.    
          .... :+++=  -===+++++++=       ++++     
                ++++-   .-+++=++++-.   :=+++.     
                .+++++++++++- .+++++++++++=.      
                  :=+++++=:     :=+++++=-.        
                                                  
```

## Configuration

The ASCII conversion settings can be modified in `function/main.go`:

```go
flags := aic_package.DefaultFlags()
flags.Dimensions = []int{50, 25}  // Width x Height
flags.Colored = true              // Enable color output
flags.SaveTxtPath = "."          // Save path for text files - useful for testing
flags.SaveImagePath = "."        // Save path for images - useful for testing
```

## Project Structure

```
.
├── function/
│   ├── main.go              # Lambda function code
│   ├── main_test.go         # Unit tests
│   ├── go.mod               # Go module definition
│   ├── go.sum               # Go module checksums
│   └── test_image.png       # Test image file
├── 0-run-tests.sh          # Run unit tests
├── 1-create-bucket.sh      # Create S3 deployment bucket
├── 2-deploy.sh             # Deploy to AWS
├── 3-invoke.sh             # Test deployed function
├── 4-cleanup.sh            # Clean up AWS resources
├── template.yml            # CloudFormation template
├── event.json              # Test event payload
└── README.md               # This file
```

## Dependencies

- [aws-lambda-go](https://github.com/aws/aws-lambda-go) - AWS Lambda Go runtime
- [ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter) - ASCII art conversion
- [go-awslambda](https://github.com/grokify/go-awslambda) - Lambda utilities

## Deployment Details

The deployment process:

1. **Creates S3 bucket** for storing deployment artifacts
2. **Builds Go binary** with Linux target and Lambda tags
3. **Packages CloudFormation template** with S3 artifacts
4. **Deploys stack** using CloudFormation with IAM capabilities
5. **Creates API Gateway** endpoint automatically

### CloudFormation Resources

- **AWS::Serverless::Function** - Lambda function
- **AWS::IAM::Role** - Execution role with basic Lambda permissions
- **AWS::ApiGateway::RestApi** - API Gateway (implicit)

## Error Handling

The function handles common error scenarios:

- Invalid multipart requests
- Unsupported image formats  
- Image processing failures
- Missing file uploads

## Performance

- **Runtime**: provided.al2023 (custom Go runtime)
- **Timeout**: 5 seconds
- **Memory**: Default Lambda allocation
- **Cold start**: ~100-500ms for Go Lambda

## Monitoring

The function includes:

- **AWS X-Ray tracing** for request tracking
- **CloudWatch logs** for debugging
- **Basic execution metrics** via CloudWatch

## Development

### Local Testing

```bash
cd function
go test -v
```

### Building Locally

```bash
cd function
GOOS=linux go build -tags lambda.norpc -o bootstrap main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [TheZoraiz/ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter) for the excellent ASCII conversion library
- AWS Lambda team for the Go runtime support
- Go community for the robust ecosystem
