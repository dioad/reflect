package reflect

import (
	"testing"
)

type Bar struct {
	Name string `json:"name,omitempty"`
}

type Foo struct {
	Bar Bar `json:"bar,omitempty"`
}

func TestName(t *testing.T) {
	result, isPointer := Name[Foo]()
	if result != "reflect.Foo" {
		t.Errorf("expected reflect.Foo got %s", result)
	}

	if isPointer {
		t.Errorf("expected isPointer=false got true")
	}
}

func TestNameWithPointer(t *testing.T) {
	result, isPointer := Name[*Foo]()
	if result != "reflect.Foo" {
		t.Errorf("expected reflect.Foo got %s", result)
	}

	if !isPointer {
		t.Errorf("expected isPointer=true got false")
	}
}

func TestHasField(t *testing.T) {
	res := HasField[string]("Name")
	if res {
		t.Errorf("expected false got true")
	}

	res = HasField[Bar]("Name")
	if !res {
		t.Errorf("expected true got false")
	}

	res = HasField[*Bar]("Name")
	if !res {
		t.Errorf("expected true got false")
	}
}
