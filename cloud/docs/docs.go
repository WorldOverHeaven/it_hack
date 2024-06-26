// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth_user": {
            "post": {
                "description": "AuthUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "AuthUser",
                "parameters": [
                    {
                        "description": "AuthUser",
                        "name": "AuthUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelscloud.AuthUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.AuthUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    }
                }
            }
        },
        "/create_user": {
            "post": {
                "description": "CreateUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "CreateUser",
                        "name": "CreateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelscloud.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    }
                }
            }
        },
        "/get_payload": {
            "post": {
                "description": "GetPayload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetPayload",
                "parameters": [
                    {
                        "description": "GetPayload",
                        "name": "GetPayload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelscloud.GetPayloadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.GetPayloadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    }
                }
            }
        },
        "/put_payload": {
            "post": {
                "description": "PutPayload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PutPayload",
                "parameters": [
                    {
                        "description": "PutPayload",
                        "name": "PutPayload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelscloud.PutPayloadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.PutPayloadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/modelscloud.BadRequestResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "modelscloud.AuthUserRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "modelscloud.AuthUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "modelscloud.BadRequestResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                }
            }
        },
        "modelscloud.CreateUserRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "modelscloud.CreateUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "modelscloud.GetPayloadRequest": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "modelscloud.GetPayloadResponse": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "string"
                }
            }
        },
        "modelscloud.PutPayloadRequest": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "modelscloud.PutPayloadResponse": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
