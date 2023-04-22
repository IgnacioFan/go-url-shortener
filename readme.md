# Go URL Shortener
A URL shortener creates shorter aliases for long and complex URLs. Users can paste a link and receive a shortened code or paste the code to be redirected to the original page. The system is designed to support high availability and scalability to handle high volumes of requests.

## Technologies Used
- `Gin` as the web framework
- `Gorm` as the ORM framework
- `Cobra` as the command-line interface framework
- `PostgreSQL` for storing relational data
- `Redis` for building the cache layer
- `Mockery` and `sqlMock` for unit test
- `Docker` for building containerized services and env
- `Wire` for supporting components using dependency injection
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
- https://github.com/vishnubob/wait-for-it (for script)
