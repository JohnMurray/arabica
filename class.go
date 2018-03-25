package aribica

// Class represents a concrete or abstract Java class
type Class struct {
	name       string
	modifiers  []Modifier
	attributes []Attribute
	interfaces []Interface
	parent     *Class
	comment    *Comment
}

// NewClass allocates a new, empty class with only the name set.
func NewClass(name string) *Class {
	return &Class{
		name:       name,
		modifiers:  make([]Modifier, 0, 4),
		attributes: make([]Attribute, 0, 4),
		interfaces: make([]Interface, 0, 4),
		parent:     nil,
		comment:    nil,
	}
}

// Validate checks to see if the class (and it's children) have been
// constructed correctly. This allows us to delcaratively build our
// objects and defer validation until the end.
func (c *Class) Validate() error {
	if err := DetectConflicts(c.modifiers); err != nil {
		return err
	}

	return nil
}

// WithModifier appends a single modifier to the list
func (c *Class) WithModifier(m *Modifier) *Class {
	// TODO: Ensure modifiers do not conflict (public + protected, synchronized + volatile, etc)
	c.modifiers = append(c.modifiers, *m)
	return c
}
