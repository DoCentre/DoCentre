// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/document": {
            "post": {
                "description": "Create a new document that belongs to the author; the author has to be a existing user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Document"
                ],
                "summary": "Create document",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateDocument.requestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateDocument.successResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateDocument.invalidResponseBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateDocument.failedResponseBody"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Check health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.CheckHealth.responseBody"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "Create a new user; the user will be created with the identity \"user\".",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserCreate.requestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserCreate.successResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserCreate.existedResponseBody"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "post": {
                "description": "Get users with the same given username.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get users by username",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GetUsersByUsername.requestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetUsersByUsername.usersNotFoundResponseBody"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login a user with username and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserLogin.requestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserLogin.successResponseBody"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserLogin.userNotFoundResponseBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateDocument.failedResponseBody": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Failed to create document"
                }
            }
        },
        "controllers.CreateDocument.invalidResponseBody": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Invalid request body"
                }
            }
        },
        "controllers.CreateDocument.requestBody": {
            "type": "object",
            "required": [
                "author_id"
            ],
            "properties": {
                "author_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "controllers.CreateDocument.successResponseBody": {
            "type": "object",
            "properties": {
                "document_id": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "controllers.GetUsersByUsername.requestBody": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "controllers.GetUsersByUsername.successResponseBody": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.UserDto"
                    }
                }
            }
        },
        "controllers.GetUsersByUsername.usersNotFoundResponseBody": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "Users not found"
                },
                "users": {
                    "description": "Should always be empty.\nXXX: Consider removing the field.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.UserDto"
                    }
                }
            }
        },
        "controllers.UserCreate.existedResponseBody": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "User/Email already exists"
                }
            }
        },
        "controllers.UserCreate.requestBody": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "controllers.UserCreate.successResponseBody": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/controllers.UserDto"
                }
            }
        },
        "controllers.UserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@mail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "identity": {
                    "type": "string",
                    "example": "user"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "controllers.UserLogin.requestBody": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "controllers.UserLogin.successResponseBody": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/controllers.UserDto"
                }
            }
        },
        "controllers.UserLogin.userNotFoundResponseBody": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "User not found"
                },
                "user": {
                    "description": "Should always be nil.\nXXX: Consider removing the field; also swaggo fails to generate example with null value.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controllers.UserDto"
                        }
                    ]
                }
            }
        },
        "router.CheckHealth.responseBody": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
