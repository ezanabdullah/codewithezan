package calculator

type User struct {
	Name    string
	address string
}
type student struct {
	Name  string
	Class string
}

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a / b
}
