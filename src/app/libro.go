package app

type Libro struct {
	Nombre    string
	Autor     string
	Id        string
	Disponile bool
}

func (libro Libro) GetNombre() string {
	return libro.Nombre
}

func (libro Libro) GetAutor() string {
	return libro.Autor
}

func (libro Libro) GetId() string {
	return libro.Id
}

func (libro Libro) GetDisponile() bool {
	return libro.Disponile
}
