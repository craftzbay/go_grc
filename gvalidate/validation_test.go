package gvalidate

import "testing"

type User struct {
	Id       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func TestValidate(t *testing.T) {
	user := User{Id: 1, Username: "username"}
	if err := Validate(user); err != nil {
		t.Errorf("%v", err.Error())
	}
}
func TestPlateNo(t *testing.T) {
	text1 := "9999уах"
	text2 := "9999УАХ"
	text3 := "9191гч"

	if !IsPlateNo(text1) {
		t.Errorf("Not valid plate no: %s", text1)
	}

	if !IsPlateNo(text2) {
		t.Errorf("Not valid plate no: %s", text2)
	}

	if !IsPlateNo(text3) {
		t.Errorf("Not valid plate no: %s", text3)
	}
}

func TestNumeric(t *testing.T) {
	text1 := "123"

	if !IsNumeric(text1) {
		t.Errorf("Not valid numeric: %s", text1)
	}

}
