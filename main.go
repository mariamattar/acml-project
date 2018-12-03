package main
import (	//"encoding/json"
 //"github.com/ryanbradynd05/go-tmdb"
"fmt"
"io/ioutil"
"log"
"net/http"
"os"
"database/sql"
_ "github.com/go-sql-driver/mysql"



)


type TMDb struct {
	apiKey string
  }
func main(){
	response, err := http.Get("https://api.themoviedb.org/3/movie/550?api_key=440ffd60b64ed0b7ef1fa3e8e3c664a7")
	
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	db, err1 := sql.Open("mysql",
	"root:secret@tcp(localhost:5432)/myapp")
	if err1 != nil {
		panic(err1)
	}
	db.SetMaxOpenConns(5)
	defer db.Close()
 
	_,err1 = db.Exec("CREATE DATABASE myapp")
	if err1 != nil {
		panic(err1)
	}
 
	_,err1 = db.Exec("USE myapp" )
	if err1 != nil {
		panic(err1)
	}
 
stmt, err := db.Prepare("CREATE Table wishlist(id int NOT NULL AUTO_INCREMENT, movie1 varchar(500), PRIMARY KEY (id));")
 if err != nil {
 fmt.Println(err.Error())
 }
 
 _, err = stmt.Exec()
 if err != nil {
 fmt.Println(err.Error())
 } else {
 fmt.Println("Table created successfully..")
 }

 stmtIns, err := db.Prepare("INSERT INTO wishlist VALUES( 1,responseData )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when w


}