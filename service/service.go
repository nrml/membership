package service

import (
	"errors"
	"fmt"
	"github.com/nrml/datalayer-go/sqlite"
	"github.com/nrml/membership-go/models"
	"log"
)

type ServiceStatus struct {
	Status    int8
	Namespace string
}

type membershipService struct {
	namespace string
	key       string
	db        *sqlite.DB
	tbl       *sqlite.Table
}

func NewMembershipService(key string, namespace string) *membershipService {
	log.Println("new membership service")
	svc := new(membershipService)
	svc.namespace = namespace
	svc.key = key
	return svc
}

func (service *membershipService) Init(key string, namespace string) (ServiceStatus, error) {
	log.Println("membership service init: " + key + "." + namespace)
	status := ServiceStatus{}
	status.Namespace = namespace

	if namespace == "" || key == "" {
		err := errors.New("key and namespace required")
		status.Status = -1
		return status, err
	}
	service.namespace = namespace
	service.key = key

	db := sqlite.DB{}
	service.db = &db
	service.db.Namespace = fmt.Sprintf("%s.%s.membership", key, namespace)
	log.Printf("initializing with namespace: %s", service.db.Namespace)
	err := service.db.Init()

	if err != nil {
		log.Println("fatal opening database")
		status.Status = -11
		return status, err
	}
	//log.Println("sanity check for registraiton")
	err = sCheck(service.db)

	if err != nil {
		log.Println("fatal initializing database")
		status.Status = -1
		return status, err
	}

	tbl := db.Table("Registration", models.Registration{})
	service.tbl = &tbl

	status.Status = 1

	return status, err
}

func sCheck(localdb *sqlite.DB) error {
	_, err := localdb.CreateTable("Registration", models.Registration{})
	if err != nil {
		log.Println("create registration table: ", err.Error())
	}
	return err
}

func (service *membershipService) Create(reg models.Registration) (models.Registration, error) {
	reg.Password = encrypt(reg.Password)
	id, err := service.tbl.Create(reg)
	reg.ID = id
	reg.Password = ""
	return reg, err
}

func (service *membershipService) Get(id int64) (models.Registration, error) {
	res, err := service.tbl.Get(id)
	if res == nil {
		return models.Registration{}, errors.New("uknown registration id")
	}
	reg := res.(models.Registration)
	reg.Password = ""
	return reg, err
}
func (service *membershipService) List() ([]models.Registration, error) {
	res, err := service.tbl.List()
	reg := make([]models.Registration, len(res))
	for i, arg := range res {
		reg[i] = arg.(models.Registration)
		reg[i].Password = ""
	}
	return reg, err
}
func (service *membershipService) Search(searchString string) ([]models.Registration, error) {
	res, err := service.tbl.Search(searchString)
	reg := make([]models.Registration, len(res))
	for i, arg := range res {
		reg[i] = arg.(models.Registration)
		reg[i].Password = ""
	}
	return reg, err
}
func (service *membershipService) Update(reg models.Registration) (models.Registration, error) {
	reg.Password = encrypt(reg.Password)
	err := service.tbl.Update(reg.ID, reg)
	reg.Password = ""
	return reg, err
}
func (service *membershipService) Delete(id int64) error {
	err := service.tbl.Delete(id)
	return err
}
