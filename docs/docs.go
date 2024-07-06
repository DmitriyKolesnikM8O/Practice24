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
        "/auth": {
            "post": {
                "description": "auth by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Auth",
                "operationId": "auth",
                "parameters": [
                    {
                        "description": "Username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_DmitriyKolesnikM8O_Practice24_internal_auth.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete product by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "DeleteProduct",
                "operationId": "delete-product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "all products from table",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "GetProducts",
                "operationId": "get-product-by-id",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update product by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "UpdateProductByID",
                "operationId": "update-product",
                "parameters": [
                    {
                        "description": "Product information",
                        "name": "information",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_product.UpdateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "new product in table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create Product",
                "operationId": "create",
                "parameters": [
                    {
                        "description": "Product information",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_product.CreateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "one product from table by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "GetProductByID",
                "operationId": "get-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/report": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create report from table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create Report",
                "operationId": "report",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_DmitriyKolesnikM8O_Practice24_internal_auth.User": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "internal_product.CreateProduct": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "internal_product.UpdateProduct": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
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
	Host:             "0.0.0.0:1234",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Generate Report API",
	Description:      "API service to generate report",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
