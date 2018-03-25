package aribica

import (
	"fmt"
)

type Modifier int

//go:generate stringer -type=Modifier
const (
	PUBLIC Modifier = iota
	PRIVATE
	PROTECTED
	VOLATILE
	SYNCHRONIZED
	STATIC
	FINAL
	ABSTRACT
)

// DetectConflicts finds issues with multiple modifiers that may not be
// allowed (by Java) to be used together.
func DetectConflicts(modifiers []Modifier) error {
	var access *Modifier = nil
	var thread *Modifier = nil

	for _, m := range modifiers {
		// Validate accesss modifiers
		if m == PUBLIC || m == PRIVATE || m == PROTECTED {
			if access != nil {
				return fmt.Errorf("Duplicate access modifiers found: %s, %s", m.String(), access.String())
			} else {
				access = &m
			}
		}

		// Validate thread/synchronization modifiers
		if m == VOLATILE || m == SYNCHRONIZED {
			if thread != nil {
				return fmt.Errorf("Duplicate thread/synchronization modifiers found: %s, %s",
					m.String(),
					thread.String())
			} else {
				thread = &m
			}
		}
	}
	return nil
}
