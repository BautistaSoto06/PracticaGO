package tests

import (
	app "ProyectoTp1/src/app"
	"testing"
)

func TestUsuario(t *testing.T) {

	l := app.Libro{
		Nombre:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	u := app.Usuario{
		Nombre:   "Pedro",
		Apellido: "Sanchez",
		Legajo:   "46460589",
	}

	u.PedirLibro(l)

	if u.GetNombre() != "Pedro" {
		t.Error("Nombre No concide")
	}

	if u.GetApellido() != "Sanchez" {
		t.Error("Apellido No concide")
	}

	if u.GetLegajo() != "46460589" {
		t.Error("Legajo No concide")
	}

	tamaño := u.GetLibro()

	if tamaño != 1 {
		t.Error("No coincide con la cantidad que tenemos")
	}

}
