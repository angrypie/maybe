package maybe

//===== Match section
func MatchOption[V any, R any](o Option[V], some func(V) R, none func() R) R {
	if o.IsSome() {
		return some(o.unsafeGet())
	}
	return none()
}

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	unsafeGet() T
	Or(T) Option[T]
	OrElse(func() Option[T]) Option[T]
	Unwrap() T
	UnwrapOr(T) T
	UnwrapOrElse(func() T) T
	UnwrapOrDefault() T
	Filter(func(T) bool) Option[T]
	Xor(Option[T]) Option[T]
}

type some[T any] func() T

func (s some[T]) unsafeGet() T {
	return s()
}

type none[T any] struct{}

func (n none[T]) unsafeGet() T {
	return *new(T)
}

func Some[T any](v T) Option[T] {
	return some[T](func() T {
		return v
	})
}

func None[T any]() Option[T] {
	return none[T]{}
}

//===== What can be implemented as metods
//===== Is* section

func (s some[T]) IsSome() bool {
	return true
}

func (n none[T]) IsSome() bool {
	return false
}

func (s some[T]) IsNone() bool {
	return false
}

func (n none[T]) IsNone() bool {
	return true
}

//Isert Section
//Cant insert because we cant convert none to some

//===== Filter section
func (o some[T]) Filter(f func(T) bool) Option[T] {
	if f(o.unsafeGet()) {
		return Some(o.unsafeGet())
	}
	return None[T]()
}

func (o none[T]) Filter(f func(T) bool) Option[T] {
	return None[T]()
}

//TODO what difference between  Clone/Copy/Or/Unwrap? does they copy data?
//===== Or section
func (o some[T]) Or(v T) Option[T] {
	return Some(o.unsafeGet())
}

func (o none[T]) Or(v T) Option[T] {
	return Some(v)
}

func (o some[T]) OrElse(f func() Option[T]) Option[T] {
	return Some(o.unsafeGet())
}

func (o none[T]) OrElse(f func() Option[T]) Option[T] {
	return f()
}

// ===== Unwrap section

func (o some[T]) Unwrap() T {
	return o.unsafeGet()
}

func (o none[T]) Unwrap() T {
	panic("Unwrap called on None")
}

func (o some[T]) UnwrapOr(v T) T {
	return o.unsafeGet()
}

func (o none[T]) UnwrapOr(v T) T {
	return v
}

func (o some[T]) UnwrapOrElse(f func() T) T {
	return o.unsafeGet()
}

func (o none[T]) UnwrapOrElse(f func() T) T {
	return f()
}

func (o some[T]) UnwrapOrDefault() T {
	return o.unsafeGet()
}

func (o none[T]) UnwrapOrDefault() T {
	return *new(T)
}

//===== Xor section
func (o some[T]) Xor(v Option[T]) Option[T] {
	if v.IsNone() {
		return Some(o.unsafeGet())
	}
	return None[T]() //if both some then none
}

func (n none[T]) Xor(v Option[T]) Option[T] {
	if v.IsNone() {
		return None[T]() //if both some then none
	}
	return Some(v.unsafeGet())
}
