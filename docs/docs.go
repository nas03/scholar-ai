package docs

import (
	"github.com/swaggo/swag"
)

// NOTE: This file provides a minimal Swagger spec without requiring the swag CLI.
// If you later run `swag init`, it will overwrite this with a generated version.

const swaggerSpec = `{
  "swagger": "2.0",
  "info": {
    "title": "Scholar AI Backend API",
    "version": "0.1.0",
    "description": "REST API for Scholar AI backend services."
  },
  "basePath": "/api/v1",
  "schemes": ["http", "https"],
  "paths": {}
}`

type s struct{}

func (s *s) ReadDoc() string { return swaggerSpec }

func init() {
	swag.Register(swag.Name, &s{})
}
