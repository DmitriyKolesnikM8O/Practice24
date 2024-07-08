package service

import "github.com/julienschmidt/httprouter"

type Service interface {
	Register(router *httprouter.Router)
}
