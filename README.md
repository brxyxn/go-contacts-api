<h1 align="center">Go PhoneBook API</h1>

<p align="center">
  <a href="https://github.com/rs/zerolog"></a>
  <img src="https://img.shields.io/badge/Golang_v1.22-4C566A.svg?logo=go&logoColor=ffffff&labelColor=00ADD8" alt="Golang">
  <img src="https://img.shields.io/badge/Swagger_v2-4C566A.svg?logo=swagger&logoColor=333333&labelColor=85EA2D" alt="Swagger">
  <img src="https://img.shields.io/badge/Chat_GPT_4-4C566A.svg?logo=openai&logoColor=ffffff&labelColor=412991" alt="Swagger">
  <img src="https://img.shields.io/badge/GitHub_Copilot-4C566A.svg?logo=githubcopilot&logoColor=ffffff&labelColor=000000" alt="Swagger">
    <br />
    <a href="https://github.com/gorilla/mux">
      <img src="https://img.shields.io/badge/Gorilla_Mux-4C566A.svg" alt="Gorilla Mux">
    </a>
    <a href="https://github.com/rs/zerolog">
      <img src="https://img.shields.io/badge/Zerolog-4C566A.svg" alt="Zerolog">
    </a>
    <a href="https://github.com/stretchr/testify">
      <img src="https://img.shields.io/badge/Testify-4C566A.svg" alt="Testify">
    </a>
</p>

This project was developed as required by AccelOne as a technical test developing and dockerizing a lite version of an
Rest API.

---

<details>
<summary>Table of contents</summary>

<!-- TOC -->
  * [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
    * [Running the API](#running-the-api)
    * [Building the API](#building-the-api)
    * [Testing the API](#testing-the-api)
    * [Checking the Server](#checking-the-server)
  * [Documentation](#documentation)
    * [Get all contacts](#get-all-contacts)
    * [Get a contact](#get-a-contact)
    * [Create a contact](#create-a-contact)
    * [Update a contact](#update-a-contact)
    * [Delete a contact](#delete-a-contact)
  * [Tests](#tests)
  * [API Design](#api-design)
    * [Decisions](#decisions)
    * [Improvements](#improvements)
<!-- TOC -->

</details>

<!-- badges -->

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [Postman](https://www.postman.com/downloads/) or
- [Curl](https://curl.se/download.html)

### Installation

```bash
git clone https://github.com/brxyxn/go-contacts-api && cd ./go-contacts-api
```

### Running the API

```bash
go run ./cmd/server/main.go
```

### Building the API

```bash
go build -o server ./cmd/server/main.go
```

### Testing the API

```bash
go test -v ./...
```

### Checking the Server

```bash
curl -X GET http://localhost:5000/status
# {"status":"active"}
```

## Documentation

> Note: For documentation purposes I prefer to use swagger in the first place.

There is a Postman collection that you can use to test the API. You can find it in the `docs` folder.

<details>
<summary>Click to expand</summary>

### Get all contacts

```http
GET /contacts
```

### Get a contact

```http
GET /contacts/${id}
```

### Create a contact

```http
POST /contacts
```

### Update a contact

```http
PUT /contacts/${id}
```

### Delete a contact

```http
DELETE /contacts/${id}
```

</details>

---

## Tests

> Note: When running tests uncomment the following line in the `contacts.go` file in `/api/v1/contacts`.

```go
package contacts

// ...

const contactsFile = "./api/v1/contacts/contacts.json"

// const contactsFile = "./contacts.json" // uncomment this line when running tests

// ...
```

```bash
$ go test -cover ./...
# avg. coverage: 50% of statements
        github.com/brxyxn/go-phonebook-api/api/status           coverage: 0.0% of statements
        github.com/brxyxn/go-phonebook-api/internal/app         coverage: 0.0% of statements
        github.com/brxyxn/go-phonebook-api/cmd/server           coverage: 0.0% of statements
ok      github.com/brxyxn/go-phonebook-api/api/v1/contacts      coverage: 67.3% of statements
ok      github.com/brxyxn/go-phonebook-api/config               coverage: 100.0% of statements
ok      github.com/brxyxn/go-phonebook-api/internal/response    coverage: 82.6% of statements
ok      github.com/brxyxn/go-phonebook-api/pkg/logger           coverage: 100.0% of statements
```

---

## API Design

### Decisions

1. **net/http**: I decided to use the `net/http` package to create the API because it is a standard package in Go that allows me to create a server and handle requests and responses, it is also more than capable of holding the requirements of the challenge. Even thought it allows me to crerate the server without any external dependencies, I decided to use Gorilla Mux as the router for the API.
2. **Gorilla Mux**: I decided to use Gorilla Mux as the router for the API because it is a powerful and flexible router that allows me to define routes with variables and query parameters. It also has a good performance and is widely used in the Go community.
3. **Response Models**: I decided to define a standard response format for the API that includes a status code, a message, an action, and a data field. This allows me to send consistent responses to the client and handle errors in a more structured way.
4. **Response Handling**: I decided to create a helper function to handle the response to the client. This eventually could become a standard way to handle responses in the API across all endpoints and other services (microservices if applicable).
5. **Logger**: I decided to use the `github.com/rs/zerolog` package to log information about the API to the console. This allows me to log information about the requests and responses and debug the API if needed. It is very flexible and allows me to log information in different formats and very structured.
6. **Environment Variables**: I decided to use environment variables to configure the API and implemented with the library `github.com/joeshaw/envdecode` because it allows us to tag the struct using default values, which requires zero configuarion when it's not provided. However, I'd use either JSON, YAML or TOML files to store the configuration in a production environment if required.

### Improvements

1. **Validation**: I would add validation to the request body and query parameters to ensure that the data is correct before processing it. This would help prevent errors and improve the reliability of the API.
2. **Error Handling**: I would improve the error handling in the API to provide more detailed error messages to the client. This would help the client understand what went wrong and how to fix it.
3. **Security**: I would add security features to the API to protect it from common attacks such as SQL injection, XSS, CSRF, etc. This would help prevent security breaches and protect the data of the users.
4. **Testing**: I would add enough unit tests and integration tests to the API to ensure that it works as expected and that it can handle different scenarios. This would help prevent regressions and improve the quality of the API.
5. **Observability**: I would add observability features to the API to monitor its performance and health. This would help detect issues early and improve the reliability of the API.

