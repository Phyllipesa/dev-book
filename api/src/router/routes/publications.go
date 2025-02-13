package routes

import (
	"api/src/controllers"
	"net/http"
)

var routePublications = []Route{
	{
		URI:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequestAuthentication: true,
	},
	{
		URI:                   "/publications",
		Method:                http.MethodGet,
		Function:              controllers.FindPublications,
		RequestAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodGet,
		Function:              controllers.FindPublicationById,
		RequestAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequestAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequestAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/publications",
		Method:                http.MethodGet,
		Function:              controllers.FindPublicationsByUser,
		RequestAuthentication: true,
	},
}
