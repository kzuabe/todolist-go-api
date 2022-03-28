// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/v1/tasks": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "ユーザのタスクを複数件取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "タスク取得",
                "parameters": [
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "タスクステータス 0: 未着手 1: 完了",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "ユーザのタスクを追加する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "タスク追加",
                "parameters": [
                    {
                        "description": "Payload Description",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            }
        },
        "/v1/tasks/{id}": {
            "get": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "ユーザのタスクを1件取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "タスク取得（1件）",
                "parameters": [
                    {
                        "type": "string",
                        "description": "タスクID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "ユーザのタスクを更新する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "タスク更新",
                "parameters": [
                    {
                        "type": "string",
                        "description": "タスクID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Description",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "TokenAuth": []
                    }
                ],
                "description": "ユーザのタスクを削除する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "タスク削除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "タスクID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Task": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "タスク説明文",
                    "type": "string"
                },
                "id": {
                    "description": "タスクID（自動で生成されるUUID）",
                    "type": "string"
                },
                "status": {
                    "description": "タスクのステータス（0: 未着手, 1: 完了）",
                    "type": "integer"
                },
                "title": {
                    "description": "タスクタイトル",
                    "type": "string"
                },
                "user_id": {
                    "description": "ユーザID",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "TokenAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "TodoList API",
	Description: "TODOリストアプリのRESTfulAPI（Go実装）",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
