ifneq (,$(wildcard ../.env))
	include ../.env
	export
endif

.PHONY: deploy

deploy:
	docker-compose up -d --build

build-images:
	eval $$(minikube docker-env)
	docker build -t api-service:latest -f images/api/Dockerfile ..
	docker build -t auth-service:latest -f images/auth/Dockerfile ..
	docker build -t user-service:latest -f images/user/Dockerfile ..

load-images:
	minikube image load auth-service
	minikube image load user-service
	minikube image load api-service

kuber-up:
	$(MAKE) build-images
	kubectl apply -f kubernetes/namespace.yaml
	kubectl create secret generic app-secrets \
	  --from-env-file=../.env \
  		--namespace=online-store
	kubectl apply -f kubernetes/configmap.yaml
	kubectl apply -f kubernetes/persistent-volume.yaml
	kubectl apply -f kubernetes/redis-deployment.yaml
	kubectl apply -f kubernetes/postgres-deployment.yaml
	kubectl apply -f kubernetes/auth-service.yaml
	kubectl apply -f kubernetes/user-service.yaml
	kubectl apply -f kubernetes/api-service.yaml
	minikube addons enable ingress
	echo "Waiting for ingress controller to be ready..."
	sleep 30
	kubectl wait --namespace ingress-nginx \
		--for=condition=ready pod \
		--selector=app.kubernetes.io/component=controller \
		--timeout=120s || true
	kubectl apply -f kubernetes/ingress.yaml

kuber-down:
	kubectl delete namespace online-store
