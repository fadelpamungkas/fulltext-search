package domain

// Job is job model reflect the 'jobs' database table
type Job struct {
	ID                  int    `json:"id" validate:"required"`
	Title_english       string `json:"title_english"`
	Description_english string `json:"description_english"`
	Title_greek         string `json:"title_greek"`
	Description_greek   string `json:"description_greek"`
}

type JobId struct {
	ID int `json:"id" validate:"required"`
}

type JobContent struct {
	Title_english       string `json:"title_english"`
	Description_english string `json:"description_english"`
	Title_greek         string `json:"title_greek"`
	Description_greek   string `json:"description_greek"`
}

// JobResponse is to response the client/request with 'job' data
type JobResponse struct {
	ID                  int    `json:"id"`
	Title_english       string `json:"title_english"`
	Description_english string `json:"description_english"`
	Title_greek         string `json:"title_greek"`
	Description_greek   string `json:"description_greek"`
}

type JobContentWithValidation struct {
	JobContent
	ValidationTag string `validate:"has_required_content"`
}

// convert 'Job' model to 'JobResponse' DTO
func (j Job) ToJobResponse() *JobResponse {
	return &JobResponse{
		ID:                  j.ID,
		Title_english:       j.Title_english,
		Title_greek:         j.Title_greek,
		Description_english: j.Description_english,
		Description_greek:   j.Description_greek,
	}
}
