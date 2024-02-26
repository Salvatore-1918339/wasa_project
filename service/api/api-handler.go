package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	// ? Session
	rt.router.POST("/session", rt.wrap(rt.doLoginHandler))

	// ? User
	rt.router.GET("/Users", rt.wrap(rt.searchUsers))
	rt.router.GET("/Users/:id", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/Users/:id", rt.wrap(rt.setMyUserName))

	// TODO
	// ? Stream
	rt.router.GET("/Users/:id/home", rt.wrap(rt.getMyStream))

	// ? Follower
	rt.router.PUT("/Users/:id/followers/:follower_id", rt.wrap(rt.followUser))
	rt.router.DELETE("/Users/:id/followers/:follower_id", rt.wrap(rt.unfollowUser))

	// ? Ban
	rt.router.PUT("/Users/:id/banned_users/:banned_user_id", rt.wrap(rt.putBan))
	rt.router.DELETE("/Users/:id/banned_users/:banned_user_id", rt.wrap(rt.deleteBan))

	// ? Photo
	rt.router.POST("/Users/:id/photos", rt.wrap(rt.uploadPhoto))
	rt.router.GET("/Users/:id/Photos/:photo_id", rt.wrap(rt.getPhoto))
	rt.router.DELETE("/Users/:id/Photos/:photo_id", rt.wrap(rt.deletePhoto))

	// ? Comments
	rt.router.POST("/Users/:id/Photos/:photo_id/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/Users/:id/Photos/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))

	// ? Likes
	rt.router.PUT("/Users/:id/Photos/:photo_id/Likes", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/Users/:id/Photos/:photo_id/Likes/:like_id", rt.wrap(rt.unlikePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
