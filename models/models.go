package models

import ()

type Registration struct {
	ID       int64  `json:"id" sql:"id integer not null primary key autoincrement"`
	Email    string `json:"email" sql:"email text"`
	Password string `json:"password" sql:"password text"`
}

type ErrorResponse struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
}
