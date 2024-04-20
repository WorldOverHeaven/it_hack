create_swagger:
	swag init -g cmd/backend/main.go -o backend/docs -d backend
	swag init -g cmd/cloud/main.go -o cloud/docs -d cloud

run:
	sudo docker-compose -f ./deploy/docker-compose.yaml -p deploy up -d back_db
	sudo docker-compose -f ./deploy/docker-compose.yaml -p deploy up -d cloud_db
	sudo docker-compose -f ./deploy/docker-compose.yaml -p deploy up -d hack_backend
	sudo docker-compose -f ./deploy/docker-compose.yaml -p deploy up -d hack_cloud
