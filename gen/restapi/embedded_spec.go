// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "for tap",
    "title": "Okaimono",
    "version": "0.0.2"
  },
  "paths": {
    "/card": {
      "post": {
        "summary": "会員証を新規に発行するAPI",
        "operationId": "CreateMembershipCard",
        "parameters": [
          {
            "name": "RegisterMemberRequest",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "companyId"
              ],
              "properties": {
                "companyId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "成功時",
            "schema": {
              "$ref": "#/definitions/Member"
            },
            "headers": {
              "location": {
                "type": "string",
                "description": "発行された会員証のURL: /card/${memberId}"
              }
            }
          },
          "default": {
            "description": "失敗時",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Member": {
      "type": "object",
      "properties": {
        "memberId": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "for tap",
    "title": "Okaimono",
    "version": "0.0.2"
  },
  "paths": {
    "/card": {
      "post": {
        "summary": "会員証を新規に発行するAPI",
        "operationId": "CreateMembershipCard",
        "parameters": [
          {
            "name": "RegisterMemberRequest",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "companyId"
              ],
              "properties": {
                "companyId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "成功時",
            "schema": {
              "$ref": "#/definitions/Member"
            },
            "headers": {
              "location": {
                "type": "string",
                "description": "発行された会員証のURL: /card/${memberId}"
              }
            }
          },
          "default": {
            "description": "失敗時",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Member": {
      "type": "object",
      "properties": {
        "memberId": {
          "type": "string"
        }
      }
    }
  }
}`))
}
