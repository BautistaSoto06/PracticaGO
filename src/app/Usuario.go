package app

type Usuario struct {
	Nombre   string
	Apellido string
	Legajo   string
	Libro    []Libro
}

func (u Usuario) GetNombre() string {
	return u.Nombre
}

func (u Usuario) GetApellido() string {
	return u.Apellido
}

func (u Usuario) GetLegajo() string {
	return u.Legajo
}

func (u Usuario) GetLibro() int {
	return len(u.Libro)
}

func (u *Usuario) PedirLibro(b *Biblioteca, titulo string) bool {
	for i, libro := range b.Libro {
		if libro.GetTitulo() == titulo {
			libro.Disponile = false
			u.Libro = append(u.Libro, libro)
			b.Libro = append(b.Libro[:i], b.Libro[i+1:]...)
			return true
		}
	}

	return false
}
