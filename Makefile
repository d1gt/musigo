build:
	@go build -o bin/musigo cmd/musigo/musigo.go

build_with_race:
	@go build -o bin/musigo cmd/musigo/musigo.go

build_and_run: build
	@./bin/musigo

build_and_run_with_race: build_with_race
	@./bin/musigo

run:
	@go run cmd/musigo/musigo.go

test: 
	@go test ./... -v
