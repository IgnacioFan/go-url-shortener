# Go URL Shortener
The Go URL shortener project provides users with the ability to generate unique shortened codes and redirect them to the original URLs. It was developed using technologies such as Gin, Cobra, Gorm, PostgreSQL, Redis, Docker, and Nginx, with a focus on ensuring high availability.

## Installation and Setup

### Prerequisite
- Go 1.16+
- Docker, the project is based on `docker-compose.yml` to boot up all runnable services

### Run it via Docker compose

1. Clone the repository to your local machine.
2. Ensure that Docker is installed and running on your machine.
3. Create a `.env` file and do some settings.
4. Run `make app.start`
5. Test the following API endpoints
    - Create a shortened URL. For example, send a POST request to `http://localhost/api/v1/urls` with a JSON body
    ```json
    {
      "url": "www.example.com/foo/bar?user=123"
    }
    ```
    - URL Redirecting. For example, send a GET request to `http://localhost/${shortcode}`
    - Delete a shortened URL. For example, send a DELETE request to `http://localhost/api/v1/urls/{shortcode}`

6. Run `make app.stop` to clean up the containers

End! To know more about other executable commands, please check out the Makefile.

## System design
TBD

## Contributors
- Weilong Fan (IgnacioFan): developer and maintainer

## References
- https://github.com/JamesYu608/piccollage-problem2-shorten-url
- https://github.com/davidwu1997/ShortURL
- https://www.youtube.com/watch?v=JQDHz72OA3c
- https://github.com/vishnubob/wait-for-it (for script)
