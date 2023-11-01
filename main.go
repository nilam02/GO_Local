package main

import (
	"fmt"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"example.com/sankey-golang-common-lib/components" // Import the created datetime package
	"example.com/sankey-golang-common-lib/config"
	"example.com/sankey-golang-common-lib/validations"
	"log"
	"os"
)

func main() {

	// testing dateTime component
	timeZone := "Asia/Kolkata" // Set your desired time zone
	dateAndTime, err := components.GetDateTimeByZone(timeZone)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current date and time in", timeZone, ":" , dateAndTime)
	fmt.Println("Current time in", timeZone, ":", dateAndTime.Format("15:04:05"))
	// The time.RFC3339 layout is one of the standard formats and it is defined as "2006-01-02T15:04:05Z07:00".
	fmt.Println("Current time in", timeZone, ":", dateAndTime.Format(time.RFC3339))




	//testing db configurations
	config := dbconfig.NewDBConfig("mysql", "root", "", "127.0.0.1", "Problem", 3306)
	// Use the config to establish a database connection
	fmt.Println("DB Config:", config)
	db, err := sql.Open(config.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Attempt a query to confirm the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

		// Example SELECT query
		// rows, err := db.Query("SELECT id, statement FROM problems")
		// if err != nil {
		// 	fmt.Println("Error executing query:", err)
		// 	return
		// }
		// defer rows.Close()
	
		// for rows.Next() {
		// 	var id int
		// 	var statement string
		// 	err = rows.Scan(&id, &statement)
		// 	if err != nil {
		// 		fmt.Println("Error scanning row:", err)
		// 		return
		// 	}
		// 	fmt.Println("ID:", id, "Statement:", statement)
		// }
	
		// err = rows.Err()
		// if err != nil {
		// 	fmt.Println("Error iterating over rows:", err)
		// 	return
		// }



	file, err := os.Create("logfile.txt")
	if err != nil {
		log.Fatal("Cannot create log file", err)
	}
	defer file.Close()

	// Use the logging function
	logger := log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	components.LogMessage(logger, "This is a log message.")
	components.LogMessage(logger, "This is another log message.")


	// test validations
	email := "test@example.com"
	valid, err := validations.ValidateEmail(email)
	if err != nil {
		fmt.Println("Email validation error:", err)
	} else {
		fmt.Println("Email is valid:", valid)
	}

	password := "1some_password"
	err = validations.ValidatePassword(password)
	if err != nil {
		fmt.Println("Password validation error:", err)
	} else {
		fmt.Println("Password is valid")
	}

	// Testing numeric value validation
	value := "42"
	min := 0
	max := 100
	err = validations.ValidateNumericValue(value, min, max)
	if err != nil {
		fmt.Println("Numeric value validation error:", err)
	} else {
		fmt.Println("Numeric value is valid")
	}

	// Testing string format validation
	input := "SomeString"
	pattern := "^[A-Z][a-z]+$"
	err = validations.ValidateStringFormat(input, pattern)
	if err != nil {
		fmt.Println("String format validation error:", err)
	} else {
		fmt.Println("String format is valid")
	}

	// Testing URL validation
	url := "https://example.com"
	err = validations.ValidateURL(url)
	if err != nil {
		fmt.Println("URL validation error:", err)
	} else {
		fmt.Println("URL is valid")
	}


}
