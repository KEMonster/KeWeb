package main

import (
	"net/http"
	"os"
	"strings"
	"./router"
)

type MyMux struct {
}


// func init(){
// 	http.Handle("/static/", http.StripPrefix("./static/", http.FileServer(http.Dir("static")))) 
// }

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path,"/static"){
		file := "."+r.URL.Path
		f,err := os.Open(file)
		defer f.Close()
		 if err != nil && os.IsNotExist(err){
			http.NotFound(w, r)
			return 
		 }
		http.ServeFile(w,r,file)
	}else{
		router.DoAction(r.URL.Path,w,r)
	}
    
    return
}


func main() {
    mux := &MyMux{}
    http.ListenAndServe(":9503", mux)
}