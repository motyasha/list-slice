package slice

import (
	"errors"
	"fmt"
	"reflect"
)

type Slice struct {
	sl []any
}

func (s *Slice) Add(data any) (index int64, err error) {
	if (s.sl[0] != nil) && (reflect.TypeOf(s.sl[0]) != reflect.TypeOf(data)) {
		return 0, errors.New("wrong type")
	}
	s.sl = append(s.sl, data)
	return int64(len(s.sl) - 1), nil
}

func (s *Slice) Delete(index int64) (ok bool) {
	if index > int64(len(s.sl)-1) {
		fmt.Println("Invalid index")
		return false
	}

	copy(s.sl[index:], s.sl[index+1:])
	s.sl[len(s.sl)-1] = 0
	s.sl = s.sl[:len(s.sl)-1]
	return true
}

func (s *Slice) Print() {
	fmt.Println(s.sl)
}

func (s *Slice) Get(index int64) (data any) {
	//TODO check index
	return s.sl[index]
}

func (s *Slice) Sort(more func(i, j any) bool) {
	for i := 0; i < len(s.sl); i++ {
		max := s.sl[0]
		iMax := 0
		for j := 0; j < len(s.sl)-i; j++ {
			if more(s.sl[j], max) {
				max = s.sl[j]
				iMax = j
			}
		}
		s.sl[len(s.sl)-i-1], s.sl[iMax] = s.sl[iMax], s.sl[len(s.sl)-i-1]
	}
}