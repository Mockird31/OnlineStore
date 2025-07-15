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

mock-all:
	mockgen -source=gen/auth/auth_grpc.pb.go -destination=mocks/mock_auth_client.go -package=mocks AuthServiceClient
	mockgen -source=gen/user/user_grpc.pb.go -destination=mocks/mock_user_client.go -package=mocks UserServiceClient

test:
	./scripts/test.sh

clean:
	$(RM) -rf *.out *.html *.tmp *.txt