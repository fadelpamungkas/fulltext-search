package job

import (
	"fmt"
	"job/search/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// jobController is type wrapper for JobService interface
type jobController struct {
	Service JobService
}

// NewJobController is new instance for jobController
func NewJobController(svc JobService) jobController {
	return jobController{Service: svc}
}

// Check if the content has both title and description
func hasRequiredContent(fl validator.FieldLevel) bool {
	content := fl.Parent().Interface().(domain.JobContentWithValidation).JobContent
	hasTitle := content.Title_english != "" && content.Title_greek != ""
	hasDescription := content.Description_english != "" && content.Description_greek != ""

	return hasTitle || hasDescription
}

// JobCreate method will process request to insert new 'Job' data and
// response with the created data back to the job (if no error found)
func (h *jobController) JobCreateController(c *gin.Context) {

	var validate = validator.New()
	_ = validate.RegisterValidation("has_required_content", func(fl validator.FieldLevel) bool {
		return hasRequiredContent(fl)
	})

	// get request data from context that containing 'Job' model information
	// and bind it to a variable matching the requested data
	var u domain.JobContentWithValidation

	// if request data binding error than return 400/ bad request
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		// exit process
		return
	}

	//use validator library to validate required fields
	if err := validate.Struct(&u); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		return
	}

	// send data to service layer to further process (create record)
	job, err := h.Service.Create(u.JobContent)

	// if error occur while trying to save the data, return 500/ internal server error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)

		// exit process
		return
	}

	//  if no error found, send 200/ status ok as well as the 'JobResponse' data
	c.JSON(
		http.StatusOK,
		job,
	)
}

// JobGetsController is method to process request to get all job data
func (h *jobController) JobGetsController(c *gin.Context) {
	q := c.Query("query")
	jobs, err := h.Service.Gets(q)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)
		return
	}

	// no error occur then send status ok and jobs data
	c.JSON(
		http.StatusOK,
		jobs,
	)
}

// JobUpdateController method will process request to update 'Job' data and
// response with the updated data
func (h *jobController) JobUpdateController(c *gin.Context) {
	// get request parameter for 'id'
	id := c.Param("id")
	uid, err := strconv.Atoi(id)

	// if error found, respnse with 400(bad request) and exit the process
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		// exit process
		return
	}

	var validate = validator.New()
	_ = validate.RegisterValidation("has_required_content", func(fl validator.FieldLevel) bool {
		return hasRequiredContent(fl)
	})

	// get request data from context that containing 'Job' model information
	// and bind it to a variable matching the requested data
	var u domain.JobContentWithValidation

	// if request data binding error than return 400/ bad request
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		// exit process
		return
	}

	//use validator library to validate required fields
	if err := validate.Struct(&u); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		return
	}

	// send data to service layer to further process (update record)
	job, err := h.Service.Update(uid, u.JobContent)

	// if error occur while trying to save the data, return 500/ internal server error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)

		// exit process
		return
	}

	//  if no error found, send 200/ status ok as well as the 'JobResponse' data
	c.JSON(
		http.StatusOK,
		job,
	)
}

// JobDeleteController method will process request to delete 'Job' data and
// response with the updated data
func (h *jobController) JobDeleteController(c *gin.Context) {
	// get request parameter for 'id'
	id := c.Param("id")
	uid, err := strconv.Atoi(id)

	// if error found, respnse with 400(bad request) and exit the process
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)

		// exit process
		return
	}

	// send data to service layer to further process (delete record)
	job, err := h.Service.Delete(uid)

	// if error occur while trying to save the data, return 500/ internal server error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": fmt.Sprintf("internal server error: %v\n", err),
			},
		)

		// exit process
		return
	}

	//  if no error found, send 200/ status ok as well as the 'JobResponse' data
	c.JSON(
		http.StatusOK,
		job,
	)
}
