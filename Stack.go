package dsa

import "log"

type stack struct {
	array []int
}

func (s *stack) push(element int) {
	s.array = append(s.array, element)
}

func (s *stack) pop() int {
	if len(s.array) == 0 {
		log.Println("Empty Stack")
		return 0
	}
	idx := len(s.array) - 1
	resutl := s.array[idx]
	s.array = s.array[:idx]
	return resutl
}
func (s *stack) peek() int {
	if len(s.array) == 0 {
		log.Println("Empty Stack")
		return 0
	}
	idx := len(s.array) - 1
	return s.array[idx]
}

func (s stack) printElements() {
	for len(s.array) != 0 {
		log.Println(s.array[len(s.array)-1])
		s.pop()
	}
}
