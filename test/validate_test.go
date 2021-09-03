package test

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `validate:"min=6,max=10"`
	Age      uint8  `validate:"gte=1,lte=10"`
	Sex      string `validate:"oneof=female male"`
}

func TestValidate(t *testing.T) {
	validate := validator.New()

	user1 := User{Username: "asong", Age: 11, Sex: "null"}
	err := validate.Struct(user1)
	if err != nil {
		fmt.Println(err)
	}

	user2 := User{Username: "王志斌王志", Age: 8, Sex: "male"}
	err = validate.Struct(user2)
	if err != nil {
		fmt.Println(err)
	}

}
