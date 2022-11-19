package ioc

import (
	"fmt"
	"testing"
)

type Tester interface {
	Say()
}

type T struct {
	Name Tester `di:"default"`
}

type S string

func (s S) Say() {
	fmt.Println(s)
}
func TestNewContainer(t *testing.T) {

}
