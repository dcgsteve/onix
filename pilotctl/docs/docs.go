// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "gatblau",
            "url": "http://onix.gatblau.org/",
            "email": "onix@gatblau.org"
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
        "/": {
            "get": {
                "description": "Checks that Artie's HTTP server is listening on the required port.\nUse a liveliness probe.\nIt does not guarantee the server is ready to accept calls.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Check that Artie's HTTP API is live",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admission": {
            "get": {
                "description": "get a list of keys of the hosts admitted into service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admission"
                ],
                "summary": "Get Host Admissions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "creates a new or updates an existing host admission by allowing to specify active status and search tags",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Admission"
                ],
                "summary": "Create or Update a Host Admission",
                "parameters": [
                    {
                        "description": "the admission to be set",
                        "name": "command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Admission"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cmd": {
            "get": {
                "description": "get all command definitions",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Command"
                ],
                "summary": "Get All Command definitions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "creates a new or updates an existing command definition",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Command"
                ],
                "summary": "Create or Update a Command",
                "parameters": [
                    {
                        "description": "the command definition",
                        "name": "command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Cmd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cmd/{id}": {
            "get": {
                "description": "get a specific a command definition",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Command"
                ],
                "summary": "Get a Command definition",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the unique id for the command to retrieve",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/host": {
            "get": {
                "description": "Returns a list of remote hosts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Host"
                ],
                "summary": "Get All Hosts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/job": {
            "get": {
                "description": "get all jobs",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "Get All Jobs Information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new job for execution on one or more remote hosts",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "Create a Job",
                "parameters": [
                    {
                        "description": "the job definition",
                        "name": "command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Cmd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/job/{id}": {
            "get": {
                "description": "get a specific a job information",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "Get Job Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the unique id for the job to retrieve",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/log": {
            "post": {
                "description": "log host events (e.g. up, down, connected, disconnected)",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Host"
                ],
                "summary": "Log Events",
                "parameters": [
                    {
                        "description": "the host logs to post",
                        "name": "logs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/core.Event"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/log/{host-id}": {
            "get": {
                "description": "get log host events (e.g. up, down, connected, disconnected) by specific host",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Host"
                ],
                "summary": "Get Events by Host",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the unique key for the host",
                        "name": "host-key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/package": {
            "get": {
                "description": "get a list of packages in the backing Artisan registry",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Registry"
                ],
                "summary": "Get Artisan Packages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/package/{name}/api": {
            "get": {
                "description": "get a list of exported functions and inputs for the specified package",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Registry"
                ],
                "summary": "Get the API of an Artisan Package",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the fully qualified name of the artisan package having the required API",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping/{host-key}": {
            "post": {
                "description": "submits a ping from a host to the control plane",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Host"
                ],
                "summary": "Ping",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the unique key for the host",
                        "name": "host-key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region": {
            "get": {
                "description": "get a list of regions where hosts are deployed",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Region"
                ],
                "summary": "Get Regions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/region/{region-key}/location": {
            "get": {
                "description": "get a list of locations within a particular region",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Region"
                ],
                "summary": "Get Locations by Region",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registers a new host and its technical details with the service",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Host"
                ],
                "summary": "Register a Host",
                "parameters": [
                    {
                        "description": "the host registration configuration",
                        "name": "registration-info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Registration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.Admission": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "key": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "core.Cmd": {
            "type": "object",
            "properties": {
                "function": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "input": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "package": {
                    "type": "string"
                }
            }
        },
        "core.Event": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "type": {
                    "description": "0: host up, 1: host down, 2: network up, 3: network down",
                    "type": "integer"
                }
            }
        },
        "core.Registration": {
            "type": "object",
            "properties": {
                "cpus": {
                    "type": "integer"
                },
                "hostname": {
                    "type": "string"
                },
                "machine_id": {
                    "description": "github.com/denisbrodbeck/machineid",
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "total_memory": {
                    "type": "number"
                },
                "virtual": {
                    "type": "boolean"
                }
            }
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
	Version:     "0.0.4",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Onix Remote Host",
	Description: "Remote Ctrl Service for Onix Pilot",
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
	swag.Register(swag.Name, &s{})
}