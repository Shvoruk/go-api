// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://example.com/support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/animals": {
            "get": {
                "description": "Retrieves a list of animals filtered by the specified category.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animals"
                ],
                "summary": "Retrieve animals by category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Animal category",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of animals",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Animal"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new animal record for the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animals"
                ],
                "summary": "Create a new animal record",
                "parameters": [
                    {
                        "description": "Animal data",
                        "name": "animal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Animal"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ID of the newly created animal",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/animals/my": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves a list of animals that belong to the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animals"
                ],
                "summary": "Retrieve authenticated user's animals",
                "responses": {
                    "200": {
                        "description": "List of user's animals",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Animal"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/animals/{id}": {
            "get": {
                "description": "Gets detailed information about an animal using its numeric ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animals"
                ],
                "summary": "Retrieve a specific animal by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Animal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/types.Animal"
                        }
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Animal not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deletes an animal record if it belongs to the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "animals"
                ],
                "summary": "Delete an animal record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Animal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid id parameter",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Animal not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "Authenticates a user using email and password, and returns a JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate a user",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token generated",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/signup": {
            "post": {
                "description": "Registers a new user, hashes the password, creates a JWT token, and returns it.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User signup data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Token generated",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request or email already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Animal": {
            "type": "object",
            "properties": {
                "age_in_month": {
                    "type": "integer"
                },
                "breed": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "contact_info": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "types.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "types.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "BeastyBonds Go API",
	Description:      "API for managing animals and users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
