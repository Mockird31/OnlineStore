.PHONY: protogen-all docker-up docker-remove

protogen-all:
	protoc -I proto proto/**/*.proto --go_out=gen --go-grpc_out=gen

docker-up:
	cd deploy/ && make deploy

docker-remove:
	-docker stop $$(docker ps -q)             
	-docker rm -f $$(docker ps -aq)           
	-docker rmi -f $$(docker images -q)
	-docker image prune -f