package basics

type AnimalAction interface {
	Speak(words string) (any, error)
}

type Animal struct {
	Specie string
}

func (a *Animal) Speak(words string) string {
	a.Specie = "HomoSap"
	return words
}

// function declared with value receiver cannot mutate the struct
// func (a Animal) Speak(words string) string {
// 	a.Specie = "HomoSap"
// 	return words
// }
