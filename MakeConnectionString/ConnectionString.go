package main

import (
	"fmt"
	"strings"
)

func main() {
	var connectionString = CreatePostgresRiderConnectionString("localhost", 5432, "postgres", "usergres", "passgres", true)
	fmt.Println(connectionString)

	var connectionString2 = CreatePostgresRiderConnectionString("localhost", 5432, "postgres", "usergres", "passgres", false)
	fmt.Println(connectionString2)

	var connectionString3 = CreatePostgresRiderConnectionString("host.aws.teddy", 0, "postgres", "usergres", "passgres", true)
	fmt.Println(connectionString3)
}

// CreatePostgresRiderConnectionString generates a PostgreSQL JDBC connection string based on the provided inputs.
// If `sslMode` is true, SSL mode is included in the connection string; otherwise, it is omitted.
// Returns an empty string if any required input variables are invalid or missing.
func CreatePostgresRiderConnectionString(host string, port int, databaseName string, username string, password string, sslMode bool) string {
	var status, statusMessage = areConnectionStringVariablesEmpty(host, port, databaseName, username, password)
	if status {
		fmt.Println(statusMessage)
		return ""
	}

	var connectionString = "jdbc:postgresql://HOST:PORT/DATABASE_NAME?SSLMODE&username=USERNAME&password=PASSWORD"

	if sslMode {
		connectionString = strings.Replace(connectionString, "SSLMODE", "sslmode=require", -1)
	} else {
		connectionString = strings.Replace(connectionString, "SSLMODE&", "", -1)
	}

	if port == 0 {
		connectionString = strings.Replace(connectionString, "HOST:PORT", host, -1)
	} else {
		connectionString = strings.Replace(connectionString, "HOST", host, -1)
		connectionString = strings.Replace(connectionString, "PORT", fmt.Sprintf("%d", port), -1)
	}

	connectionString = strings.Replace(connectionString, "DATABASE_NAME", databaseName, -1)
	connectionString = strings.Replace(connectionString, "USERNAME", username, -1)
	connectionString = strings.Replace(connectionString, "PASSWORD", password, -1)

	return connectionString
}

func areConnectionStringVariablesEmpty(host string, port int, databaseName string, username string, password string) (bool, string) {
	var variablesAreEmpty = false
	var errorMessage = ""

	if host == "" {
		errorMessage = "Host is empty"
		variablesAreEmpty = true
		return variablesAreEmpty, errorMessage
	}

	if port < 0 {
		errorMessage = "Port is negative"
		variablesAreEmpty = true
		return variablesAreEmpty, errorMessage
	}

	if databaseName == "" {
		errorMessage = "Database name is empty"
		variablesAreEmpty = true
		return variablesAreEmpty, errorMessage
	}

	if username == "" {
		errorMessage = "Username is empty"
		variablesAreEmpty = true
		return variablesAreEmpty, errorMessage
	}

	if password == "" {
		errorMessage = "Password is empty"
		variablesAreEmpty = true
		return variablesAreEmpty, errorMessage
	}

	return variablesAreEmpty, errorMessage
}
