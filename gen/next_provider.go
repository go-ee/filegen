package gen

type NextProvider[T any] interface {
	Next() (ret T)
}

type ArrayNextProvider[T any] struct {
	Items          []T
	deliveredIndex int
}

func (o *ArrayNextProvider[T]) Next() (ret T) {
	if o.deliveredIndex < len(o.Items) {
		ret = o.Items[o.deliveredIndex]
		o.deliveredIndex++
	}
	return
}
