package main

import (
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi"

	//bc "github.com/yockyyuwono/golang2/businesscore"
	hlp "github.com/yockyyuwono/golang2/helper"
)

/*
type Configuration struct {
	Db struct {
		Host     string
		User     string
		Password string
		Database string
	}
}
*/
//Config := hlp.Configuration{}

func main() {
	/*
		http.HandleFunc("/hellogreetingmain", hellogreetingmain)
		http.HandleFunc("/getmsusersingle", getmsusersingle)
		http.HandleFunc("/getmsuserlist", getmsuserlist)
		http.HandleFunc("/addmsusersingle", addmsusersingle)
		http.HandleFunc("/addmsuserbulk", addmsuserbulk)
	*/
	////http.ListenAndServe("192.168.205.20:8080", nil)

	myserver, _ := hlp.LoadConfiguration()
	//http.ListenAndServe(myserver.AppServer, nil)

	router := chi.NewRouter()
	router.Get("/api/jobs", hlp.GetJobs)
	router.Get("/api/GetGreetingFunction", hlp.GetGreetingFunction)
	//run it on port 8080
	//err := http.ListenAndServe("0.0.0.0:8080", router)
	err := http.ListenAndServe(myserver.AppServer, router)
	if err != nil {
		log.Fatal(err)
	}

}

/*
type Person struct {
	Name string
	Age  int
}
*/

/*
func hellogreetingmain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		//var personModel mdl.Person

		////var resultstring string

		////var resultstring string = GreetingFunction("Gladys")
		////var resultstring string = "saya"
		//resultstring := bc.GreetingFunction("saya")
		//var result, err1 = json.Marshal(resultstring)
		//fmt.Println(resultstring)
		//var result, err = json.Marshal(resultstring)

		//err := json.NewDecoder(r.Body).Decode(&personModel)

		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		////fmt.Fprintf(w, "Person: %+v", personModel.Name)
		////w.Write(p)

		resultstring := bc.GreetingFunction("saya")
		var result, err1 = json.Marshal(resultstring)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func getmsusersingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var userModel mdl.User

		err1 := json.NewDecoder(r.Body).Decode(&userModel)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		var resultstring, err2 = bc.GetMsUserSingle(userModel.UserCode)
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

func getmsuserlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		var resultstring, err1 = bc.GetMsUserList()
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

func addmsusersingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var userSaveModel mdl.UserSave

		err1 := json.NewDecoder(r.Body).Decode(&userSaveModel)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		var resultstring, err2 = bc.AddNewUserSingle(userSaveModel)
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

func addmsuserbulk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var userSaveModel []mdl.UserSave

		err1 := json.NewDecoder(r.Body).Decode(&userSaveModel)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		var resultstring, err2 = bc.AddNewUserBulk(userSaveModel)
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
*/
