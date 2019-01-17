package controller

import(
	"fmt"
	"encoding/json"
	"../base"
	"../model"
)

type OrderController struct{
	base.Controller
}

type Result struct{
	Result int `json:"result"`
	OrderCode int64 `json:"ordercode"`
	Status string `json:"status"`
}


func (this *OrderController) AddOrder(){
	

	var model_order model.ModelOrder
	model_order.Method = this.GetParam("method")
	model_order.ZipCode = this.GetParam("zipcode") 
	model_order.SendName = this.GetParam("sendname") 
	model_order.SendAddress = this.GetParam("sendaddress") 
	model_order.RecipientName = this.GetParam("recipientname") 
	model_order.RecipientAddress = this.GetParam("recipientaddress")
	model_order.UserCode = this.GetParam("usercode")

	sMsg := model_order.CheckParam()
	
	//http://test.keweisen.com:9503/order/addorder?method=UPSs&zipcode=sa1222&sendname=老师&sendaddress=学校&recipientname=学生&recipientaddress=家里&usercode=A11
	var id int64
	result := &Result{}
	if sMsg == ""{
		id = model_order.CreateOrder()
	}
	
	if id > 0 {
		result.Result = 0
		result.OrderCode = id
		result.Status = "success"
	}else{
		result.Result = -1
		result.Status = sMsg
	}
	jsonStr, err := json.Marshal(result)
    if err != nil {
        fmt.Println("error:", err)
    }
	fmt.Fprintf(this.W, string(jsonStr))

}