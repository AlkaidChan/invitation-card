.PHONY: tidy docker docker-release build-release

tidy:
	go mod tidy

docker:
	docker build . -t alkaidchen/invitation-card:latest

docker-release: docker
	docker push alkaidchen/invitation-card:latest

build-release: tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/invitation-card-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/invitation-card-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/invitation-card-darwin-amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/invitation-card-darwin-arm64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/invitation-card-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/invitation-card-windows-arm64.exe main.go
