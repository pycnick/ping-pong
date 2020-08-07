all: server client

server:
	go build -o server cmd/server/main.go

client:
	go build -o client cmd/client/main.go

run_server:
	./server

run_client:
	./client

clean:
	rm client && rm server
