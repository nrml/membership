package endpoints

import (
	"encoding/json"
	"errors"
	"github.com/nrml/membership-go/models"
	"github.com/nrml/membership-go/service"
	"github.com/nrml/server-go"
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	ref := r.Referer()

	if ref == "" {
		ref = "localhost"
	}
	service := service.NewMembershipService()
	service.Init(ref)

	var reg models.Registration

	switch r.Method {
	case "GET":
		sID := r.URL.Path[1:]
		log.Println("routing GET", sID)
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
	service := service.NewMembershipService()
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
	//TODO: add static handler
	return
}
