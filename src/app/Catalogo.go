package app

type Catalogo struct {
	libros    []Libro
}

func (c Catalogo) GetLibros() []Libro {
	return c.libros
}

func (c *Catalogo) AgregarLibro(libro Libro) {
	c.libros = append(c.libros, libro)
}