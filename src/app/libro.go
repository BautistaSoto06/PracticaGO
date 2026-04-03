package app

type Libro struct {
	Titulo    string
	Autor     string
	Id        string
	Disponile bool
}

func (libro Libro) GetTitulo() string {
	return libro.Titulo
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
