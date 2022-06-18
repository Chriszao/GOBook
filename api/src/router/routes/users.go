package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Function:  controllers.InsertUser,
		IsPrivate: false,
	},
	{
		URI:       "/users",
		Method:    http.MethodGet,
		Function:  controllers.FetchUsers,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}",
		Method:    http.MethodGet,
		Function:  controllers.GetUserById,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}",
		Method:    http.MethodPut,
		Function:  controllers.UpdateUser,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}",
		Method:    http.MethodDelete,
		Function:  controllers.DeleteUser,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}/follow",
		Method:    http.MethodPost,
		Function:  controllers.FollowUser,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}/unfollow",
		Method:    http.MethodPost,
		Function:  controllers.UnFollowUser,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}/followers",
		Method:    http.MethodGet,
		Function:  controllers.FindFollowers,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}/following",
		Method:    http.MethodGet,
		Function:  controllers.FindFollowing,
		IsPrivate: true,
	},
	{
		URI:       "/users/{id}/reset-password",
		Method:    http.MethodPost,
		Function:  controllers.ResetPassword,
		IsPrivate: true,
	},
}
