// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LoginAdminHandlerFunc turns a function with the right signature into a login admin handler
type LoginAdminHandlerFunc func(LoginAdminParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginAdminHandlerFunc) Handle(params LoginAdminParams) middleware.Responder {
	return fn(params)
}

// LoginAdminHandler interface for that can handle valid login admin params
type LoginAdminHandler interface {
	Handle(LoginAdminParams) middleware.Responder
}

// NewLoginAdmin creates a new http.Handler for the login admin operation
func NewLoginAdmin(ctx *middleware.Context, handler LoginAdminHandler) *LoginAdmin {
	return &LoginAdmin{Context: ctx, Handler: handler}
}

/*
	LoginAdmin swagger:route POST /auth/admin_login Auth loginAdmin

Login Admin
*/
type LoginAdmin struct {
	Context *middleware.Context
	Handler LoginAdminHandler
}

func (o *LoginAdmin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLoginAdminParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
