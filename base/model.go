package base

import(
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)


var DbCon *sql.DB


func init(){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_test_go?charset=utf8")
	CheckErr(err)
	DbCon = db
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
