package tests

import (
	app "ProyectoTp1/src/app"
	"testing"
)

func TestCrearLibro(t *testing.T) {

	libro := app.Libro{
		Nombre:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	if libro.GetNombre() != "TodoPasa" {
		t.Error("Nombre no coincide con el libro")
	}

	if libro.GetAutor() != "Carlos Alberto" {
		t.Error("Autor no coincide con el libro")
	}

	if libro.GetId() != "43401" {
		t.Error("Id no coincide con el libro")
	}

	if libro.GetDisponile() != true {
		t.Error("Disponile no coincide con el libro")
	}

}
