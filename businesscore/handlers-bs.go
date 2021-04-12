package businesscore

import (
	"encoding/json"
	"net/http"

	//bc "github.com/yockyyuwono/golang2/businesscore"
	mdl "github.com/yockyyuwono/golang2/model"
)

//A handler to fetch all the jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	//make a slice to hold our jobs data
	var jobs []mdl.Job
	///add some job to the slice
	jobs = append(jobs, mdl.Job{ID: 1, Name: "hellogreetingmain"})
	jobs = append(jobs, mdl.Job{ID: 2, Name: "Programming"})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func GetGreetingFunction(w http.ResponseWriter, r *http.Request) {
	//resultstring := bc.GreetingFunction("saya")
	resultstring := GreetingFunction("kamu")
	var result, err1 = json.Marshal(resultstring)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
	//return
}
