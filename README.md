# genvaeg


# Description

https://github.com/shortcut/sweden-coding-challenge-backend/

### Examples of other URL shorteners

- https://free-url-shortener.rb.gy/
- https://www.shorturl.at/shortener.php

# Layout

- The application uses the Echo routing framework.
- Since this is a small application, the handlers communicate directly with the database.
- The database is SQLite3, which means it's only a file, so there is no database server to start.
- The shortener creates a short code, and will retry if a code already exists.

# Development

- See Makefile for commands.
- You can run the server with a watch (`$> make serverwatch`). It uses Air for live reload (https://github.com/cosmtrek/air).

# Automatic tests

- The tests will run against it's own database, which is called test.db.
- There is only a test for the creation of the short URL code.

# Manually testing the app using cURL

To reset the database and get an easier read of users and URLs, see CLI below

### Pre

1. Build and start server: `make server`
2. Build CLI: `make build_cli`

### Run

1. Signup
   - `curl -X POST -L -v 'http://localhost:8080/signup' -d 'name=abe&pw=abe&repeatpw=abe' -H 'Content-Type: application/x-www-form-urlencoded'`
2. Login
   - `curl -X POST -L -v -c cookies.txt 'http://localhost:8080/login' -d 'name=abe&pw=abe' -H 'Content-Type: application/x-www-form-urlencoded'`
3. Create short URL
   - `curl -L -v -b cookies.txt 'http://localhost:8080/create?url=http://vg.no'`
4. List short codes
   - `./bin/cli urls`
5. Use short url
   - `curl -L -v 'http://localhost:8080/r/<a short code from p4>'`
6. Logout
   - `curl -L -v 'http://localhost:8080/logout'`

# CLI (kind of a backoffice)

With the CLI you can:
- reset the database
- create and list users
- list all URLs

### Run

- Build : `$> make build_cli`
- Help  :  `./bin/cli`

# Further work towards production

- Remove .env file from Git repo, since this will be different in production, and will probably contain secrets codes.
- Use TLS. The Echo is using Let's Encrypt.
- Use a database server, like MySQL/MariaDb, PostgreSQL, preferable with Docker.
- Dockerize the server for easier handling behind a firewall.
- Add extra layers to the design - like a service layer - if needed.
- Fill out the functionality where needed, e.g. with more CRUD operations.
- Add authorization if needed.
- More tests, e.g. integration tests from http to database.

