package main

import (
	csvStorage "biodata/pkg/storage/csv"
	"biodata/pkg/student"
	"log"
	"os"
	"strconv"
)

func main() {
	// Create repository
	r, err := csvStorage.NewStorage("./batch5.csv")
	if err != nil {
		log.Fatalf("anjay: %s", err)
	}

	// Create Service
	s := student.NewService(r)

	// Get arg
	arg := os.Args[1]
	studentId, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}

	// Call
	s.PrintStudent(studentId)
}
