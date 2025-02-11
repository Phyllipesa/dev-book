package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindAllWithNameOrNick,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.FindUserById,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequestAuthentication: true,
	},
}
