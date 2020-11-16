import "database/sql"
import _ "go-sql-driver/mysql"


// Configure the database connection (always check errors)
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")



// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
err := db.Ping()
