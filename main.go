package main

import (
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	//"github.com/go-chi/jwtauth"

	//"github.com/go-chi/jwtauth/v5"

	bc "github.com/yockyyuwono/golang2/businesscore"
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
var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("Rahasia1780"), nil)
	/*
		// For debugging/example purposes, we generate and print
		// a sample jwt token with claims `user_id:123` here:
		_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
		fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	*/
}

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
	addr := myserver.AppServer
	fmt.Printf("Starting server on %v\n", myserver.AppServer)
	http.ListenAndServe(addr, router())

}
func router() http.Handler {
	router := chi.NewRouter()
	//// Protected routes
	router.Group(func(router chi.Router) {
		//// Seek, verify and validate JWT tokens
		router.Use(jwtauth.Verifier(tokenAuth))

		//// Handle valid / invalid tokens. In this example, we use
		//// the provided authenticator middleware, but you can write your
		//// own very easily, look at the Authenticator method in jwtauth.go
		//// and tweak it, its not scary.
		router.Use(jwtauth.Authenticator)

		////r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		////	_, claims, _ := jwtauth.FromContext(r.Context())
		////	w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		////})
		router.Get("/api/GetGreetingFunction", bc.GetGreetingFunction)
		router.Get("/api/GetUserList", bc.GetUserList)
		router.Get("/api/GetUserByCode", bc.GetUserByCode)
		router.Post("/api/SaveUser", bc.SaveUser)
		router.Post("/api/SaveUserBulk", bc.SaveUserBulk)
	})
	// Public routes
	//router.Get("/api/GetToken", bc.GetToken)
	router.Get("/api/GetToken", bc.GetToken)
	//router.Get("/api/CekHeaderToken", bc.VerifyToken1)
	//router.Get("/api/CekHeaderToken", bc.VerifyToken1)

	return router
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
