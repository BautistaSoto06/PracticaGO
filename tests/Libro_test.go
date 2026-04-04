package tests

import (
	app "ProyectoTp1/src/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLibro_GetTitulo(t *testing.T) {
	libro := app.Libro{Titulo: "TodoPasa"}
	assert.Equal(t, "TodoPasa", libro.GetTitulo(), "Titulo no coincide")
}

func TestLibro_GetAutor(t *testing.T) {
	libro := app.Libro{Autor: "Carlos Alberto"}
	assert.Equal(t, "Carlos Alberto", libro.GetAutor(), "Autor no coincide")
}

func TestLibro_GetId(t *testing.T) {
	libro := app.Libro{Id: "43401"}
	assert.Equal(t, "43401", libro.GetId(), "Id no coincide")
}

func TestLibro_GetDisponile(t *testing.T) {
	libro := app.Libro{Disponile: true}
	assert.Equal(t, true, libro.GetDisponile(), "Disponile no coincide")
}
