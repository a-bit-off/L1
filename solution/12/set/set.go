package set

type set struct {
	storage map[any]struct{}
}

func New() *set {
	return &set{storage: map[any]struct{}{}}
}

func (s *set) Add(values ...any) {
	for _, value := range values {
		s.storage[value] = struct{}{}
	}
}

func (s *set) Delete(values ...any) {
	for _, value := range values {
		delete(s.storage, value)
	}
}

func (s set) Contain(data any) bool {
	_, ok := s.storage[data]
	return ok
}

func (s set) GetKeys() []any {
	keys := make([]any, 0, len(s.storage))

	for k, _ := range s.storage {
		keys = append(keys, k)
	}

	return keys
}
