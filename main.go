package main

import (
	"fmt"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"example.com/sankey-golang-common-lib/components" // Import the created datetime package
	"example.com/sankey-golang-common-lib/config"
	"example.com/sankey-golang-common-lib/validations"
	"example.com/sankey-golang-common-lib/http"
	"log"
	"os"
	// "io/ioutil"
	// "net/http"
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
	url := "https://examplecom"
	err = validations.ValidateURL(url)
	if err != nil {
		fmt.Println("URL validation error:", err)
	} else {
		fmt.Println("URL is valid")
	}


	// Testing get api 
	url1 := "https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=747810ce9291121b47d6c93660d58490"
	resp, err := http.MakeRequest(url1, "GET", nil, nil)
	fmt.Println(valid)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	fmt.Println("Response:", string(resp))


	// testing post api
	url2 := "https://jsonplaceholder.typicode.com/posts"
	body := []byte(`{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`)
	response, err := http.MakeRequest(url2, "POST", nil, body)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	fmt.Println("Response:", string(response))




	//Testing encryption and decryption

	key := []byte("examplekey123456") // Define your key here or in env file

	// Test encryption
	plaintext := []byte("Hello, World!") // Define the plaintext here
	ciphertext, err := components.Encrypt(key, plaintext)
	if err != nil {
		fmt.Println("Error during encryption:", err)
	} else {
		fmt.Println("Ciphertext:", ciphertext)
	}

	// Test decryption
	decryptedText, err := components.Decrypt(key, ciphertext)
	if err != nil {
		fmt.Println("Error during decryption:", err)
	} else {
		fmt.Println("Decrypted text:", string(decryptedText))
	}


	//testing error handler
	err := errors.New("example error")
	if handleError := components.HandleError(err, "An error occurred:"); handleError != nil {
		// Handle the error
		fmt.Println("handleError", handleError)
	}


}
