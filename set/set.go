package set

import "fmt"

type set struct {
	setMap map[interface{}]bool
}

func NewSet() *set {
	return &set{
		setMap: map[interface{}]bool{},
	}
}

func (s *set) Add(element interface{}) {

	if element == nil {
		return
	}

	if _, ok := s.setMap[element]; !ok {
		s.setMap[element] = true
	}
}

func (s *set) Delete(element interface{}) {
	delete(s.setMap, element)
}

func (s *set) Contains(element interface{}) bool {
	if _, ok := s.setMap[element]; ok {
		return true
	}

	return false
}

func (s *set) Print() {
	fmt.Print("\nSet: ")
	for k := range s.setMap {
		fmt.Print(" ", k)
	}
}

func (s *set) Intersect(s2 *set) *set {
	intersectSet := &set{setMap: map[interface{}]bool{}}

	for element := range s.setMap {
		if _, ok := s2.setMap[element]; ok {
			intersectSet.setMap[element] = true
		}
	}
	return intersectSet
}

func (s *set) Union(s2 *set) *set {
	unionSet := &set{setMap: map[interface{}]bool{}}

	for element := range s.setMap {
		unionSet.setMap[element] = true
	}

	for element := range s2.setMap {
		unionSet.setMap[element] = true
	}

	return unionSet
}

func (s *set) Subtract(s2 *set) *set {
	resultSet := *s
	for element := range s2.setMap {
		delete(resultSet.setMap, element)
	}
	return &resultSet
}
