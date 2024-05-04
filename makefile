targets: run migrate

run:
	go run cmd/url-shortener/main.go --config=./config/local.yaml
migrate:
	go run ./cmd/migrator --storage-path=./storage/storage.db --migrations-path=./migrations