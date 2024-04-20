.PHONY: build run-server


build:
	docker build -t gin-server .

run-server:
	docker run --rm -p 8080:8080 --name gogolook-interview-api-server gin-server

stop-server:
	docker stop gogolook-interview-api-server

