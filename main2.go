package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type User struct {
	Id       string `json:"Id"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

func test(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Main website",
		"content": "content is here123",
	})
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var post User
	json.Unmarshal(reqBody, &post)
	json.NewEncoder(w).Encode(post)

	newData, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/post", createNewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
