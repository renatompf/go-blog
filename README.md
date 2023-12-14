# Go-Blog

This is a CRUD Web Application made in [Golang](https://go.dev/) made to simulate a blog.
This application used [PostgreSQL](https://www.postgresql.org/) running on [Docker](https://www.docker.com/) in order to have a database.

To build this it was also used the [GIN Framework](https://gin-gonic.com/) as Web Framework and [GORM Framework](https://gorm.io) to handle the data. 

# How to run it:

1. You have to start the database by making `docker compose up`
2. Then, inside of [go-blog/migrate](migrate) directory run: `go run .`. This will run the migration and create the `Post` table in the database;
3. To finish and start the application in the main directory you make `go run .`

# How to test it:

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
