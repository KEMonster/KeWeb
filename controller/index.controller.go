package controller

import(
	//"net/http"
	"fmt"
	"text/template"
	"../base"
)

type IndexController struct{
	base.Controller
}

func (this *IndexController)Index(){
	t, err := template.ParseFiles("./tpl/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	if err := t.Execute(this.W,nil); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
	return
}