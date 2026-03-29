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

func (u Usuario) GetLibro() []Libro {
	return u.Libro
}
