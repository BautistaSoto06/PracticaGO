package tests

import (
	app "ProyectoTp1/src/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsuario_GetNombre(t *testing.T) {
	u := app.Usuario{Nombre: "Pedro"}
	assert.Equal(t, "Pedro", u.GetNombre(), "Nombre no coincide")
}

func TestUsuario_GetApellido(t *testing.T) {
	u := app.Usuario{Apellido: "Sanchez"}
	assert.Equal(t, "Sanchez", u.GetApellido(), "Apellido no coincide")
}

func TestUsuario_GetLegajo(t *testing.T) {
	u := app.Usuario{Legajo: "46460589"}
	assert.Equal(t, "46460589", u.GetLegajo(), "Legajo no coincide")
}

func TestUsuario_GetLibro(t *testing.T) {
	u := app.Usuario{}
	assert.Equal(t, 0, u.GetLibro(), "El usuario deberia tener 0 libros inicialmente")
	
	u.Libro = append(u.Libro, app.Libro{Titulo: "Test"})
	assert.Equal(t, 1, u.GetLibro(), "El usuario deberia tener 1 libro")
}

func TestUsuario_PedirLibro(t *testing.T) {
	l := app.Libro{
		Titulo:    "TodoPasa",
		Disponile: true,
	}

	u := app.Usuario{Nombre: "Pedro"}
	b := app.Biblioteca{}
	b.AgregarLibro(l)

	assert.Equal(t, true, u.PedirLibro(&b, "TodoPasa"), "El usuario no pudo pedir un libro disponible")
	assert.Equal(t, 1, u.GetLibro(), "No coincide con la cantidad que tenemos")
	assert.Equal(t, 0, b.LibrosDisponile(), "La biblioteca deberia quedarse sin libros disponibles")
}

func TestUsuario_PedirLibroInexistente(t *testing.T) {
	u := app.Usuario{Nombre: "Juan"}
	b := app.Biblioteca{}
	l := app.Libro{
		Titulo:    "OtroLibro",
		Disponile: true,
	}
	b.AgregarLibro(l)

	resultado := u.PedirLibro(&b, "Libro Inexistente")

	assert.Equal(t, false, resultado, "El usuario pudo pedir un libro que no existe")
	assert.Equal(t, 0, u.GetLibro(), "El usuario se quedo con un libro que no deberia tener")
}
