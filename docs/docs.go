// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "email": "a.kapitonov@limehd.tv"
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
        "/internal/api/v1/event": {
            "post": {
                "description": "Method receive message and send to queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Receive event fields",
                "parameters": [
                    {
                        "description": "request event",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Event"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "no content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/internal/api/v1/event-batch": {
            "post": {
                "description": "Method receive messages and send to queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Receive event fields with batch mode",
                "parameters": [
                    {
                        "description": "request events",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Event"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "no content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/internal/api/v1/event-debug": {
            "post": {
                "description": "Method check event message is valid or not",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Debug event fields",
                "parameters": [
                    {
                        "description": "request debug event",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Event"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/gateway.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Event": {
            "type": "object",
            "properties": {
                "android_id": {
                    "type": "string"
                },
                "app": {
                    "type": "string"
                },
                "app_package_name": {
                    "type": "string"
                },
                "app_version_name": {
                    "type": "string"
                },
                "application_id": {
                    "type": "integer"
                },
                "appmetrica_device_id": {
                    "type": "string"
                },
                "browser": {
                    "type": "string"
                },
                "browser_version": {
                    "type": "string"
                },
                "connection_type": {
                    "type": "string"
                },
                "cookie_enabled": {
                    "type": "boolean"
                },
                "device_id": {
                    "type": "string"
                },
                "device_locale": {
                    "type": "string"
                },
                "device_manufacturer": {
                    "type": "string"
                },
                "device_model": {
                    "type": "string"
                },
                "device_type": {
                    "type": "string"
                },
                "event_datetime": {
                    "$ref": "#/definitions/domain.EventDatetime"
                },
                "event_json": {
                    "type": "string"
                },
                "event_name": {
                    "type": "string"
                },
                "event_timestamp": {
                    "type": "integer"
                },
                "google_aid": {
                    "type": "string"
                },
                "hardware_or_gui": {
                    "type": "string"
                },
                "installation_id": {
                    "type": "string"
                },
                "ios_ifa": {
                    "type": "string"
                },
                "ios_ifv": {
                    "type": "string"
                },
                "js_enabled": {
                    "type": "boolean"
                },
                "mcc": {
                    "type": "string"
                },
                "mnc": {
                    "type": "string"
                },
                "operator_name": {
                    "type": "string"
                },
                "os_name": {
                    "type": "string"
                },
                "os_version": {
                    "type": "string"
                },
                "physical_screen_height": {
                    "type": "integer"
                },
                "physical_screen_width": {
                    "type": "integer"
                },
                "platform": {
                    "type": "string"
                },
                "profile_id": {
                    "type": "string"
                },
                "referer": {
                    "type": "string"
                },
                "screen_aspect_ratio": {
                    "type": "string"
                },
                "screen_height": {
                    "type": "integer"
                },
                "screen_orientation": {
                    "type": "string"
                },
                "screen_weight": {
                    "type": "integer"
                },
                "sdk_version": {
                    "type": "integer"
                },
                "session_id": {
                    "type": "string"
                },
                "timezone": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                },
                "uniq_id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "user_agent": {
                    "type": "string"
                },
                "utm_campaign": {
                    "type": "string"
                },
                "utm_content": {
                    "type": "string"
                },
                "utm_medium": {
                    "type": "string"
                },
                "utm_source": {
                    "type": "string"
                },
                "utm_term": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                },
                "xlhd_agent": {
                    "type": "string"
                }
            }
        },
        "domain.EventDatetime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "gateway.ErrorMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "request_body": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "lm.limehd.tv",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "OwnMetrics aka AppMetrica gateway service",
	Description:      "Under construct",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
