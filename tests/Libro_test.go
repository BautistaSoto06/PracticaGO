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

func TestCatalogo_GetCategoria(t *testing.T) {
	c := app.Catalogo{
		Libro: app.Libro{
			Titulo:    "El Principito",
			Autor:     "Saint-Exupery",
			Id:        "1",
			Disponile: true,
		},
		Categoria: "Infantil",
		Estante:   "A1",
	}

	assert.Equal(t, "Infantil", c.GetCategoria(), "La categoria no coincide")
}

func TestCatalogo_GetEstante(t *testing.T) {
	c := app.Catalogo{
		Libro: app.Libro{
			Titulo:    "1984",
			Autor:     "George Orwell",
			Id:        "2",
			Disponile: true,
		},
		Categoria: "Novela",
		Estante:   "B3",
	}

	assert.Equal(t, "B3", c.GetEstante(), "El estante no coincide")
}

func TestCatalogo_HeredaDatosDeLibro(t *testing.T) {
	c := app.Catalogo{
		Libro: app.Libro{
			Titulo:    "Clean Code",
			Autor:     "Robert C. Martin",
			Id:        "3",
			Disponile: true,
		},
		Categoria: "Programacion",
		Estante:   "C2",
	}

	assert.Equal(t, "Clean Code", c.GetTitulo(), "El titulo no coincide")
	assert.Equal(t, "Robert C. Martin", c.GetAutor(), "El autor no coincide")
	assert.Equal(t, "3", c.GetId(), "El id no coincide")
	assert.Equal(t, true, c.GetDisponile(), "La disponibilidad no coincide")
}