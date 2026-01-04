package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func CountLinesInCsv() (int, error) {
	file, err := os.Open("db.csv")
	if err != nil {
		log.Fatal("could not read db.csv")
		return 0, err
	}

	reader := csv.NewReader(file)
	count := 0
	for {
		_, err := reader.Read()
		if err != nil {
			break
		}
		count++
	}
	return count, nil
}

func MarkTodoAsCompleted(id string) error {
	file, err := os.Open("db.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, v := range records {
		if v[0] == id {
			v[2] = "complete"
			records[i] = v
		}
	}

	file, err = os.OpenFile("db.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(records)

}
