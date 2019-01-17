package model

import (
	"../base"
	"database/sql"
	"fmt"
)

// CREATE TABLE `tb_order`(
//     `sOrderCode` int(10) NOT NULL AUTO_INCREMENT,
//     `sMethod` varchar(20) NULL DEFAULT '',
//     `sZipCode` varchar(10) NULL DEFAULT '',
//     `sSendName` varchar(50) NULL DEFAULT '',
//     `sSendAddress` varchar(128) NULL DEFAULT '',
//     `sRecipientName` varchar(50) NULL DEFAULT '',
//     `sRecipientAddress` varchar(128) NULL DEFAULT '',
//     `sUserCode` varchar(20) NOT NULL DEFAULT '',
//     PRIMARY KEY(`sOrderCode`),
//     KEY `idx_usercode` (`sUserCode`),
//     KEY `idx_method` (`sMethod`) 
// )ENGINE=Innodb DEFAULT charset=UTF8;

type ModelOrder struct{
	OrderCode string
	Method string
	ZipCode string
	SendName string
	SendAddress string
	RecipientName string
	RecipientAddress string
	UserCode string
}

var DB *sql.DB

func init(){
	DB = base.DbCon
}

func (m *ModelOrder)CreateOrder() int64 {
	stmt, err :=DB.Prepare("Insert tb_order set sMethod=?,sZipCode=?,sSendName=?,sSendAddress=?,sRecipientName=?,sRecipientAddress=?,sUserCode=?")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	res, err := stmt.Exec( m.Method,m.ZipCode,m.SendName,m.SendAddress,m.RecipientName,m.RecipientAddress,m.UserCode)
    if err != nil {
		fmt.Println(err)
		return 0
	}
    id, err := res.LastInsertId()
    if err != nil {
		fmt.Println(err)
		return 0
	}
    return id
}

func (m *ModelOrder)CheckParam() string {
	if m.Method == ""{
		return "Method must not empty"
	}else if m.ZipCode == ""{
		return "ZipCode must not empty"
	}else if m.SendName == ""{
		return "SendName must not empty"
	}else if m.SendAddress == ""{
		return "SendAddress must not empty"
	}else if m.RecipientName == ""{
		return "RecipientName must not empty"
	}else if m.RecipientAddress == ""{
		return "RecipientAddress must not empty"
	}else if m.UserCode == ""{
		return "UserCode must not empty"
	}
	return ""
}