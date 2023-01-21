// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Lev Yuvensky",
            "email": "levayu22@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chats/add": {
            "post": {
                "description": "Create a new chat with name and list of users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Create new chat",
                "operationId": "createChat",
                "parameters": [
                    {
                        "description": "Parameters for creating a chat.",
                        "name": "createChatObject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateChatRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ChatId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/chats/get": {
            "post": {
                "description": "Get a list of user chats by user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Get a list of user chats",
                "operationId": "getUserChats",
                "parameters": [
                    {
                        "description": "Parameters for getting user chats.",
                        "name": "getUserChatsObject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GetUserChatsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.GetUserChatsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/messages/add": {
            "post": {
                "description": "Create a new message with chat's id, author's id and text.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "summary": "Create new message",
                "operationId": "createMessage",
                "parameters": [
                    {
                        "description": "Parameters for creating a message.",
                        "name": "createMessageObject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateMessageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.MessageId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/messages/get": {
            "post": {
                "description": "Get a list of chat messages by chat ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "summary": "Get a list of chat messages",
                "operationId": "GetChatMessages",
                "parameters": [
                    {
                        "description": "Parameters for getting chat messages.",
                        "name": "getChatsMessagesObject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GetChatMessagesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Message"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/users/add": {
            "post": {
                "description": "Create a new user with username.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create new user",
                "operationId": "createUser",
                "parameters": [
                    {
                        "description": "Parameters for creating a user.",
                        "name": "createUserObject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.ErrorJSON"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apperror.ErrorJSON": {
            "type": "object",
            "properties": {
                "developerMessage": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ChatId": {
            "type": "object",
            "properties": {
                "chatId": {
                    "type": "integer"
                }
            }
        },
        "models.CreateChatRequest": {
            "type": "object",
            "required": [
                "name",
                "users"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.CreateMessageRequest": {
            "type": "object",
            "required": [
                "author",
                "chat",
                "text"
            ],
            "properties": {
                "author": {
                    "type": "integer"
                },
                "chat": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserRequest": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "models.GetChatMessagesRequest": {
            "type": "object",
            "required": [
                "chat"
            ],
            "properties": {
                "chat": {
                    "type": "integer"
                }
            }
        },
        "models.GetUserChatsRequest": {
            "type": "object",
            "required": [
                "user"
            ],
            "properties": {
                "user": {
                    "type": "integer"
                }
            }
        },
        "models.GetUserChatsResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "integer"
                },
                "chat": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.MessageId": {
            "type": "object",
            "properties": {
                "messageId": {
                    "type": "integer"
                }
            }
        },
        "models.UserId": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Avito.tech's task for backend trainee (2020 year)",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
