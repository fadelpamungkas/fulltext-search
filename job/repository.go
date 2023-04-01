package job

import (
	"context"
	"fmt"
	"job/search/domain"
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
	entry := map[string]interface{}{"id": job.ID, "title": job.Title, "description": job.Description}

	task, err := client.Index.AddDocuments(entry)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Create method will insert new record to database. 'C' part of the CRUD
func (pool Database) Create(job domain.JobContent) (*domain.Job, error) {
	// sql for inserting new record
	q := `INSERT INTO job (title,description)
          VALUES ($1,$2) RETURNING id,title,description`

	// execute query to insert new record. it takes 'job' variable as its input
	// the result will be placed in 'row' variable
	row := pool.DB.QueryRow(context.Background(), q,
		job.Title, job.Description)

	// create 'u' variable as 'Job' type to contain scanned data value from 'row' variable
	u := new(domain.Job)

	// scan 'row' variable and place the value to 'u' variable as well as check for error
	err := row.Scan(
		&u.ID,
		&u.Title,
		&u.Description,
	)

	// return nil and error if scan operation is fail/ error found
	if err != nil {
		return nil, err
	}

	// return 'u' and nil if no error found
	return u, nil
}

func (client DatabaseSearch) Search(query string) ([]domain.JobId, error) {
	res, err := client.Index.Search(query,
		&meilisearch.SearchRequest{
			Limit:                10,
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
		id := getIntField(hit, "id")
		jobResults = append(jobResults, domain.JobId{ID: id})
	}

	return jobResults, nil
}

func getIntField(hit map[string]interface{}, field string) int {
	val, ok := hit[field]
	if !ok {
		return 0
	}
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case int64:
		return int(v)
	default:
		panic(fmt.Sprintf("unexpected type %T for %s field", v, field))
	}
}

// Gets method will get all job data. extended 'R' part of the CRUD
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
			// create 'u' for struct 'Job'
			u := new(domain.Job)

			// scan rows and place it in 'u' (job) container
			err := rows.Scan(
				&u.ID,
				&u.Title,
				&u.Description,
			)

			// return nil and error if scan operation fail
			if err != nil {
				return nil, err
			}

			// add u to job slice
			job = append(job, u)
		}
	}

	// return job slice and nil for the error
	return job, nil
}

func (client DatabaseSearch) Update(id int, job domain.JobContent) (*meilisearch.TaskInfo, error) {
	documents := []map[string]interface{}{
		{
			"id":          id,
			"title":       job.Title,
			"description": job.Description,
		},
	}
	return client.Index.UpdateDocuments(documents)
}

// Update will update job record based on their id
func (pool Database) Update(id int, job domain.JobContent) (*domain.Job, error) {
	// prepare update query
	q := `UPDATE job SET 
            title = $2,
            description  = $3
          WHERE id = $1
          RETURNING id, title, description;
         `
	// execute update query
	row := pool.DB.QueryRow(context.Background(), q, id,
		job.Title, job.Description)

	// create container variable for Job
	u := new(domain.Job)

	// scan data and place it on 'u' variable we create before and check for error
	if err := row.Scan(
		&u.ID,
		&u.Title,
		&u.Description,
	); err != nil {
		return nil, err
	}

	// return variable 'u' as Job and nil/ no error
	return u, nil
}

func (client DatabaseSearch) Delete(id int) (*meilisearch.TaskInfo, error) {
	return client.Index.DeleteDocuments([]string{
		strconv.Itoa(id),
	})
}

// Delete method will delete job record based on its 'id'
func (pool Database) Delete(id int) (*domain.Job, error) {
	// query for deleting job data
	q := `DELETE FROM job WHERE id = $1 RETURNING id,title,description;`

	// execute query
	row := pool.DB.QueryRow(context.Background(), q, id)

	// create container variable for Job
	u := new(domain.Job)

	// if error occur, return the error
	if err := row.Scan(
		&u.ID,
		&u.Title,
		&u.Description,
	); err != nil {
		return nil, err
	}

	// return nil if no error found
	return u, nil

}
