package main

func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}
