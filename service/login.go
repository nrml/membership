package service

import (
	"errors"
	"github.com/nrml/membership/models"
)

func (service *MembershipService) Login(reg models.Registration) (models.Registration, error) {
	res, err := service.tbl.Search("email='" + reg.Email + "'")
	if len(res) == 0 {
		return reg, errors.New("invalid username: " + reg.Email)
	}

	match := res[0].(models.Registration)

	if passMatch(reg.Password, match.Password) {
		reg.Password = ""
		return reg, err

	} else {
		return reg, errors.New("invalid password")
	}
}
