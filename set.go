package main

// Set implementation in Go using a map
type Set map[int]struct{}

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(val int) {
    s[val] = struct{}{}
}

func (s Set) Remove(val int) {
    delete(s, val)
}

func (s Set) Contains(val int) bool {
    _, exists := s[val]
    return exists
}
