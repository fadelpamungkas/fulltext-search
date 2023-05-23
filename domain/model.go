package domain

// Job is job model reflect the 'jobs' database table
type Job struct {
	ID    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type JobId struct {
	ID int `json:"id" validate:"required"`
}

type JobContent struct {
	Title string `json:"title" validate:"required"`
}

// JobResponse is to response the client/request with 'job' data
type JobResponse struct {
	ID    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

// convert 'Job' model to 'JobResponse' DTO
func (j Job) ToJobResponse() *JobResponse {
	return &JobResponse{
		ID:    j.ID,
		Title: j.Title,
	}
}
