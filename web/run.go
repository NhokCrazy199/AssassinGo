package web

import (
	"github.com/NhokCrazy199/AssassinGo/assassin"
	"github.com/AmyangXYZ/sweetygo"
)

var (
	daddy *assassin.Daddy
)

func init() {
	daddy = assassin.NewDaddy()
}

// Run Web GUI.
func Run() {
	app := sweetygo.New()
	app.SetTemplates("/web/templates", nil)
	SetMiddlewares(app)
	SetRouter(app)
	app.Run(":8000")
}
