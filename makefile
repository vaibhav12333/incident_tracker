.PHONY: server client start

server:
	cd server && go run main.go

client:
	cd client/tracker && npm start

start:
	$(MAKE) server &
	$(MAKE) client