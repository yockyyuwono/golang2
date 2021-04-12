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

func GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		var resultstring, err1 = GetUserList_bc()
		//fmt.Println(resultstring)
		var result, err2 = json.Marshal(resultstring)

		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

		return
	}
}
func GetUserByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var userModel mdl.User_mdl

		err1 := json.NewDecoder(r.Body).Decode(&userModel)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		var resultstring, err2 = GetUserByCode_bc(userModel.UserCode)
		//fmt.Println(resultstring)
		var result, err3 = json.Marshal(resultstring)

		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

		return
	}
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var userSaveModel mdl.UserSave_mdl

		err1 := json.NewDecoder(r.Body).Decode(&userSaveModel)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		var resultstring, err2 = SaveUser_bc(userSaveModel)
		//fmt.Println(resultstring)
		var result, err3 = json.Marshal(resultstring)

		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

		return
	}
}

func SaveUserBulk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var userSaveModel []mdl.UserSave_mdl

		err1 := json.NewDecoder(r.Body).Decode(&userSaveModel)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		var resultstring, err2 = SaveUserBulk_bc(userSaveModel)
		//fmt.Println(resultstring)
		var result, err3 = json.Marshal(resultstring)

		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

		return
	}
}
