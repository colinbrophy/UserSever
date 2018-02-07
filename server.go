package main

import ("net/http"
        "fmt"
	    "log"
	"encoding/json"
	"path"
    "strconv"
	"io/ioutil"
)

func main() {
	initDatabase()

	http.HandleFunc("/allusers/", users)
	http.HandleFunc("/users/", user)
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func users (w http.ResponseWriter, r *http.Request) {
		usrs := Users()
		usrJson, _ := json.Marshal(&usrs)
		fmt.Fprintf(w, "%s\n", string(usrJson))
}

type jsonUser struct {
	FirstName string
	LastName string
}

func user(w http.ResponseWriter, r *http.Request) {
	idStr := path.Base(r.URL.Path)
	id, _ := strconv.Atoi(idStr)

	switch (r.Method) {
	case "GET":
		usr := UserFromDb(id)
		usrJson, _ := json.Marshal(&usr)
		fmt.Fprintf(w, "%s\n", string(usrJson))
	case "DELETE":
		DeleteUsr(id)
	case "PUT":
		var updateUsr jsonUser;
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body ,&updateUsr)
		UpdateUsr(id, updateUsr.FirstName, updateUsr.LastName)
	case "POST":
		var newUsr jsonUser;
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body ,&newUsr)
		CreateUsr(newUsr.FirstName, newUsr.LastName)
	}

}
