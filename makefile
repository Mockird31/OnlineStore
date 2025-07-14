ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.PHONY: protogen-all 

protogen-all:
	protoc -I proto proto/**/*.proto --go_out=gen --go-grpc_out=gen

docker-up:
	cd deploy/ && make deploy

migrate_up:
	tern migrate -c db/migrations/tern.conf --migrations db/migrations

migrate_down:
	tern migrate -c db/migrations/tern.conf --migrations db/migrations -d 1

populate:
	make migrate_up

docker-remove:
	-docker stop $$(docker ps -q)             
	-docker rm -f $$(docker ps -aq)           
	-docker rmi -f $$(docker images -q)
	-docker image prune -f