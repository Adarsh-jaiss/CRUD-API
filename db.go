	package main

	import(
		"database/sql"
		_ "github.com/lib/pq"
	)

	var DB *sql.DB

	// TODO - Fix the DB connection and also test the application
	func OpenDB()error{
		var err error
		DB,err = sql.Open("postgres","user=postgres dbname=users password='Lets do it' sslmode=disable")
		if err!= nil{
			return err
		}
		return nil
	}

	func CloseDB() error {
		return DB.Close()
	}


