package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nrml/membership/models"
	"github.com/nrml/membership/service"
	"github.com/nrml/server"
	"net/http"
	"strconv"
)

type Registration struct {
	ID       int64  `json:"id" sql:"id integer not null primary key autoincrement"`
	Email    string `json:"email" sql:"email text"`
	Password string `json:"password" sql:"password text"`
}

type ErrorResponse struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	ref := r.Referer()

	if ref == "" {
		ref = "localhost"
	}
	service := service.MembershipService{}
	service.Init(ref)

	var reg models.Registration

	switch r.Method {
	case "GET":
		sID := r.URL.Path[1:]
		fmt.Println("routing GET", sID)
		var id int64 = 0

		if sID == "" {
			id = 0
		} else {
			convid, err := strconv.ParseInt(sID, 10, 64)
			id = convid

			if err != nil {
				server.JsonErr(w, err)
				return
			}
		}

		if id == 0 {
			reg, err := service.List()
			if err != nil {
				server.JsonErr(w, err)
				return
			}
			server.Json(w, reg)
			return
		} else {

			reg, err := service.Get(id)

			if err != nil {
				server.JsonErr(w, err)
				return
			}

			server.Json(w, reg)
			return
		}

	case "POST":
		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&reg)

		if err != nil {
			server.JsonErr(w, err)
			return
		}

		reg, err := service.Create(reg)

		if err != nil {
			server.JsonErr(w, err)
			return
		}

		server.Json(w, reg)
		return

	case "PUT":
		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&reg)

		if err != nil {
			server.JsonErr(w, err)
			return
		}

		reg, err = service.Update(reg)

		if err != nil {
			server.JsonErr(w, err)
			return
		}

		server.Json(w, r)
		return
	case "DELETE":
		sID := r.URL.Path[1:]
		id, err := strconv.ParseInt(sID, 10, 64)

		err = service.Delete(id)

		if err != nil {
			server.JsonErr(w, err)
			return
		}
		server.JsonStatus(w, 1, "registration deleted")
		return

	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	ref := r.Referer()

	if ref == "" {
		ref = "localhost"
	}
	service := service.MembershipService{}
	service.Init(ref)

	m := r.Method
	if m != "POST" {
		server.JsonErr(w, errors.New("must POST when logging in"))
	}

	dec := json.NewDecoder(r.Body)

	var reg models.Registration

	err := dec.Decode(reg)

	if err != nil {
		server.JsonErr(w, err)
		return
	}

	reg, err = service.Login(reg)

	if err != nil {
		server.JsonErr(w, err)
	}

	server.Json(w, reg)
	return
}
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	return
}
