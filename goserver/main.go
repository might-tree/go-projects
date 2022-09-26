package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name:=r.FormValue("name")
	address=r.FormValue("address")
	fmt.Fprintln(w,"Name=",name)
	fmt.Fprintln(w,"Address=",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello!")	//Fprintf formates according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
}

func main(){
	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Starting server at port 8080")

	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}