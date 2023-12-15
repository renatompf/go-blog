# Go-Blog

## Description
This is a CRUD Web Application made in [Golang](https://go.dev/) made to simulate a blog.
This application used [PostgreSQL](https://www.postgresql.org/) running on [Docker](https://www.docker.com/) in order to have a database.

To build this it was also used the [GIN Framework](https://gin-gonic.com/) as Web Framework and [GORM Framework](https://gorm.io) to handle the data. 

## How to run it:

1. Since the models will be automated created as tables in the database, to start the application you can simply make the following command, and then you can start to make requests to `localhost:8080`.

```shell
make run
```

## How to test it:

### Get All Posts:
 * GET request to `localhost:8080/posts`

### Get Posts by Id:
 * GET request to `localhost:8080/posts/:id`

### Add new Post:
 * POST request to `localhost:8080/posts`

```json
{
  "title": "How to Write a Web Application in Go",
  "body": "Using Gin and GORM makes writing a web application in Go really easy, as both frameworks are incredibly helpful to the user."
}
```

### Update Post By Id:
 * PUT request to `localhost:8080/posts/:id`
```json
{
  "title": "Effective Concurrency in Go",
  "body": "Leveraging Go's concurrency primitives like goroutines and channels is key to building robust and highly performant applications. Mastering these features allows developers to create scalable and efficient software systems."
}
```

### Delete Post By Id:
 * DELETE request to `localhost:8080/posts/:id`
