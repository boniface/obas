package users

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {
	result, err := GetUsers()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestGetUser(t *testing.T) {
	expected := "JEAN"
	result, err := GetUser("7896541230")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.FirstName)

}

func TestCreateUser(t *testing.T) {
	userC := domain.User{"m@gt.com", "", "JEAN", "PAUL", "MATUTO", time.Time{}}
	result, err := CreateUser(userC)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUser(t *testing.T) {
	userC := domain.User{"m@gt.com", "", "JEAN", "PAUL", "MATUTO", time.Time{}}
	result, err := UpdateUser(userC, "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUser(t *testing.T) {
	userC := domain.User{"m@gt.com", "", "JEAN", "PAUL", "MATUTO", time.Time{}}
	result, err := DeleteUser(userC)
	assert.Nil(t, err)
	assert.True(t, result)

}
func TestCreateUser2(t *testing.T) {
	user, err := excelize.OpenFile("C:/Users/Nicole Abrahams/go/src/obas/util/files/user.xlsx")
	var newUser domain.User
	if err != nil {
		fmt.Println(err)
		return
	}
	cellVal, err := user.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(cellVal)
	time.Now()
	for _, value := range cellVal {
		/***reading the first value in the first row***/
		//fmt.Println(index)
		//fmt.Println(value[2])
		newUser = domain.User{value[3], value[4], value[0], "", value[2], time.Now()}
		//for _,value1:=range value{
		//	fmt.Println(value1[0])
		//}
		fmt.Println(newUser)
	}

}
