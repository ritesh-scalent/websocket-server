postgresinit:
	docker run --name postgres15 -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

# go project 

goserver:

	go build -o gowebsocketserver .
	./gowebsocketserver

.PHONY:goserver postgresinit
