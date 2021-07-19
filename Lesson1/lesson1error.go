package lesson1

import (
	"fmt"
)

type Lesson1Error struct {
	err       error
	timestamp int64
}

func (e *Lesson1Error) Error() string {
	return fmt.Sprintf("Time error %d \n%s", e.timestamp, e.err)
}

func (e *Lesson1Error) Unwrap() error {
	return e.err
}

func WrapLesson1Error(err error, timestamp int64) error {
	return &Lesson1Error{
		err:       err,
		timestamp: timestamp,
	}
}
