package job

import (
	"job/search/domain"
)

// Interface to Account service
type JobService interface {
	Gets(query string) ([]*domain.JobResponse, error)
	Create(job domain.JobContent) (*domain.JobResponse, error)
	Update(id int, job domain.JobContent) (*domain.JobResponse, error)
	Delete(id int) (*domain.JobResponse, error)
}

// jobService is wrapper for Database struct
type jobService struct {
	db       Database
	dbsearch DatabaseSearch
}

// NewJobService will create jobService instance
func NewJobService(db Database, dbsearch DatabaseSearch) *jobService {
	return &jobService{db: db, dbsearch: dbsearch}
}

// Create method will send create record request to datastore/ repository
func (s *jobService) Create(job domain.JobContent) (*domain.JobResponse, error) {
	// call Create from repository/
	u, err := s.db.Create(job)

	// if error occur, return nil rfor the response as well as return the error
	if err != nil {
		return nil, err
	}

	// call Create from repository/
	_, err = s.dbsearch.Create(*u)

	// if error occur, return nil rfor the response as well as return the error
	if err != nil {
		return nil, err
	}

	return u.ToJobResponse(), nil
}

// Gets method will get all job record from repository/ datastore
func (s *jobService) Gets(query string) ([]*domain.JobResponse, error) {

	// Call Gets from repository/ to retreive all Job record
	res, err := s.dbsearch.Search(query)

	if err != nil {
		return nil, err
	}

	if len(res) < 1 {
		return []*domain.JobResponse{}, err
	}

	// Call Gets from repository/ to retreive all Job record
	users, err := s.db.Gets(res)

	// if error occur, return nil for the response slice as well as return the error
	if err != nil {
		return nil, err
	}

	// if no error found, convert all 'Job' record to JobResponse dto
	var uRes []*domain.JobResponse
	for _, job := range users {
		uRes = append(uRes, job.ToJobResponse())
	}

	// return response slice and nil if no error found
	return uRes, nil
}

// Update will send update request to datastore/ repository
func (s *jobService) Update(id int, job domain.JobContent) (*domain.JobResponse, error) {
	// call Database Update method from repository/ to update certain record
	u, err := s.db.Update(id, job)

	// return nil and the error if error occur
	if err != nil {
		return nil, err
	}

	// call Database Search Update method from repository/ to update certain record
	_, err = s.dbsearch.Update(id, job)

	if err != nil {
		return nil, err
	}

	// return job response dto and nil for the error
	return u.ToJobResponse(), nil
}

// Delete method will send request to delete record to datastore/ repository
// based on job 'id'
func (s *jobService) Delete(id int) (*domain.JobResponse, error) {
	// call Delete method from repository/ datastore
	u, err := s.db.Delete(id)

	// check if error occur while executing Delete method
	if err != nil {
		return nil, err
	}

	_, err = s.dbsearch.Delete(id)

	// check if error occur while executing Delete method
	if err != nil {
		return nil, err
	}

	// return job response dto and nil if no error found
	return u.ToJobResponse(), nil
}
