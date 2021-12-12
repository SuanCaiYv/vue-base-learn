package config

import (
	"fmt"
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	obj := NewConfiguration()
	fmt.Println(obj.Roles)
}
