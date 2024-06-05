
# iniciando migrate
migrate create -ext=sql -dir=sql/migrations -seq init

-ext = extensão
-dir = dir aonde vai ser salvo
-seq = é como vai versiona o arquivo, no caso sequencia


# criar tabelas
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

# dropar tabelas
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

# sqlc
sqlc generate

# rodar sqlc no windows usando docker

docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate


# instalado via go get

sqlc generate
