package job

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine, controller jobController) {
	// prepare router
	// main group api endpoint url : /v1
	v1 := r.Group("/v1")

	// job app group api endpoint : /v1/job
	job := v1.Group("/job")
	job.POST("/", controller.JobCreateController)
	job.PUT("/:id", controller.JobUpdateController)
	job.DELETE("/:id", controller.JobDeleteController)
	job.GET("/", controller.JobGetsController)
}
