create_swagger:
	swag init -g cmd/backend/main.go -o backend/docs -d backend
	swag init -g cmd/cloud/main.go -o cloud/docs -d cloud
