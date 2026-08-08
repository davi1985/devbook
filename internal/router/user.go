package router

import (
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Func:                  controller.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Func:                  controller.GetAllUsers,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Func:                  controller.GetUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Func:                  controller.UpdateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Func:                  controller.DeleteUser,
		RequireAuthentication: false,
	},
}
