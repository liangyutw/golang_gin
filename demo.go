package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type demoData struct {
	Data map[string]string
}

type fiboData struct {
	Fibo map[int]int
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("template/demo.html"))
}

func demo(w http.ResponseWriter, r *http.Request) {

	// type 1 取值方法
	// id := r.URL.Query()
	// fmt.Println("data =>", id)

	// type 2 取值方法
	vars := mux.Vars(r)
	data := demoData{Data: vars}
	fmt.Println("data =>", data)
	tmpl.Execute(w, data)
}

func fibonacci(s int) int {
	if s < 2 {
		return s
	}
	return fibonacci(s-1) + fibonacci(s-2)
}

func loopFinbo(w http.ResponseWriter, r *http.Request) {

	//定義map
	myMap := make(map[int]int)

	//接參數
	vars := mux.Vars(r)

	//檢查參數數量
	if (len(vars)) == 1 {

		//判斷是否有值
		value, isExist := vars["num"]
		if isExist != true {
			log.Fatal(isExist)
		}

		//轉成數字
		intVar, _ := strconv.Atoi(value)

		//執行費氏數列
		for i := 1; i < intVar; i++ {
			myMap[i] = fibonacci(i)
		}
	}

	//塞入結構map
	data := fiboData{Fibo: myMap}
	fmt.Println("data =>", data)
	//與網頁綁定執行
	tmpl.Execute(w, data)
}

func main() {

	// type 1 取 query string
	// ex: http://localhost:9999/?aa=123
	// http.HandleFunc("/", demo)
	// http.ListenAndServe(":9999", nil)

	// type 2 使用 mux 取得路由參數
	r := mux.NewRouter()
	// r.HandleFunc("/{xx}/{dd}", demo)
	r.HandleFunc("/fibonacci/{num}", loopFinbo)
	http.ListenAndServe(":9999", r)
}
