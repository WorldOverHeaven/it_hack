ROOT_FOLDER = ${PWD}

create_swagger:
	swag init -g cmd/backend/main.go -o backend/docs -d backend
	swag init -g cmd/cloud/main.go -o cloud/docs -d cloud

run:
	set -ex
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d back_db
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d cloud_db
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d hack_backend
	sudo docker-compose -f $(ROOT_FOLDER)/deploy/docker-compose.yaml -p deploy up --build -d hack_cloud