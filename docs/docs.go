// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "User login route",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authDao.LoginDao"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in",
                        "schema": {
                            "$ref": "#/definitions/authDao.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid email or password",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "description": "User logout route",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "User logged out"
                    }
                }
            }
        },
        "/auth/refresh": {
            "get": {
                "description": "Refresh access token route",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh access token",
                "responses": {
                    "200": {
                        "description": "Access token refreshed",
                        "schema": {
                            "$ref": "#/definitions/authDao.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid refresh token",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "User registration route",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "User registration",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authDao.RegisterDao"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered",
                        "schema": {
                            "$ref": "#/definitions/authDao.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        },
        "/organizations": {
            "get": {
                "description": "List organizations route",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organizations"
                ],
                "summary": "List organizations",
                "responses": {
                    "200": {
                        "description": "List of organizations",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessDataResponse-array_model_Organizations"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create organization route",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organizations"
                ],
                "summary": "Create organization",
                "parameters": [
                    {
                        "description": "Organization creation",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/organizationDao.OrganizationInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Organization created",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessDataResponse-model_Organizations"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    },
                    "403": {
                        "description": "Owner ID must be the same as the authenticated user",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        },
        "/organizations/{orgId}": {
            "get": {
                "description": "Get organization by slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organizations"
                ],
                "summary": "Get organization",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization Slug",
                        "name": "orgId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessDataResponse-model_Organizations"
                        }
                    },
                    "404": {
                        "description": "Organization not found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        },
        "/organizations/{orgSlug}/teams": {
            "get": {
                "description": "Get list of teams",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get list of teams",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization slug",
                        "name": "orgSlug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of teams",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessDataResponse-array_model_Teams"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch teams",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Create a team",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization slug",
                        "name": "orgSlug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Team data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/teamsDao.CreateTeamInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Team created",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessDataResponse-model_Teams"
                        }
                    },
                    "400": {
                        "description": "Slug already exists",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    },
                    "500": {
                        "description": "Failed to create team",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "User info",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessDataResponse-model_Users"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch user info",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse-string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authDao.AuthResponse": {
            "type": "object",
            "properties": {
                "expirationDurationSeconds": {
                    "type": "integer"
                }
            }
        },
        "authDao.LoginDao": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "authDao.RegisterDao": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Organizations": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Teams": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization_id": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Users": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_reset_at": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "verified_at": {
                    "type": "string"
                }
            }
        },
        "organizationDao.OrganizationInput": {
            "type": "object",
            "required": [
                "name",
                "owner_id",
                "slug"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse-string": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.SuccessDataResponse-array_model_Organizations": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Organizations"
                    }
                }
            }
        },
        "response.SuccessDataResponse-array_model_Teams": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Teams"
                    }
                }
            }
        },
        "response.SuccessDataResponse-model_Organizations": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Organizations"
                }
            }
        },
        "response.SuccessDataResponse-model_Teams": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Teams"
                }
            }
        },
        "response.SuccessDataResponse-model_Users": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Users"
                }
            }
        },
        "teamsDao.CreateTeamInput": {
            "type": "object",
            "required": [
                "name",
                "slug"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
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
