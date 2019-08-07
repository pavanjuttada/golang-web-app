package models

import (
	"fmt"
)

type Page struct {
	Title string
	Body  []byte
}
type Pagedata struct {
	// Title string
	Datasql []Entry
    Userdata []Entry
}

type Entry struct {
    Number int
    ID, Name, Email, Address string
}

func EditUser(user_id string) (p *Pagedata) {
    
     //SELECT DATA BY ID
    users, err := db.Query("SELECT * FROM users WHERE id="+user_id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer users.Close()

    results := []Entry{}
    setData := Entry{}

    for users.Next() {
        var id, name, email, address string
        users.Scan(&id, &name, &email, &address)
        setData.ID = id
        setData.Name = name
        setData.Email = email
        setData.Address = address
        results = append(results, setData)
    }

    fmt.Println(results)

    //SELECT ALL DATA
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    // Fetch rows

    resultsRow := []Entry{}
    setDataRow := Entry{}
    i := 0
    for rows.Next() {
        var id, name, email, address string
        rows.Scan(&id, &name, &email, &address)
        i += 1
        setDataRow.Number = i
        setDataRow.ID = id
        setDataRow.Name = name
        setDataRow.Email = email
        setDataRow.Address = address
        resultsRow = append(resultsRow, setDataRow)
    }

    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    return &Pagedata{Datasql: resultsRow, Userdata: results}
}

func GetAllUsers() (p *Pagedata){
	initDB()
	rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // text := ""
    resultsRow := []Entry{}
    setDataRow := Entry{}
    i := 0
    for rows.Next() {
        var id, name, email, address string
        rows.Scan(&id, &name, &email, &address)
        i += 1
        setDataRow.Number = i
        setDataRow.ID = id
        setDataRow.Name = name
        setDataRow.Email = email
        setDataRow.Address = address
        resultsRow = append(resultsRow, setDataRow)
    }

    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	return &Pagedata{Datasql: resultsRow}
}
func UpdateUser(id, name, email, address string) bool{
        // Execute the query
    stmtIns, err := db.Query("UPDATE users SET name='"+name+"', email='"+email+"', address='"+address+"' WHERE id="+id)
    
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	defer stmtIns.Close()
	
	return true
}
func CreateUser(name, email, address string) bool {
        // Execute the query
    stmtIns, err := db.Query("INSERT INTO users (name, email, address) VALUES ('"+name+"','"+email+"','"+address+"')")
    
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close()
	return true
}
func DeleteUser(user_id string) bool{
	
    fmt.Println(user_id);
        // Execute the query
    stmtIns, err := db.Query("DELETE FROM users WHERE id="+user_id)
    
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close()
	return true
}

