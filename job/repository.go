package job

import (
	"context"
	"job/search/domain"
	"job/search/utils"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
)

// Database
type Database struct {
	DB *pgxpool.Pool
}

// Database
type DatabaseSearch struct {
	Index *meilisearch.Index
}

// NewSelector is an initializer for Selector
func NewDatabase(pgx *pgxpool.Pool) Database {
	return Database{DB: pgx}
}

func NewDatabaseSearch(index *meilisearch.Index) DatabaseSearch {
	return DatabaseSearch{Index: index}
}

func (client DatabaseSearch) Create(job domain.Job) (*meilisearch.TaskInfo, error) {
	entry := map[string]interface{}{
		"id":    job.ID,
		"title": job.Title,
	}

	task, err := client.Index.AddDocuments(entry)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Create method will insert new record to database.
func (pool Database) Create(job domain.JobContent) (*domain.Job, error) {
	// sql for inserting new record
	q := `
  INSERT INTO job (title)
  VALUES ($1) 
  RETURNING id, title;
  `

	// execute query to insert new record. it takes 'job' variable as its input
	// the result will be placed in 'row' variable
	row := pool.DB.QueryRow(
		context.Background(),
		q,
		job.Title,
	)

	// create 'j' variable as 'Job' type to contain scanned data value from 'row' variable
	j := new(domain.Job)

	// scan 'row' variable and place the value to 'j' variable as well as check for error
	err := row.Scan(
		&j.ID,
		&j.Title,
	)

	// return nil and error if scan operation is fail/ error found
	if err != nil {
		return nil, err
	}

	// return 'j' and nil if no error found
	return j, nil
}

func (client DatabaseSearch) Search(query string) ([]domain.JobId, error) {
	res, err := client.Index.Search(query,
		&meilisearch.SearchRequest{
			// Limit:                10,
			AttributesToRetrieve: []string{"id"},
		})
	if err != nil {
		return nil, err
	}

	// Parse the search results
	var hits []map[string]interface{}
	for _, hit := range res.Hits {
		hits = append(hits, hit.(map[string]interface{}))
	}

	// Print the search results
	var jobResults []domain.JobId
	for _, hit := range hits {
		// get the item id
		id := utils.GetIntField(hit, "id")
		jobResults = append(jobResults, domain.JobId{ID: id})
	}

	return jobResults, nil
}

// Gets method will get all job data.
func (pool Database) Gets(results []domain.JobId) ([]*domain.Job, error) {
	// Construct the SQL query
	query := "SELECT * FROM job WHERE id IN ("
	placeholders := make([]string, len(results))
	for i, r := range results {
		placeholders[i] = strconv.Itoa(r.ID)
	}
	query += strings.Join(placeholders, ",") + ")"

	// execute query
	rows, err := pool.DB.Query(context.Background(), query)

	// check if any error occur while executing the query
	if err != nil {
		return nil, err
	}

	// close rows if error ocur
	defer rows.Close()

	// iterate Rows
	var job []*domain.Job
	if rows != nil {
		for rows.Next() {
			// create 'j' for struct 'Job'
			j := new(domain.Job)

			// scan rows and place it in 'j' (job) container
			err := rows.Scan(
				&j.ID,
				&j.Title,
			)

			// return nil and error if scan operation fail
			if err != nil {
				return nil, err
			}

			// add j to job slice
			job = append(job, j)
		}
	}

	// return job slice and nil for the error
	return job, nil
}

func (client DatabaseSearch) Update(id int, job domain.JobContent) (*meilisearch.TaskInfo, error) {
	documents := []map[string]interface{}{
		{
			"id":    id,
			"title": job.Title,
		},
	}
	return client.Index.UpdateDocuments(documents)
}

// Update will update job record based on their id
func (pool Database) Update(id int, job domain.JobContent) (*domain.Job, error) {
	// prepare update query
	q := `
  UPDATE job SET 
  title= $2,
  WHERE id = $1
  RETURNING id, title;
  `
	// execute update query
	row := pool.DB.QueryRow(
		context.Background(),
		q,
		id,
		job.Title,
	)

	// create container variable for Job
	j := new(domain.Job)

	// scan data and place it on 'j' variable we create before and check for error
	if err := row.Scan(
		&j.ID,
		&j.Title,
	); err != nil {
		return nil, err
	}

	// return variable 'j' as Job and nil/ no error
	return j, nil
}

func (client DatabaseSearch) Delete(id int) (*meilisearch.TaskInfo, error) {
	return client.Index.DeleteDocuments([]string{
		strconv.Itoa(id),
	})
}

// Delete method will delete job record based on its 'id'
func (pool Database) Delete(id int) (*domain.Job, error) {
	// query for deleting job data
	q := `
  DELETE FROM job 
  WHERE id = $1 
  RETURNING id, title;
  `

	// execute query
	row := pool.DB.QueryRow(context.Background(), q, id)

	// create container variable for Job
	j := new(domain.Job)

	// if error occur, return the error
	if err := row.Scan(
		&j.ID,
		&j.Title,
	); err != nil {
		return nil, err
	}

	// return nil if no error found
	return j, nil
}
