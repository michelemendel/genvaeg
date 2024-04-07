# --------------------------------------------------------------------------------
# Build

build_server:
	@go build -o bin/server ./cmd/api/...

build_pg:
	@go build -o bin/pg main.go

# --------------------------------------------------------------------------------
# Run

server: build_server
	@echo "Starting web server"
	@./bin/server

serverwatch: build_server
	@echo "Starting web server with air to watch for changes"
	@air

pg: build_pg
	@./bin/pg