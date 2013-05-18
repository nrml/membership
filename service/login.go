package service

import (
	"errors"
	"github.com/nrml/membership-go/models"
)

func (service *MembershipService) Login(reg models.Registration) (models.Registration, error) {
	res, err := service.tbl.Search("email='" + reg.Email + "'")
	if len(res) == 0 {
		return reg, errors.New("invalid username: " + reg.Email)
	}

	match := res[0].(models.Registration)

	if passMatch(reg.Password, match.Password) {
		reg.Password = ""
		reg.ID = match.ID
		return reg, err

	}
	return reg, errors.New("invalid password")
	
}
