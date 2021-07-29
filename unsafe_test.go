package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "zhangsan"},
	{ID: 2, Name: "zhsan"},
	{ID: 3, Name: "zh"},
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\":\""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "not found")
	}

}

func TestUnsafe(t *testing.T) {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
}
