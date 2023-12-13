# go-test-api

## Prerequisites

Make sure you have Go installed. You can download it from [Go official website](https://golang.org/dl/).

## Installation

To build the project, run the following command in the terminal:

```bash
make
```

## Start App

To start the app, use the following command:

```bash
make run
```

This will launch the "go test API" on your local machine.

## Configuration

Create a `.env` file in the root directory using the provided `.env.example` as a template. Fill in the necessary attributes with your specific values.

## API Endpoints

### Get Info

**Request**

- URL: `{{baseurl}}/v1/info`
- Request Type: GET

**Response**

```json
{
  "status": "success",
  "payload": {
    "name": "go-test-api",
    "version": "1.0.0"
  }
}
```

### Get My Data

**Request**

- URL: `{{baseurl}}/v1/my`
- Request Type: GET
- Headers: `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjI1MTYyMzkwMjJ9.EYWLJUkEpZJauIt8QPYHc2EKRK3eqtSVq2XYRzjVEn0`

**Response**

```json
{
  "status": "success",
  "payload": {
    "exp": 2516239022,
    "iat": 1516239022,
    "name": "John Doe",
    "sub": "1234567890"
  }
}
```

Make sure to replace \`{{baseurl}}\` with the actual base URL of your deployed application.
