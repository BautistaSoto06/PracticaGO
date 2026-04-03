package app

type Biblioteca struct {
	Libro   []Libro
	Usuario []Usuario
}

func (b Biblioteca) LibrosDisponile() int {
	return len(b.Libro)
}

func (b *Biblioteca) AgregarLibro(l Libro) {
	b.Libro = append(b.Libro, l)

}

func (b Biblioteca) LibrosPrestados() int {
	total := 0

	for _, Usuario := range b.Usuario {
		total += Usuario.GetLibro()
	}

	return total
}
