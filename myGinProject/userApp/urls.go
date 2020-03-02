package userApp

import (
	"github.com/gin-gonic/gin"
)

func urls_init(r *gin.Engine) {

	authorized := r.Group("/")
	authorized.Use(AuthLogin)
	{
		//authorized.GET("/", hello)
		authorized.GET("logout",logout)
	}
	g := r.Group("/")
	{
		g.POST("login",login)
		g.POST("register",register)
		g.GET("profile",profile)
		g.POST("update",update)
		g.POST("updatePassword",updatePassword)
		g.GET("autologin",autologin)
		g.GET("getProblems",getProblems)
		g.GET("getOneProblem",getOneProblem)
		g.POST("submitCode",submitCode)
		g.GET("getHistory",getHistory)
		g.POST("searchSubmissions",searchSubmissions)
		g.POST("searchProblems",searchProblems)
		g.GET("getSubmission",getSubmission)
		g.POST("searchContests",searchContests)
		g.GET("enterContest",enterContest)
		g.POST("checkContestPwd",checkContestPwd)
		g.GET("getContestContent",getContestContent)
		g.GET("getCproblems",getCproblems)
		g.GET("getOneContestProblem",getOneContestProblem)
		g.POST("submitContestCode",submitContestCode)
		g.POST("searchCsubmissions",searchCsubmissions)
		g.GET("getCsubmission",getCsubmission)
		g.GET("getBaseCS",getBaseCS)
		g.GET("getCproblemHistory",getCproblemHistory)
		g.GET("getLabelList",getLabelList)
		g.GET("getRankList",getRankList)
	}
}
