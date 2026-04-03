package tests

import (
	app "ProyectoTp1/src/app"
	"testing"
)

func TestUsuario(t *testing.T) {

	l := app.Libro{
		Titulo:    "TodoPasa",
		Autor:     "Carlos Alberto",
		Id:        "43401",
		Disponile: true,
	}

	u := app.Usuario{
		Nombre:   "Pedro",
		Apellido: "Sanchez",
		Legajo:   "46460589",
	}

	b := app.Biblioteca{}
	b.AgregarLibro(l)

	//esperado := "NuncaPasa"
	//obtenido := l.GetTitulo()

	//pedir := u.PedirLibro(&b, esperado)

	//if pedir != obtenido

	if !u.PedirLibro(&b, "TodoPasa") {
		t.Fatal("El usuario no pudo pedir un libro disponible")
	}

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

	if b.LibrosDisponile() != 0 {
		t.Error("La biblioteca deberia quedarse sin libros disponibles")
	}

}

//func LibroInexistente(t *testing.T) {

//}
