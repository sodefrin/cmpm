package example

type AAA struct {
	A string
	B string
}

type Example interface {
	Go(AAA) string
}
