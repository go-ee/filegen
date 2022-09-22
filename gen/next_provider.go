package gen

type NextProvider[T any] interface {
	Next() (ret T)
	Reset()
}

type ArrayNextProvider[T any] struct {
	Items        []T
	currentIndex int
}

func (o *ArrayNextProvider[T]) Next() (ret T) {
	if o.currentIndex < len(o.Items) {
		ret = o.Items[o.currentIndex]
		o.currentIndex++
	}
	return
}

func (o *ArrayNextProvider[T]) Reset() {
	o.currentIndex = 0
}
