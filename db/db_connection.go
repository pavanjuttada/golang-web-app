package db;

// var db *sql.DB
func createDBConnection() *sql.DB {
    con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demodb") //localhost:3306
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    return con;
}