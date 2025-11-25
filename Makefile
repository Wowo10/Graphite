# Build the application
build:	
	@GOOS=linux GOARCH=amd64 go build -o graphite cmd/graphite/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f graphite
