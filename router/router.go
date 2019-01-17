package router
import(
	"../base"
	"../controller"
	"net/http"
)


func DoAction(path string,w http.ResponseWriter, r *http.Request){
	b := base.Controller{W:w,R:r}
	if path == "/" || path == "/index" {
		c := controller.IndexController{b}
		c.Index()
    }else if path == "/order/addorder"{
		c := controller.OrderController{b}
		c.AddOrder()
	}else{
		http.NotFound(w, r)
	}

}