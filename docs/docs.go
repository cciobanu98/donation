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
        "termsOfService": "http://swagger.io/terms/",
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
        "/project": {
            "get": {
                "description": "get all projects",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "summary": "get projects",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "create a new project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "summary": "create a new project",
                "parameters": [
                    {
                        "description": "Project Info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateProjectInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/project/{project_id}": {
            "get": {
                "description": "Get project using id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "summary": "Get project by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "delete project using id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "project"
                ],
                "summary": "delete project by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateProjectInput": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Project API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}