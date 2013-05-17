package service

import (
	"errors"
	"fmt"
	"github.com/nrml/datalayer/sqlite"
	"github.com/nrml/membership/models"
	"log"
)

type MembershipService struct {
	Namespace string
	db        *sqlite.DB
	tbl       *sqlite.Table
}

func (service *MembershipService) Init(namespace string) {
	service.Namespace = namespace + ".membership"
	db := sqlite.DB{}
	service.db = &db
	service.db.Namespace = namespace + ".membership"
	err := service.db.Init()

	if err != nil {
		fmt.Println("fatal error initialzing database")
		log.Fatal("fatal error initializing database")
	}
	fmt.Println("sanity check for registraiton")
	sCheck(service.db)
	tbl := db.Table("Registration", models.Registration{})
	service.tbl = &tbl
}

func sCheck(localdb *sqlite.DB) {
	_, err := localdb.CreateTable("Registration", models.Registration{})
	if err != nil {
		fmt.Println("create registration table: ", err.Error())
		log.Fatal("fatal error when trying to create the registation table: " + err.Error())
	}
}

func (service *MembershipService) Create(reg models.Registration) (models.Registration, error) {
	reg.Password = encrypt(reg.Password)
	id, err := service.tbl.Create(reg)
	reg.ID = id
	reg.Password = ""
	return reg, err
}

func (service *MembershipService) Get(id int64) (models.Registration, error) {
	res, err := service.tbl.Get(id)
	if res == nil {
		return models.Registration{}, errors.New("uknown registration id")
	}
	reg := res.(models.Registration)
	return reg, err
}
func (service *MembershipService) List() ([]models.Registration, error) {
	res, err := service.tbl.List()
	reg := make([]models.Registration, len(res))
	for i, arg := range res {
		reg[i] = arg.(models.Registration)
	}
	return reg, err
}
func (service *MembershipService) Search(searchString string) ([]models.Registration, error) {
	res, err := service.tbl.Search(searchString)
	reg := make([]models.Registration, len(res))
	for i, arg := range res {
		reg[i] = arg.(models.Registration)
	}
	return reg, err
}
func (service *MembershipService) Update(reg models.Registration) (models.Registration, error) {
	reg.Password = encrypt(reg.Password)
	err := service.tbl.Update(reg.ID, reg)
	reg.Password = ""
	return reg, err
}
func (service *MembershipService) Delete(id int64) error {
	err := service.tbl.Delete(id)
	return err
}
