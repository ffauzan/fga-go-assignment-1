package student

import (
	"fmt"
	"log"
)

type Service interface {
	PrintStudent(int) error
}

type Repository interface {
	GetStudent(int) (*Student, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) PrintStudent(studentId int) error {
	student, err := s.repo.GetStudent(studentId)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Id: ", student.Id)
	fmt.Println("Nama: ", student.Name)
	fmt.Println("Alamat: ", student.Address)
	fmt.Println("Pekerjaan: ", student.Job)
	fmt.Println("Alasan memilih kelas: ", student.Reason)
	return nil
}
