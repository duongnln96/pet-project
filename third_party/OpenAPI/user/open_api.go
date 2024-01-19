package user

import "github.com/getkin/kin-openapi/openapi3"

// NewOpenAPI3 instantiates the OpenAPI specification for this service
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "User API",
			Description: "REST APIs used for interacting with the User Service",
			Version:     "0.0.0",
			Contact: &openapi3.Contact{
				URL: "https://github.com/duongnln96/blog-realworld",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://127.0.0.1",
			},
			&openapi3.Server{
				URL:         "Docker environment",
				Description: "http://0.0.0.0",
			},
		},

		Components: initComponent(),
		Paths:      initPath(),
	}

	return swagger
}

func initComponent() *openapi3.Components {

	openAPI3Component := &openapi3.Components{}

	// init schema
	openAPI3Component.Schemas = openapi3.Schemas{
		"UserDTO": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("name", openapi3.NewStringSchema()).
				WithProperty("bio", openapi3.NewStringSchema()).
				WithProperty("created_date", openapi3.NewDateTimeSchema()).
				WithProperty("updated_date", openapi3.NewDateTimeSchema()),
		),
	}

	// init body request
	openAPI3Component.RequestBodies = openapi3.RequestBodies{
		"RegisterUserDTO": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Body request").
				WithRequired(true).
				WithJSONSchema(
					openapi3.NewSchema().
						WithProperty("name", openapi3.NewStringSchema().WithMaxLength(500)).
						WithProperty("bio", openapi3.NewStringSchema().WithMaxLength(500)).
						WithProperty("email", openapi3.NewStringSchema().WithMaxLength(500)).
						WithProperty("password", openapi3.NewStringSchema()),
				),
		},
		"LoginUserDTO": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Body request").
				WithRequired(true).
				WithJSONSchema(
					openapi3.NewSchema().
						WithProperty("email", openapi3.NewStringSchema().WithMaxLength(500)).
						WithProperty("password", openapi3.NewStringSchema()),
				),
		},
		"UpdateUserDTO": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Body request").
				WithRequired(true).
				WithJSONSchema(
					openapi3.NewSchema().
						WithProperty("id", openapi3.NewUUIDSchema()).
						WithProperty("name", openapi3.NewStringSchema().WithMaxLength(500)).WithNullable().
						WithProperty("bio", openapi3.NewStringSchema().WithMaxLength(500)).WithNullable().
						WithProperty("email", openapi3.NewStringSchema().WithMaxLength(500)).WithNullable().
						WithProperty("password", openapi3.NewStringSchema()).WithNullable(),
				),
		},
	}

	// init body response
	openAPI3Component.Responses = openapi3.ResponseBodies{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Body response").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithProperty("code", openapi3.NewIntegerSchema()).
						WithProperty("error_code", openapi3.NewStringSchema().WithNullable()).
						WithProperty("error_msg", openapi3.NewStringSchema().WithNullable()),
				)),
		},
		// user
		"UserRegister": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Body response").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithProperty("code", openapi3.NewIntegerSchema()).
						WithPropertyRef("data", &openapi3.SchemaRef{
							Ref: "#/components/schemas/UserDTO",
						}),
				)),
		},
		"UserDetail": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Body response").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithProperty("code", openapi3.NewIntegerSchema()).
						WithPropertyRef("data", &openapi3.SchemaRef{
							Ref: "#/components/schemas/UserDTO",
						}),
				)),
		},
		"UserUpdate": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Body response").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithProperty("code", openapi3.NewIntegerSchema()).
						WithPropertyRef("data", &openapi3.SchemaRef{
							Ref: "#/components/schemas/UserDTO",
						}),
				)),
		},

		// profile
	}

	return openAPI3Component
}

func initPath() *openapi3.Paths {

	openAPI3Paths := openapi3.NewPaths()

	// user paths
	openAPI3Paths.Set("/api/v1/user/register", &openapi3.PathItem{
		Put: &openapi3.Operation{
			OperationID: "user/register",
			RequestBody: &openapi3.RequestBodyRef{
				Ref: "#/components/requestBodies/RegisterUserDTO",
			},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(200, &openapi3.ResponseRef{
					Ref: "#/components/responses/UserRegister",
				}),
				openapi3.WithStatus(400, &openapi3.ResponseRef{
					Ref: "#/components/responses/ErrorResponse",
				}),
				openapi3.WithStatus(500, &openapi3.ResponseRef{
					Ref: "#/components/responses/ErrorResponse",
				}),
			),
		},
	})

	openAPI3Paths.Set("/api/v1/user/:user_id", &openapi3.PathItem{
		Get: &openapi3.Operation{
			OperationID: "user/user_id",
			Parameters: []*openapi3.ParameterRef{
				{
					Value: openapi3.NewPathParameter("user_id").
						WithSchema(openapi3.NewUUIDSchema()),
				},
			},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(200, &openapi3.ResponseRef{
					Ref: "#/components/responses/UserDetail",
				}),
				openapi3.WithStatus(400, &openapi3.ResponseRef{
					Ref: "#/components/responses/ErrorResponse",
				}),
				openapi3.WithStatus(500, &openapi3.ResponseRef{
					Ref: "#/components/responses/ErrorResponse",
				}),
			),
		},
	})

	openAPI3Paths.Set("/api/v1/user/update", &openapi3.PathItem{
		Post: &openapi3.Operation{
			OperationID: "user/update",
			RequestBody: &openapi3.RequestBodyRef{
				Ref: "#/components/requestBodies/UpdateUserDTO",
			},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(200, &openapi3.ResponseRef{
					Ref: "#/components/responses/UserUpdate",
				}),
				openapi3.WithStatus(400, &openapi3.ResponseRef{
					Ref: "#/components/responses/ErrorResponse",
				}),
				openapi3.WithStatus(500, &openapi3.ResponseRef{
					Ref: "#/components/responses/ErrorResponse",
				}),
			),
		},
	})

	// profile path

	return openAPI3Paths
}
