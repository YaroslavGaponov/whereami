cli:
	DATAFILE=geodata/worldcities.zip@worldcities.csv go run cmd/cli/main.go
	
server:
	DATAFILE=geodata/worldcities.zip@worldcities.csv PORT=8080 go run cmd/server/main.go

build:
	go build -o whereamid cmd/server/main.go

run:
	go build -o whereamid cmd/server/main.go
	DATAFILE=geodata/worldcities.zip@worldcities.csv PORT=8080 ./whereamid

docker:
	docker build -t whereamid:latest .
	docker run -p 8080:8080 whereamid:latest  