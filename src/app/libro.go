package app

type Libro struct {
	Nombre    string
	Autor     string
	Id        string
	Disponile bool
}

func (libro Libro) getNombre() string {
	return libro.Nombre
}

func (libro Libro) getAutor() string {
	return libro.Autor
}

func (libro Libro) getId() string {
	return libro.Id
}

func (libro Libro) getDisponile() bool {
	return libro.Disponile
}
