package model

type Storage interface {
	Add(data any) (index int64, err error)
	Delete(index int64) (ok bool)
	Print()
	Get(index int64) (data any)
	Sort(more func(i, j any) bool)
}
