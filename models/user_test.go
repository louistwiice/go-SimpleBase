package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SetPassword(t *testing.T) {
	plain_password := "mikepassword"
	mockUser := User{
		ID: uuid.New().String(), Email: "mike@gmail.com", FirstName: "Mike", LastName: "Spenser", Password: plain_password,
	}

	result := mockUser.SetPassword()
	assert.Nil(t, result)
	assert.NotEqual(t, mockUser.Password, plain_password)	
}

func Test_CheckPassword(t *testing.T) {
	plain_password := "mikepassword"
	mockUser := User{
		ID: uuid.New().String(), Email: "mike@gmail.com", FirstName: "Mike", LastName: "Spenser", Password: plain_password,
	}
	
	mockUser.SetPassword()

	err := mockUser.CheckPassword(plain_password)
	assert.Nil(t, err)
}

func Test_Validate(t *testing.T) {
	mockUser_withoutemail := User{
		ID: uuid.New().String(), Email: "", FirstName: "Mike", LastName: "Spenser",
	}
	mockUser_withoutfirsttname := User{
		ID: uuid.New().String(), Email: "mike@mail.com", FirstName: "", LastName: "Spenser",
	}

	mockUser_withoutlasttname := User{
		ID: uuid.New().String(), Email: "mike@mail.com", FirstName: "Mike", LastName: "",
	}

	mockUser_withoutboth := User{
		ID: uuid.New().String(), Email: "", FirstName: "Mike", LastName: "",
	}

	err := mockUser_withoutemail.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNullField)

	err = mockUser_withoutfirsttname.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNullField)

	err = mockUser_withoutlasttname.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNullField)

	err = mockUser_withoutboth.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNullField)
	assert.Equal(t, err.Error(), ErrNullField.Error())
}