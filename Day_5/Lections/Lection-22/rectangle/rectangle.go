package rectangle

type Rectangle struct {
	A, B  int    // Экспортируемые поля
	color string // не экспортируемое поле
}

func New(newA int, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}
