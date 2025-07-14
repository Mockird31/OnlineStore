ifneq (,$(wildcard ../.env))
	include ../.env
	export
endif

.PHONY: deploy

deploy:
	docker-compose up -d --build