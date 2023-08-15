package database

import (
	"testing"
)

func (t *testing.T) {
 _, err := DBSet()
 if err != nil {
	t.Fatal("DBSet() failed with error: %v", err)
 }
}