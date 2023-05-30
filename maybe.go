package maybe

// Maybe is a generic interface for optional types,
// it implements all Haskell Maybe functions.
type Maybe[T any] interface {
	IsJust() bool         //return true if the Maybe is a Just
	IsNothing() bool      //return true if the Maybe is a Nothing
	FromJust() T          //return the value of the Maybe, if it is a Just, otherwise panic
	FromMaybe(T) T        //return the value of the Maybe, if it is a Just, otherwise return the default value
	MaybeToList() []T     //returns an empty list when given Nothing or a singleton list when not given Nothing.
	Or(Maybe[T]) Maybe[T] // return self value if present or argument if not (<|> from alterantive )
}

func Match[T any, R any](m Maybe[T], just func(T) R, nothing func() R) R {
	if m.IsJust() {
		return just(m.FromJust())
	}
	return nothing()
}

type maybe[T any] struct {
	value  T
	isJust bool
}

func (j *maybe[T]) Or(other Maybe[T]) Maybe[T] {
	if j.isJust {
		return j
	}
	return other
}

func Just[T any](v T) Maybe[T] {
	return &maybe[T]{value: v, isJust: true}
}
func Nothing[T any]() Maybe[T] {
	return &maybe[T]{isJust: false}
}

func (j *maybe[T]) IsJust() bool {
	return j.isJust
}

func (j *maybe[T]) IsNothing() bool {
	return !j.isJust
}

func (j *maybe[T]) FromJust() T {
	if j.isJust {
		return j.value
	}
	panic("Maybe is Nothing")
}

func (j *maybe[T]) FromMaybe(def T) T {
	if j.isJust {
		return j.value
	}
	return def
}

func (j *maybe[T]) MaybeToList() []T {
	if j.isJust {
		return []T{j.value}
	}
	return []T{}
}

//returns Nothing on an empty list or Just a where a is the first element of the list.
func ListToMaybe[T any](list []T) Maybe[T] {
	if len(list) == 0 {
		return Nothing[T]()
	}
	return Just(list[0])
}

//CatMaybes returns a list of Just values from list of Maybes.
func CatMaybes[T any](list []Maybe[T]) []T {
	var result []T
	for _, v := range list {
		if v.IsJust() {
			result = append(result, v.FromJust())
		}
	}
	return result
}

//MapMaybe is a function is a version of map which can throw out elements.
//Only the results of f() which are Just will be in the returned list.
func MapMaybe[T, U any](f func(T) Maybe[U], list []T) []U {
	var result []U
	for _, v := range list {
		maybe := f(v)
		if maybe.IsJust() {
			result = append(result, maybe.FromJust())
		}
	}
	return result
}

//MaybeFunc returns default value if m is Nothing or result of f(m) if m is Just
func MaybeFunc[T, U any](def U, f func(T) U, m Maybe[T]) U {
	if m.IsJust() {
		return f(m.FromJust())
	}
	return def
}

//Join squashes two layers of optionality into one.
func Join[T any](m Maybe[Maybe[T]]) Maybe[T] {
	if m.IsJust() {
		return m.FromJust()
	}
	return Nothing[T]()
}

//Regular map function
func Map[T, U any](f func(T) U, list []T) []U {
	var result []U
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}
