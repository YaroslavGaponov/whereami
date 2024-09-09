cli:
	DATAFILE=geodata/worldcities.zip@worldcities.csv go run cmd/cli/main.go
server:
	DATAFILE=geodata/worldcities.zip@worldcities.csv PORT=7777 go run cmd/server/main.go
build:
	go build  -o whereamid cmd/server/main.go
docker:
	docker build -t whereamid:latest .