
brew install golang-migrate

Migrate create Schema 

migrate create -ext sql -dir db/migration -seq init_schema
/Users/jsiddiqui/db/migration/000001_init_schema.up.sql
/Users/jsiddiqui/db/migration/000001_init_schema.down.sql


migrate -help
migrate --path db/migration --database "postgresql://root:secret@localhost:5432/database_name" --verbose up 

If you receive SSL issue
migrate --path db/migration --database "postgresql://root:secret@localhost:5432/database_name?sslmode=disable" --verbose up 



sqlc init - It will create sqlc.yaml
sqlc generate




Some Docker Commands

docker pull image:tag
docker run --name container_name -e environment_variable -d image:tag
docker run --name container_name -e environment_variable -p hostport:containerport -d image:tag



postgres://postgres:postgrespw@localhost:32769

docker run -p 5050:80 -e 'PGADMIN_DEFAULT_EMAIL=pgadmin4@pgadmin.org' -e 'PGADMIN_DEFAULT_PASSWORD=admin' -d --name pgadmin4 dpage/pgadmin4



Docker logs <container_name>

Docker logs c44e63e43539af533af426f03731fd8d113d533b667e546ca4aa35a441b07370





Docker stop container_name

List down all container regarding of there running status
Ps -a 


Docker exec -it <container_name> /bin/sh

Created —username=root —owner=root simple_bank

History | grep “docker run”

