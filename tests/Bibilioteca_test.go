package tests

import (
	app "ProyectoTp1/src/app"
	"testing"
)

func TestBiblio(t *testing.T) {

	l1 := app.Libro{
		Titulo:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	l2 := app.Libro{
		Titulo:    "Habia una vez",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	l3 := app.Libro{
		Titulo:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	l4 := app.Libro{
		Titulo:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	l5 := app.Libro{
		Titulo:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	u1 := app.Usuario{
		Nombre:   "Pedro",
		Apellido: "Sanchez",
		Legajo:   "46460589",
	}

	u2 := app.Usuario{
		Nombre:   "Pedro",
		Apellido: "Sanchez",
		Legajo:   "46460589",
	}

	b := app.Biblioteca{}
	b.AgregarLibro(l1)
	b.AgregarLibro(l2)
	b.AgregarLibro(l3)
	b.AgregarLibro(l4)
	b.AgregarLibro(l5)
	b.Usuario = append(b.Usuario, u1, u2)

	if b.LibrosDisponile() != 5 {
		t.Fatal("La biblioteca deberia tener cinco libros disponibles")
	}

	if !b.Usuario[0].PedirLibro(&b, "TodoPasa") {
		t.Fatal("El primer usuario no pudo pedir un libro")
	}

	if !b.Usuario[1].PedirLibro(&b, "Habia una vez") {
		t.Fatal("El segundo usuario no pudo pedir un libro")
	}

	if b.LibrosDisponile() != 3 {
		t.Error("La cantidad de libros disponibles no coincide")
	}

	if b.LibrosPrestados() != 2 {
		t.Error("La cantidad de libros prestados no coincide")
	}
}
