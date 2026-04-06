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

func TestCatalogo_GetEditorial(t *testing.T) {
	c := app.Catalogo{}

	libro := app.Libro{
		Titulo:    "1984",
		Autor:     "George Orwell",
		Id:        "2",
		Disponile: true,
	}

	c.AgregarLibro(libro)

	assert.Equal(t, 1, len(c.GetLibros()), "El catalogo deberia tener 1 libro")
	assert.Equal(t, "1984", c.GetLibros()[0].GetTitulo(), "El titulo no coincide")
	assert.Equal(t, "George Orwell", c.GetLibros()[0].GetAutor(), "El autor no coincide")
}

func TestCatalogo_AgregarLibro(t *testing.T) {
	c := app.Catalogo{}

	libro1 := app.Libro{
		Titulo:    "Clean Code",
		Autor:     "Robert C. Martin",
		Id:        "3",
		Disponile: true,
	}

	libro2 := app.Libro{
		Titulo:    "The Pragmatic Programmer",
		Autor:     "Andrew Hunt",
		Id:        "4",
		Disponile: true,
	}

	c.AgregarLibro(libro1)
	c.AgregarLibro(libro2)

	assert.Equal(t, 2, len(c.GetLibros()), "El catalogo deberia tener 2 libros")
	assert.Equal(t, "Clean Code", c.GetLibros()[0].GetTitulo(), "El primer libro no coincide")
	assert.Equal(t, "The Pragmatic Programmer", c.GetLibros()[1].GetTitulo(), "El segundo libro no coincide")
}
