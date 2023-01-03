install-dependencies:
	echo "Installing dependencies..."
	go get ./...

graphql-generate:
	echo "Generating graphql files..."
	go get github.com/99designs/gqlgen@v0.17.22
	go run github.com/99designs/gqlgen generate --verbose

run-playground:
	echo "Running graphql playground..."
	go run ./server.go