package helpers

import (
	"testing"
)

type User struct {
	Id       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func TestValidate(t *testing.T) {
	user := User{}
	if err := Validate(user); err != nil {
		t.Errorf("%v", err.Error())
	}

}
