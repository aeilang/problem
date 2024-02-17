package handler

import "github.com/lang/problem/internal/service"

type Handler struct {
	serv service.Servicer
}

func NewHandler(serv service.Servicer) Handler {
	return Handler{
		serv: serv,
	}
}
