package base

import(
	"net/http"
)

type Controller struct{
	W http.ResponseWriter
	R *http.Request
}

func (this *Controller) GetParam(name string) (string) {
	r := this.R
	r.ParseForm()
	keys, ok := r.Form[name]
	if !ok || len(keys) < 1 {
		return ""
	}
	return r.Form[name][0]
}


// func (m *ModelOrder)CheckParam() string {
// 	if m.Method == ""{
// 		return "Method must not empty"
// 	}else if m.ZipCode == ""{
// 		return "ZipCode must not empty"
// 	}else if m.SendName == ""{
// 		return "SendName must not empty"
// 	}else if m.SendAddress == ""{
// 		return "SendAddress must not empty"
// 	}else if m.RecipientName == ""{
// 		return "RecipientName must not empty"
// 	}else if m.RecipientAddress == ""{
// 		return "RecipientAddress must not empty"
// 	}else if m.UserCode == ""{
// 		return "UserCode must not empty"
// 	}
// 	return ""
// }