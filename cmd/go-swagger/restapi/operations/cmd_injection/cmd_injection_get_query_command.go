// Code generated by go-swagger; DO NOT EDIT.

package cmd_injection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CmdInjectionGetQueryCommandHandlerFunc turns a function with the right signature into a cmd injection get query command handler
type CmdInjectionGetQueryCommandHandlerFunc func(CmdInjectionGetQueryCommandParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CmdInjectionGetQueryCommandHandlerFunc) Handle(params CmdInjectionGetQueryCommandParams) middleware.Responder {
	return fn(params)
}

// CmdInjectionGetQueryCommandHandler interface for that can handle valid cmd injection get query command params
type CmdInjectionGetQueryCommandHandler interface {
	Handle(CmdInjectionGetQueryCommandParams) middleware.Responder
}

// NewCmdInjectionGetQueryCommand creates a new http.Handler for the cmd injection get query command operation
func NewCmdInjectionGetQueryCommand(ctx *middleware.Context, handler CmdInjectionGetQueryCommandHandler) *CmdInjectionGetQueryCommand {
	return &CmdInjectionGetQueryCommand{Context: ctx, Handler: handler}
}

/*
	CmdInjectionGetQueryCommand swagger:route GET /cmdInjection/exec.Command/query/{safety} cmd_injection cmdInjectionGetQueryCommand

demonstrates Command Injection via query, with vulnerable function exec.Command
*/
type CmdInjectionGetQueryCommand struct {
	Context *middleware.Context
	Handler CmdInjectionGetQueryCommandHandler
}

func (o *CmdInjectionGetQueryCommand) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCmdInjectionGetQueryCommandParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
