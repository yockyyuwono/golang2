package businesscore

import (
	"encoding/json"
	"net/http"

	dal "github.com/yockyyuwono/golang2/dal"

	"github.com/go-chi/jwtauth"

	mdl "github.com/yockyyuwono/golang2/model"
)

/*
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
*/
var tokenAuth *jwtauth.JWTAuth

/*
const USERNAME = "1"
const PASSWORD = "1"
*/

func GetToken(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
	}
	sqlQuery := "SELECT COUNT(*) "
	sqlQuery += "FROM Auth With(NoLock) "
	sqlQuery += "WHERE UserName = '" + username + "' "
	sqlQuery += "AND Passwords = '" + password + "' "
	sqlQuery += "AND Active = '1' "

	isValid, errBool := dal.GetUserAuth_dal(sqlQuery)
	if errBool != nil {
		http.Error(w, errBool.Error(), http.StatusInternalServerError)
		return
	}
	//isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return
	}

	tokenAuth = jwtauth.New("HS256", []byte("Rahasia1780"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": username + password})

	var result, err1 = json.Marshal(tokenString)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
	/*
		tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

		// For debugging/example purposes, we generate and print
		// a sample jwt token with claims `user_id:123` here:
		_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
		//fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

		resultstring := tokenString
		var result, err1 = json.Marshal(resultstring)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		//return
	*/
}

func GetGreetingFunction(w http.ResponseWriter, r *http.Request) {
	//resultstring := bc.GreetingFunction("saya")
	//_, claims, _ := jwtauth.FromContext(r.Context())
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
