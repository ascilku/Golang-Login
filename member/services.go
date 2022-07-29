package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Sevices interface {
	SaveSevices(inputMember InputMember) (Member, error)
	LoginServis(loginInput LoginInput) (Member, error)
}

type sevices struct {
	respository Respository
}

func NewSevices(respository Respository) *sevices {
	return &sevices{respository}
}

func (s *sevices) SaveSevices(inputMember InputMember) (Member, error) {
	keyServices := Member{}
	keyServices.Nama = inputMember.Nama
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputMember.Password), bcrypt.MinCost)
	if err != nil {
		return keyServices, err
	} else {
		keyServices.Password = string(passwordHash)
		newSevices, err := s.respository.SaveRespository(keyServices)
		if err != nil {
			return newSevices, err
		} else {
			return newSevices, nil
		}
	}
}

func (s *sevices) LoginServis(loginInput LoginInput) (Member, error) {
	nama := loginInput.Nama
	password := loginInput.Password

	newFindByEmailServices, err := s.respository.FindByEmail(nama)
	if err != nil {
		return newFindByEmailServices, err
	} else {
		if newFindByEmailServices.Id == 0 {
			return newFindByEmailServices, errors.New("Tidak Ada Data Servis")
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(newFindByEmailServices.Password), []byte(password))
			if err != nil {
				return newFindByEmailServices, err
			} else {
				return newFindByEmailServices, nil
			}
		}
	}
}
