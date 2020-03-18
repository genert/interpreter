package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

// The environment is what we use to keep track of
// value by associating them with a name.
// The name “environment” is a classic one,
// used in a lot of other interpreters, especially Lispy ones.
// But even though the name may sound sophisticated,
// at its heart the environment is a hash map that associates strings with objects.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Closures are functions that “close over” the environment they
// were defined in. They carry their
// own environment around and whenever they’re called they can access it.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
