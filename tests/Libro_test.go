package tests

import (
	app "ProyectoTp1/src/app"
	"testing"
)

func TestCrearLibro(t *testing.T) {

	libro := app.Libro{
		Titulo:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	if libro.GetTitulo() != "TodoPasa" {
		t.Error("Titulo no coinciden")
	}

	if libro.GetAutor() != "Carlos Alberto" {
		t.Error("Autor no coinciden")
	}

	if libro.GetId() != "43401" {
		t.Error("Id no coincide")
	}
	if libro.GetDisponile() != true {
		t.Error("Disponile no coincide")
	}

}
