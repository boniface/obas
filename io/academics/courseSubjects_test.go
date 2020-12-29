package academics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/academics"
	"testing"
)

func TestCreateCourseSubject(t *testing.T) {
	var token = "eyJraWQiOiJvYmFzYXBpX29uXzE1NS4yMzguMzIuMjE5IiwiYWxnIjoiRVMyNTYifQ.eyJpc3MiOiJIQVNIQ09ERS5aTSIsImF1ZCI6IlNJVEVVU0VSUyIsImV4cCI6MTU4Mjg4NDk5MSwianRpIjoiS202cXVMMUk1NW9BUUdaYkVpYmRHZyIsImlhdCI6MTU4Mjc5ODU5MSwibmJmIjoxNTgyNzk4NDcxLCJzdWIiOiJTaXRlIEFjY2VzcyIsImVtYWlsIjoiMjE1MDM4MTQyQG15Y3B1dC5hYy56YSIsInJvbGUiOiJQUlJVLTJWWFlLIn0.41AH3ZiufbmH4P1nDNuzajK5Wpfqc2rqQjk9YigKpTqiPpEQQj-Z4tWs1Y28URv45q-CdtZ5Y6AwjqsJBl-AOg"
	obj := domain.CourseSubject{"000000", "A0000"}
	result, err := CreateCourseSubject(obj, token)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, obj, result)
}

func TestGetCourseSubject(t *testing.T) {
	result, err := GetCourseSubject("0001", "002")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestDeleteCourseSubject(t *testing.T) {
	obj := domain.CourseSubject{"AAAD-OTKLY", "AMNN-00MWQ"}
	result, err := DeleteCourseSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestGetAllCourseSubject(t *testing.T) {
	result, err := GetAllCourseSubject()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
