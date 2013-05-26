package service

import (
	"errors"
	"github.com/nrml/membership-go/models"
	"log"
)

func (service *MembershipService) Login(reg models.Registration) (models.Registration, error) {
	log.Printf("trying to log in: %v\n", reg)
	res, err := service.tbl.Search("email='" + reg.Email + "'")
	if len(res) == 0 {
		return reg, errors.New("invalid username: " + reg.Email)
	}

	match := res[0].(models.Registration)
	log.Printf("match: %v\n", match)
	if passMatch(reg.Password, match.Password) {
		reg.Password = ""
		reg.ID = match.ID
		return reg, err

	}
	log.Println("bad password")

	return reg, errors.New("invalid password")

}
