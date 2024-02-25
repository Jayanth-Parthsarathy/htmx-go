run: build
	@tmp/main
build:generate
	@go build -o tmp/main cmd/htmx-go/main.go
generate:
	@templ generate
