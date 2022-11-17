run:
	air -c .air.conf
    #godotenv -f .env go run cmd/admin/main.go --port=3100

build:
	go build -ldflags "-s -w" -o ./tmp/server ./auth/cmd/main.go