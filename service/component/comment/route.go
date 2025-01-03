package comment

import "github.com/gin-gonic/gin"

//import "github.com/alaa2amz/g1/mw"

func Register(r *gin.Engine) {
	cg := r.Group("") //common group
	//	cg.Use(mw.Logged)

	wg := cg.Group(Path) //web group
	{
		wg.GET("/new", nw)         //new
		wg.POST("", cr)            //create
		wg.GET("", rt)             //retrieve
		wg.GET("/list", rt)        //retrieve
		wg.GET("/:id", gt)         //get one
		wg.GET("/:id/edit", ed)    //edit
		wg.POST("/:id/update", up) //update
		wg.POST("/:id/delete", dl) //delete

		wg.POST("/:id/:path", crAs)
	}

	ag := cg.Group("/api" + Path) //api group
	{
		ag.POST("", cr)
		ag.GET("", rt)
		ag.GET("/:id", gt)
		ag.PUT("/:id", up)
		ag.DELETE("/:id", dl)

		ag.POST("/a/:assoc/*actions", cr)   //create
		ag.POST("/aa/:assoc/:aid", cr)      //create
		ag.POST("/aaa/:id/:assoc/:aid", cr) //create
	}

}
