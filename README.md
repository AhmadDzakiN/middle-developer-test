# Employee API 
An Employee API by Ahmad Dzaki Naufal for mid-developer-test

## For more detailed information, please see this

## Tech Stack
- Go
- PostgreSQL

## How to Start
- Clone this repo
- Copy file `params/.env.sample` to `params/.env`
- Modify `params/.env` as your needs
- Ensure your PostgreSQL database already up
- Run the project by `go run main.go` or `go build && ./middle-developer-test`

## List of Comamnds
### API:
- `go run main.go` or `go build && ./middle-developer-test`: To run the API server
- `go run main.go migrate`: To up all available database migrations
- `go run main.go migratenew {your_migration_file_name}`: To create a new database migration file (the file will be created at migrations/sql folder)
- `go run main.go migratedown`: To migrate down only one database migration

### Docker:
- `docker build -t {your-postgresdb-image-name} -f postgresdb.dockerfile .`: To build PostgreSQL docker image
- `docker container run -d --name {your-postgresdb-container-name} -p 5432:5432 {your-postgresdb-image-name}`: To run/up PostgreSQL docker container based on previous point config with 5432 exposed port
- `docker-compose up -d --build`: To set up and run PostgreSQL & API docker image & container