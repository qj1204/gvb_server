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
        "/api/advert": {
            "get": {
                "description": "广告列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "广告列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示条数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "创建广告",
                "parameters": [
                    {
                        "description": "表示多个参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advert.AdvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "删除广告",
                "parameters": [
                    {
                        "description": "广告id列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/advert/:id": {
            "put": {
                "description": "更新广告",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "广告管理"
                ],
                "summary": "更新广告",
                "parameters": [
                    {
                        "description": "广告的一些参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/advert.AdvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/image": {
            "get": {
                "description": "图片列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示条数",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/image_name": {
            "get": {
                "description": "图片名称列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片名称列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/image.ImageResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "advert.AdvertRequest": {
            "type": "object",
            "required": [
                "href",
                "image",
                "is_show",
                "title"
            ],
            "properties": {
                "href": {
                    "description": "广告链接",
                    "type": "string"
                },
                "image": {
                    "description": "广告图片",
                    "type": "string"
                },
                "is_show": {
                    "description": "是否显示",
                    "type": "boolean"
                },
                "title": {
                    "description": "广告标题",
                    "type": "string"
                }
            }
        },
        "image.ImageResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "models.RemoveRequest": {
            "type": "object",
            "properties": {
                "id_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "response.ListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {}
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gvb_server API文档",
	Description:      "gvb_server API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
