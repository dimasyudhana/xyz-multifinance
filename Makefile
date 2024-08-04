.PHONY: create-migration cert secure-ssl docker docker-compose-up docker-compose-down

create-migration:
	migrate create -ext sql -dir migration -seq customers
	migrate create -ext sql -dir migration -seq product
	migrate create -ext sql -dir migration -seq transaction
	migrate create -ext sql -dir migration -seq transaction_products

cert:
	@echo Generating SSL certificates
	cd ./deployment/ssl && sh instructions.sh
	cd ..
	cd ./deployment/nginx && openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx-selfsigned.key -out nginx-selfsigned.crt

secure-ssl:
	chmod 400 ca.key nginx-selfsigned.key
	chmod 444 ca.crt nginx-selfsigned.crt nginx-selfsigned.pem

docker:
	docker images | grep backend-service
	docker rmi --force backend-service
	go mod tidy
	docker build -t backend-service .

docker-compose-up:
	docker compose -f docker-compose-service.yml up && docker compose up -d && docker compose logs -f be
