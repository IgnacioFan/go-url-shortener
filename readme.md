# Go URL Shortener
A URL shortener creates shorter aliases for long and complex URLs. Users can paste a link and receive a shortened code or paste the code to be redirected to the original page. The system is designed to support high availability and scalability to handle high volumes of requests.

## Technologies Used
- `Gin` as the web framework
- `Gorm` as the ORM framework
- `Cobra` as the command-line interface framework
- `PostgreSQL` for storing relational data
- `Redis` for building the cache layer
- `Mockery` and `sqlMock` for supporting unit test
- `Docker` for quickly setting up containerized services
- `Github actions` for CI

## Installation and Setup

### Prerequisite
- Go 1.16+
- Docker, we use `docker-compose` to boot up all required services

### Run it via Docker compose

1. Clone this repository to local
2. Navigate to the repo

```
// start all services
docker compose up

// stop all services
docker compose down
```

## Usage

### API: Shorten a URL
- Send a POST request to `http://localhost:3000/api/shorten` with a JSON body containing a long URL you want to shorten. For example:

```json
{
  "url": "www.example.com/foo/bar?user=123"
}
```

- The response will be a JSON object containing the original URL, the shortened URL, and a shortcode. For example:

```json
{
    "message": "Short URL created successfully",
    "data": {
        "short_url": "abc",
        "expiration": null
    }
}
```

### API: Redirect to the original URL
-  send a GET request to `http://localhost:3000/{shortcode}`, For example, `http://localhost:8080/abc` will redirect to `www.example.com/foo/bar?user=123`.

## Project Structure
```
├── build/                  # docker build file
├── cmd/
│   ├── migrate.go          # sub-command interface for implementing schema migration
│   ├── root.go             # main command interface
│   └── server.go           # sub-command interface for booting up the Http Server
├── deployment/
│   ├── config/             # configuration code and files
│   └── migration/          # migration code and files
├── internal/
│   ├── delivery/           # web layer
│   │   ├── handler/        # HTTP request handlers for the API endpoints and their tests
│   │   └── http_server.go  # HTTP server initializer and API routes settings
│   ├── entity/             # domanin layer
│   ├── mocks/              # contains generated mock files for different test purposes
│   ├── repository
│   │   ├── postgres/       # data access layer based on PostgreSQL
│   │   └── redis/          # cache layer based on Redis
│   └── usecase             # implemenation layer
├── scripts
└── main.go                 # main entry
```

## Test

We can run tests without `docker compose up`

```
// run thru all test cases
go test -v ./internal/...

// run specific test files
go test -v ./internal/usecase/.

// run thru all test cases with test coverage
go test -v -coverprofile=coverage.txt -cover ./internal/...
```

## Contributors
- Weilong Fan (IgnacioFan): developer and maintainer

## References
- https://github.com/JamesYu608/piccollage-problem2-shorten-url
- https://github.com/davidwu1997/ShortURL
- https://www.youtube.com/watch?v=JQDHz72OA3c
