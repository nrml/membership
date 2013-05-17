package service

import (
	"github.com/jameskeane/bcrypt"
)

func encrypt(p string) string {
	hash, _ := bcrypt.Hash(p)
	return hash
}

func passMatch(p string, hash string) bool {
	if bcrypt.Match(p, hash) {
		return true
	}
	return false
}
