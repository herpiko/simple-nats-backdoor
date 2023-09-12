run:
	GOWORK=off go run main.go
nats:
	docker run -d --name nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222

req:
	nats request update "hehehe"
