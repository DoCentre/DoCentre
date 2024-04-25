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
                    "document"
                ],
                "summary": "Create document",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.createDocumentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.createDocumentSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.createDocumentInvalidResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.createDocumentFailedResponse"
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
                    "health"
                ],
                "summary": "Check health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.healthResponse"
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
                            "$ref": "#/definitions/controllers.userCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.userCreateSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.userCreateExistedResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
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
        "controllers.createDocumentFailedResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Failed to create document"
                }
            }
        },
        "controllers.createDocumentInvalidResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Invalid request body"
                }
            }
        },
        "controllers.createDocumentRequest": {
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
        "controllers.createDocumentSuccessResponse": {
            "type": "object",
            "properties": {
                "document_id": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "controllers.userCreateExistedResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "User/Email already exists"
                }
            }
        },
        "controllers.userCreateRequest": {
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
        "controllers.userCreateSuccessResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/controllers.UserDto"
                }
            }
        },
        "router.healthResponse": {
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
