package base

import "testing"

func TestNewObject(t *testing.T) {

	type s struct {
		Name string
		Age int
	}
	d := new(s)
	d.Age = 10
	d.Name = "tom"
	t.Log(d)
}
