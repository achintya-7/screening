gen:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go screening/db/sqlc Store

server:
	go run main.go