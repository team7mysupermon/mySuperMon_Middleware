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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Login/{Username}/{Password}": {
            "get": {
                "description": "this is a request to give the middleware user information. this will allow the middleware to set up the authentication token need to start and stop the recording.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mySuperMon Middleware"
                ],
                "summary": "Send middleware user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": ":Username",
                        "name": "Username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": ":Password",
                        "name": "Password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User information is accepted"
                    },
                    "400": {
                        "description": "User information is not correct"
                    }
                }
            }
        },
        "/Start/{Usecase}/{Appiden}": {
            "get": {
                "description": "This endpoint is to stop a recording and needs a usecase and a applicationIdentifier as parameters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mySuperMon Middleware"
                ],
                "summary": "Start a recording",
                "parameters": [
                    {
                        "type": "string",
                        "description": ":Usecase",
                        "name": "Usecase",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": ":Appiden",
                        "name": "Appiden",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recording has started"
                    },
                    "400": {
                        "description": "Wrong parameters"
                    },
                    "500": {
                        "description": "Connection error to mySuperMon webservice"
                    }
                }
            }
        },
        "/Stop/{Usecase}/{Appiden}": {
            "get": {
                "description": "This endpoint is to stop a recording and needs a usecase and a applicationIdentifier as parameters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mySuperMon Middleware"
                ],
                "summary": "Stop a recording",
                "parameters": [
                    {
                        "type": "string",
                        "description": ":Usecase",
                        "name": "Usecase",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": ":Appiden",
                        "name": "Appiden",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recording has stopped"
                    },
                    "400": {
                        "description": "Wrong parameters"
                    },
                    "500": {
                        "description": "Connection error to mySuperMon webservice"
                    }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
