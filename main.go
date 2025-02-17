package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type EmailJSON struct {
	Verified   []string `json:"verified"`
	Unverified []string `json:"unverified"`
}

func main() {
	inputPath := "internal/data/empty-verified-and-unverified.csv"
	outputPath := "empty-verified-and-unverified.csv"

	// 1) Open the input CSV
	inFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Failed to open input CSV: %v", err)
	}
	defer inFile.Close()

	// 2) Create CSV reader
	reader := csv.NewReader(inFile)

	// 3) Read header row (assuming your CSV has a header)
	headers, err := reader.Read()
	if err != nil {
		log.Fatalf("Failed to read header row: %v", err)
	}
	fmt.Printf("Headers: %v\n", headers)

	// Find the index of the "emails" column
	emailsIndex := -1
	for i, h := range headers {
		if h == "emails" {
			emailsIndex = i
			break
		}
	}
	if emailsIndex == -1 {
		log.Fatalf("Couldn't find an 'emails' column in headers!")
	}

	// 4) Prepare output CSV file
	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output CSV: %v", err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Write the same header to the output file
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Failed to write headers to output CSV: %v", err)
	}

	// 5) Iterate over each data row
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Skipping row due to read error: %v\n", err)
			continue
		}

		// Get the JSON from the "emails" column
		emailsCol := row[emailsIndex]

		var ej EmailJSON
		if err := json.Unmarshal([]byte(emailsCol), &ej); err != nil {
			log.Printf("Skipping row due to JSON parse error: %v\n", err)
			continue
		}

		// 6) Check if BOTH verified and unverified are empty
		if len(ej.Verified) == 0 && len(ej.Unverified) == 0 {
			// Write the entire row to our output CSV
			if err := writer.Write(row); err != nil {
				log.Printf("Error writing row to CSV: %v\n", err)
			}
		}
	}

	fmt.Println("Done! See empty-verified-and-unverified.csv for the filtered records.")
}
