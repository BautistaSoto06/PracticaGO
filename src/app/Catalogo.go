type Catalogo struct {
	libros 
	private string editorial
}

func (c Catalogo) GetEditorial() string {
	return c.editorial
}

func (c Catalogo) GetLibros() []Libro {
	return c.libros
}

func (c Catalogo) Agreg|arLibro(libro Libro) Catalogo {
	c.libros = append(c.libros, libro)
	return c
}


