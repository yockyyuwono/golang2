package helper

import (
	"encoding/json"
	"net/http"

	mdl "github.com/yockyyuwono/golang2/model"
)

//A handler to fetch all the jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	//make a slice to hold our jobs data
	var jobs []mdl.Job
	///add some job to the slice
	jobs = append(jobs, mdl.Job{ID: 1, Name: "Accounting"})
	jobs = append(jobs, mdl.Job{ID: 2, Name: "Programming"})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
