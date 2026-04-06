package tests

import (
	app "ProyectoTp1/src/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatalogo_IniciaVacio(t *testing.T) {
	c := app.Catalogo{}

	assert.Equal(t, 0, len(c.GetLibros()), "El catalogo deberia iniciar vacio")
}

func TestCatalogo_AgregarLibro(t *testing.T) {
	c := app.Catalogo{}

	libro := app.Libro{
		Titulo:    "1984",
		Autor:     "George Orwell",
		Id:        "2",
		Disponile: true,
	}

	c.AgregarLibro(libro)

	libros := c.GetLibros()

	assert.Equal(t, 1, len(libros), "El catalogo deberia tener 1 libro")
	assert.Equal(t, "1984", libros[0].GetTitulo(), "El titulo no coincide")
	assert.Equal(t, "George Orwell", libros[0].GetAutor(), "El autor no coincide")
	assert.Equal(t, "2", libros[0].GetId(), "El id no coincide")
	assert.Equal(t, true, libros[0].GetDisponile(), "La disponibilidad no coincide")
}

func TestCatalogo_AgregarVariosLibros(t *testing.T) {
	c := app.Catalogo{}

	libro1 := app.Libro{
		Titulo:    "Clean Code",
		Autor:     "Robert C. Martin",
		Id:        "3",
		Disponile: true,
	}

	libro2 := app.Libro{
		Titulo:    "Refactoring",
		Autor:     "Martin Fowler",
		Id:        "4",
		Disponile: false,
	}

	c.AgregarLibro(libro1)
	c.AgregarLibro(libro2)

	libros := c.GetLibros()

	assert.Equal(t, 2, len(libros), "El catalogo deberia tener 2 libros")
	assert.Equal(t, "Clean Code", libros[0].GetTitulo(), "El primer libro no coincide")
	assert.Equal(t, "Refactoring", libros[1].GetTitulo(), "El segundo libro no coincide")
}

func TestCatalogo_ConservaOrdenDeInsercion(t *testing.T) {
	c := app.Catalogo{}

	libro1 := app.Libro{Titulo: "Libro A", Autor: "Autor A", Id: "1", Disponile: true}
	libro2 := app.Libro{Titulo: "Libro B", Autor: "Autor B", Id: "2", Disponile: true}
	libro3 := app.Libro{Titulo: "Libro C", Autor: "Autor C", Id: "3", Disponile: true}

	c.AgregarLibro(libro1)
	c.AgregarLibro(libro2)
	c.AgregarLibro(libro3)

	libros := c.GetLibros()

	assert.Equal(t, "Libro A", libros[0].GetTitulo(), "No conserva el orden")
	assert.Equal(t, "Libro B", libros[1].GetTitulo(), "No conserva el orden")
	assert.Equal(t, "Libro C", libros[2].GetTitulo(), "No conserva el orden")
}