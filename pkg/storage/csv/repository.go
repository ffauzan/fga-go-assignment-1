package json

import (
	"biodata/pkg/student"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type storage struct {
	students *[]Student
}

func NewStorage(fileName string) (*storage, error) {
	s := new(storage)

	students, err := readFromFile(fileName)
	if err != nil {
		return s, nil
	}
	s.students = students
	return s, nil
}

func readFromFile(fileName string) (*[]Student, error) {
	var students []Student = []Student{}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	// Skip first line (column name)
	if _, err := csvReader.Read(); err != nil {
		return nil, err
	}

	// Loop through record
	// record = Id,Nama,Batch,Alamat,Pekerjaan
	var i int
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		i++
		student := Student{
			Id:      i,
			Name:    record[1],
			Address: record[3],
			Job:     record[4],
			Reason:  "Because Golang is blazingly fast",
		}

		students = append(students, student)
	}

	return &students, nil

}

func (s *storage) GetStudent(id int) (*student.Student, error) {
	if id > 60 {
		return nil, fmt.Errorf("student id out of index")
	}
	students := *s.students
	targetStudent := &students[id-1]

	result := student.Student{
		Id:      targetStudent.Id,
		Name:    targetStudent.Name,
		Address: targetStudent.Address,
		Job:     targetStudent.Job,
		Reason:  targetStudent.Reason,
	}

	return &result, nil
}
