package main

import (
	"VIVASOFT1/src/apis/user_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	db, err := sql.Open("mysql","root:1703179Faysal@tcp(localhost:3306)/user_activity")
// 	if err !=nil {
// 		fmt.Println("Error")
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("Error")
// 		panic(err.Error())
// 	}

// 	fmt.Println("Connected")
// }

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/findall", user_api.FindAll).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println((err))
	}
}
