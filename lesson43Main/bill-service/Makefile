DB_URL='postgres://postgres:1111@localhost:5433/gollang?sslmode=disable'
run:
	go run cmd/main.go
mig-up:
	migrate -database ${DB_URL} -path migrations up
mig-down:
	migrate -database ${DB_URL} -path migrations down
mig-force:
	migrate -database ${DB_URL} -path migrations force 1
mig-goto:
	migrate -database ${DB_URL} -path migrations goto 1
mig-file:
	migrate create -ext sql -dir migrations -seq user

mod:
	go mod init api_getaway_service
	go mod tidy
	go mod vendor
tidy:
	go mod tidy
	go mod vendor



