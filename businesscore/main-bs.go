package businesscore

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	//"github.com/go-chi/jwtauth"
	dal "github.com/yockyyuwono/golang2/dal"

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
//var tokenAuth *jwtauth.JWTAuth

/*
var tokenAuth *jwtauth.JWTAuth


type tknStruct struct {
	Token string //Harus diawali huruf besar
}

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
		w.Write([]byte(`invalid username/password`))
		return
	}

	tokenAuth = jwtauth.New("HS256", []byte("Rahasia1780"), nil)
	//jwtauth.SetExpiry(claim )
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": username + password})
	tknJson := &tknStruct{string(tokenString)}

	var result, err1 = json.Marshal(tknJson)
	fmt.Println(string(result))
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)

}
*/
type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type tknStruct struct {
	Token string //Harus diawali huruf besar
}

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
		w.Write([]byte(`invalid username/password`))
		return
	}

	//expiredTime := time.Now().AddDate(0, 0, 7*1)
	expiredTime := time.Now().UTC().AddDate(0, 0, 7*1).Unix() // 7 hari
	//expiredTime := time.Now().UTC().Add(60 * time.Second).Unix()	// 60 Detik

	claims := customClaims{
		Username: username + password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
			Issuer:    "mygolang",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("Rahasia1780"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//var result, err1 = json.Marshal(signedToken)
	tknJson := &tknStruct{string(signedToken)}
	var result, err1 = json.Marshal(tknJson)
	//fmt.Println(string(result))
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
}

/*
func CekHeaderToken(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	resultstring := GreetingFunction(reqToken)
	var result, err1 = json.Marshal(resultstring)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
	//return
}

func VerifyToken1(r *http.Request) bool {

	reqToken := r.Header.Get("Authorization")
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("Rahasia1780"), nil
	})
	if err == nil && token.Valid {
		fmt.Println("valid token")
		return true
	} else {
		fmt.Println("invalid token")
		return false
	}

}
func VerifyToken2(w http.ResponseWriter, r *http.Request) bool {
	// Token from another example.  This token is expired
	reqToken := r.Header.Get("Authorization")
	//var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			return false
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return false
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return false
	}

	// Output: Timing is everything
}
*/
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
