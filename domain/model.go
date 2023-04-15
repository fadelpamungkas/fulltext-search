package domain

// Job is job model reflect the 'jobs' database table
type Job struct {
	ID            int    `json:"id" validate:"required"`
	Title_greek   string `json:"title_greek" validate:"required"`
	Title_english string `json:"title_english,omitempty"`
	Keywords      string `json:"keywords,omitempty"`
}

type JobId struct {
	ID int `json:"id" validate:"required"`
}

type JobContent struct {
	Title_greek   string `json:"title_greek" validate:"required"`
	Title_english string `json:"title_english,omitempty"`
	Keywords      string `json:"keywords,omitempty"`
}

// JobResponse is to response the client/request with 'job' data
type JobResponse struct {
	ID            int    `json:"id" validate:"required"`
	Title_greek   string `json:"title_greek" validate:"required"`
	Title_english string `json:"title_english,omitempty"`
	Keywords      string `json:"keywords,omitempty"`
}

// type JobContentWithValidation struct {
// 	JobContent
// 	ValidationTag string `validate:"has_required_content"`
// }

// convert 'Job' model to 'JobResponse' DTO
func (j Job) ToJobResponse() *JobResponse {
	return &JobResponse{
		ID:            j.ID,
		Title_english: j.Title_english,
		Title_greek:   j.Title_greek,
		Keywords:      j.Keywords,
	}
}
