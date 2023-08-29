GO = go
BUILD = build
RUN = run
SERVER = server.go
CLIENT = client.go

all:
		go build $(SERVER)
		go build $(CLIENT)
fclean:
		rm -rf server
		rm -rf client

re: fclean all

.PHONY: fclean all
