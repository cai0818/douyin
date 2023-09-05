package router

import (
	"github.com/gin-gonic/gin"
	"simpledouyin/Controller/Base"
	"simpledouyin/Controller/Interaction"
	"simpledouyin/Controller/Socialize"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./static")
	//apiRouter := r.Group("/douyin")

	r.GET("/douyin/feed/", Base.Feed)
	r.GET("/douyin/publish/list/", Base.PublishList)
	r.GET("/douyin/user/", Base.UserInfo)
	r.POST("/douyin/user/login/", Base.Login)
	r.POST("/douyin/user/register/", Base.Register)
	r.POST("/douyin/publish/action/", Base.UploadVideo)

	r.GET("/douyin/favorite/list/", Interaction.Favorite)
	r.POST("/douyin/favorite/action/", Interaction.Like)
	r.POST("/douyin/comment/action/", Interaction.Conmment)
	r.GET("/douyin/comment/list/", Interaction.CommentList)

	r.POST("/douyin/relation/action/", Socialize.Follow)
	r.GET("/douyin/relation/follow/list/", Socialize.FollowList)
	r.GET("/douyin/relation/follower/list/", Socialize.FanList)
	r.GET("/douyin/relation/friend/list/", Socialize.FriendList)
	r.GET("/douyin/message/chat/", Socialize.MessageList)
	r.POST("/douyin/message/action/", Socialize.AddMessage)
}
