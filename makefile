redis:
	docker run -d --name counter -p 6379:6379 --rm redis

run:
	docker run -d --name counter -p 6379:6379 --rm redis
	go run ./cmd/api/main.go