package adminApp

import "github.com/gin-gonic/gin"

func urls_init(r *gin.Engine)  {
	authorized := r.Group("/admin")
	authorized.Use(Authorrize)
	{
		authorized.POST("problemHandle",problemHandle)
		authorized.GET("delFile",delFile)
		authorized.GET("getProblems",getProblems)
		authorized.GET("getOneProblem",getOneProblem)
		authorized.GET("deleteProblem",deleteProblem)
		authorized.POST("addTestdatas",addTestdatas)
		authorized.GET("removeTestdatas",removeTestdatas)
		authorized.GET("getTestdatas",getTestdatas)
		authorized.GET("getOneTestdata",getOneTestdata)
		authorized.GET("removeOneTestdata",removeOneTestdata)
		authorized.GET("searchProbelmTitle",searchProbelmTitle)
		authorized.POST("contestHandle",contestHandle)
		authorized.POST("searchContests",searchContests)
		authorized.GET("getContest",getContest)
		authorized.POST("addCproblems",addCproblems)
		authorized.GET("getCproblems",getCproblems)
		authorized.GET("deleteContest",deleteContest)
		authorized.GET("getContestCode",getContestCode)
		authorized.POST("searchCsubmissions",searchCsubmissions)
		authorized.GET("rejudge",rejudge)

	}
	g := r.Group("/admin")
	{
		g.POST("login",login)
		g.GET("logout",logout)
		g.GET("autologin",autologin)
	}
}