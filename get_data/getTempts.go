package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// To get the value data from the website, we have to create a data struct that matches the JSON file
// then we can put the data into it and use the data after that. Which is what the "Response" struct is for.

// Response struct encapsulates the structure of the JSON data
type Response struct {
	Data struct {
		TempReading struct {
			Time  float64 `json:"time"`
			Value string  `json:"value"`
		} `json:"[2][temp]"`
	} `json:"data"`
}

// The url for Arduino 5.
const url = "https://iotwebserver/api/get-device?device=5"

func main() {
	// We send the url for the specific Arduino we want to get the temperature form.
	result := getTimeAndTempt(url)

	// We then print "result.Data.TempReading.Value", which is equal to: [2][tempt]: value: "24.69" from the JSON body.
	fmt.Println("The time is:", int64(result.Data.TempReading.Time))
	fmt.Println("The value is:", result.Data.TempReading.Value)

	// The name of the Excel file will be: Arduino Temperaturer [DATO]. A new Excel file will then be created, once a
	// new day occurs, because the name will be different.
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02")
	excelName := "Arduino Temperaturer " + formattedTime

	// We check if the Excel file exists. If it does not, we create it and add one row of elements along the way.
	_, err := os.Stat(excelName + ".xlsx")
	if os.IsNotExist(err) {
		// We create the Excel file, and store the error message in "err".
		err = createExcel(excelName, result.Data.TempReading.Time, result.Data.TempReading.Value)
		if err != nil {
			fmt.Println("An error occurred:", err)
			os.Exit(1)
		} else {
			fmt.Println("Excel file created successfully!")
		}

		// If it does exist, we update the file.
	} else {
		err = updateExcel(excelName, result.Data.TempReading.Time, result.Data.TempReading.Value)
		if err != nil {
			fmt.Println("An error occurred:", err)
			os.Exit(1)
		} else {
			fmt.Println("Excel file updated successfully!")
		}
	}

}

func getTimeAndTempt(url string) Response {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("An error occurred:", err)
		}
	}(response.Body)

	// Was "ioutil.ReadAll", but it is deprecated, so we just use "io.ReadAll", which is actually what "ioutil.ReadAll" does now.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read the response body: %s\n", err)
	}

	// We unmarshal the body of the JSON file and store it in "result"
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %s\n", err)
	}

	return result

}

func createExcel(excelName string, arduinoTime float64, arduinoTemp string) error {
	// Create a new Excel file
	f := excelize.NewFile()

	// Create a new sheet
	index, _ := f.NewSheet("Sheet1")

	// Changes the time to a simplified version: "HH:MM".
	currentTime := time.Unix(int64(arduinoTime), 0)
	formattedTime := currentTime.Format("15:04")

	// Set value in cell A1
	err := f.SetCellValue("Sheet1", "A1", formattedTime)
	if err != nil {
		fmt.Print("Error setting value to sheet A1")
		return err
	}

	// Set value in cell B1
	err4 := f.SetCellValue("Sheet1", "B1", arduinoTemp)
	if err4 != nil {
		fmt.Print("Error setting value to sheet B1")
		return err
	}

	// Set the active sheet to Sheet1
	f.SetActiveSheet(index)

	// Save the file
	if err2 := f.SaveAs(excelName + ".xlsx"); err2 != nil {
		fmt.Printf("Failed to save Excel file: %s\n", err2)
		return err2
	}

	// Closes the Excel file and checks for error
	err3 := f.Close()
	if err3 != nil {
		return err3
	}

	// Assuming no errors occurred, we return nil.
	return nil
}

func updateExcel(excelName string, arduinoTime float64, arduinoTemp string) error {
	// Opens the file
	f, err := excelize.OpenFile(excelName + ".xlsx")
	if err != nil {
		fmt.Printf("Failed to open Excel file: %s\n", err)
		return err
	}

	// Gets the index of the first sheet, then its name.
	index := f.GetActiveSheetIndex()
	sheetName := f.GetSheetName(index)

	// Find the next empty row
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	// By finding the length of all the active rows and adding one, we get the next empty one.
	nextRow := len(rows) + 1

	// Changes the time to a simplified version: "HH:MM".
	currentTime := time.Unix(int64(arduinoTime), 0)
	formattedTime := currentTime.Format("15:04")

	// Append time and temperature to the next empty row
	err = f.SetCellValue(sheetName, fmt.Sprintf("A%d", nextRow), formattedTime)
	if err != nil {
		fmt.Println("Error setting value to sheet A#", err)
		return err
	}
	err = f.SetCellValue(sheetName, fmt.Sprintf("B%d", nextRow), arduinoTemp)
	if err != nil {
		fmt.Println("Error setting value to sheet B#", err)
		return err
	}

	// Save the file
	err = f.Save()
	if err != nil {
		fmt.Printf("Failed to save Excel file: %s\n", err)
		return err
	}

	// Closes the Excel file and checks for error
	err = f.Close()
	if err != nil {
		return err
	}

	// Assuming no errors occurred, we return nil.
	return nil
}

func createChart(excelName string) error {

	return nil
}
