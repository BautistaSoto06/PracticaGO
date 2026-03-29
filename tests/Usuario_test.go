package tests

import (
	app "ProyectoTp1/src/app"
	"testing"
)

func TestUsuario(t *testing.T) {

	u := app.Usuario{
		Nombre:   "Pedro",
		Apellido: "Sanchez",
		Legajo:   "46460589",
	}

	if u.Nombre != "Pedro" {
		t.Error("Nombre No concide")
	}

	if u.Apellido != "Sanchez" {
		t.Error("Apellido No concide")
	}

	if u.Legajo != "46460589" {
		t.Error("Legajo No concide")
	}
}
