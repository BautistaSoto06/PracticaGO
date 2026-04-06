package app

type Catalogo struct {
	libros    []Libro
	editorial string
}

func (c Catalogo) GetEditorial() string {
	return c.editorial
}

func (c Catalogo) GetLibros() []Libro {
	return c.libros
}

func (c *Catalogo) AgregarLibro(libro Libro) {
	c.libros = append(c.libros, libro)
}