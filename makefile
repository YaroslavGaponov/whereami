build_cli:
		go build -o whereamicli cmd/cli/main.go
build:
		go build -o whereamid cmd/server/main.go

run: build
		LOG_LEVEL="all" DATA_FILE="data/worldcities.zip@worldcities.csv" SERVER_ADDRESS=":8080" ./whereamid


build_docker:
		docker build -t YaroslavGaponov/whereamid:latest .

run_docker: build_docker
		docker run -p 8080:8080 YaroslavGaponov/whereamid:latest