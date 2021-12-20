build:
	docker-compose build petprjct

run:
	docker-compose up petprjct

test:
	go test -v ./...
