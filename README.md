# genvaeg


# Description

https://github.com/shortcut/sweden-coding-challenge-backend/

# Layout

I don't like complications for the reason that there may be a need in the future. The code should be simple and easy to understand today to make changes tomorrow easier. Even though this is small application, it has enough moving parts that warrant a minimum set of layers; API, routes, handlers, repository, plus some utility directories.

- I've been following a common structure of the code. The API and CLI start in the cmd folder.
- The application uses the Echo routing framework, but any routing framework would do the job just as well.
- Since this is a small application, the handlers communicate directly with the database.
- For simplicity sake I chose to use SQLite3 since it's only a file, so there is no database server to start. This can easily be replaced by a db server like MySQL or PostgreSQL.
- The shortener creates a random short code, and will retry until this code is unique.
- I like to to "make", since it makes my life easier.

# Development

- See Makefile for commands.
- You can run the server with a watch (`$> make serverwatch`). It uses Air for live reload (https://github.com/cosmtrek/air).

# Automatic tests

- The tests will run against it's own database, which is called test.db.
- There is only a test for the creation of the short URL code.

# Manually testing the app using cURL

To reset the database and get a prettier list of users and URLs, see CLI below

### Pre

1. Build and start server: `make server`
2. Build CLI: `make build_cli`

### Run

1. Signup as user `abe` with password `abe`
   - `curl -X POST -L 'http://localhost:8080/signup' -d 'name=abe&pw=abe&repeatpw=abe' -H 'Content-Type: application/x-www-form-urlencoded'`
2. Login as `abe`
   - `curl -X POST -L -c cookies.txt 'http://localhost:8080/login' -d 'name=abe&pw=abe' -H 'Content-Type: application/x-www-form-urlencoded'`
3. Create short URL
   - `curl -L -b cookies.txt 'http://localhost:8080/create' -d 'url=http://vg.no' -H 'Content-Type: application/x-www-form-urlencoded'`
4. List short codes to use for to test redirect.
   - `./bin/cli urls`
5. Test redirect
   - `curl -L 'http://localhost:8080/r/<a short code from p4>'`
6. Logout
   - `curl -L 'http://localhost:8080/logout'`

# CLI (kind of a backoffice)

With the CLI you can:
- reset the database
- create and list users
- list all URLs

### Run

- Build : `$> make build_cli`
- Help  :  `./bin/cli`

# Further work towards production

- Use a database server, like MySQL/MariaDb, PostgreSQL, preferable with Docker.
- Remove .env file from Git repo, since this will be different in production, and will probably contain secrets codes.
- Use TLS. The Echo is using Let's Encrypt.
- Dockerize the server for easier handling behind a firewall.
- More tests, e.g. integration tests from http to database.
- Add extra layers to the design - like a service layer - if needed.
- Fill out the functionality where needed, e.g. with more CRUD operations.
- Add authorization if needed.
- Maybe move code behind the internal folder, so it won't be seen by other projects.

# Known bugs

- Logout doesn't seem to invalidate the cookie.




