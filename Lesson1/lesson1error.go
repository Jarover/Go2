package lesson1

import (
	"fmt"
	"time"
)

type Lesson1Error struct {
	err       error
	timestamp time.Time
}

func (e *Lesson1Error) Error() string {
	return fmt.Sprintf("Time %d error: %s", e.timestamp, e.err)
}

func (e *Lesson1Error) Unwrap() error {
	return e.err
}

func WrapLesson1Error(err error, timestamp time.Time) error {
	return &Lesson1Error{
		err:       err,
		timestamp: timestamp,
	}
}
