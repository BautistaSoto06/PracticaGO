package tests

import (
	app "ProyectoTp1/src/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiblioteca_AgregarLibro(t *testing.T) {
	b := app.Biblioteca{}
	l1 := app.Libro{Titulo: "Libro 1"}
	b.AgregarLibro(l1)

	assert.Equal(t, 1, len(b.Libro), "La biblioteca deberia tener 1 libro tras agregarlo")
	assert.Equal(t, "Libro 1", b.Libro[0].Titulo, "El libro agregado no es el correcto")
}

func TestBiblioteca_LibrosDisponile(t *testing.T) {
	b := app.Biblioteca{}
	l1 := app.Libro{Titulo: "Libro 1"}
	l2 := app.Libro{Titulo: "Libro 2"}
	
	assert.Equal(t, 0, b.LibrosDisponile(), "La biblioteca deberia tener 0 libros inicialmente")

	b.AgregarLibro(l1)
	b.AgregarLibro(l2)

	assert.Equal(t, 2, b.LibrosDisponile(), "La biblioteca deberia tener 2 libros disponibles")
}

func TestBiblioteca_LibrosPrestados(t *testing.T) {
	b := app.Biblioteca{}
	
	u1 := app.Usuario{Nombre: "Pedro"}
	u1.Libro = append(u1.Libro, app.Libro{Titulo: "Libro Prestado 1"})
	
	u2 := app.Usuario{Nombre: "Juan"}
	u2.Libro = append(u2.Libro, app.Libro{Titulo: "Libro Prestado 2"}, app.Libro{Titulo: "Libro Prestado 3"})

	b.Usuario = append(b.Usuario, u1, u2)

	assert.Equal(t, 3, b.LibrosPrestados(), "La cantidad de libros prestados debe ser 3")
}
