package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"github.com/adrianoff/golang-otus-homework/hw09_struct_validator/validators"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}
	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "John Doe",
				Age:    50,
				Email:  "john@example.com",
				Role:   "admin",
				Phones: []string{"11231231234"},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "John Doe",
				Age:    120,
				Email:  "example.com",
				Role:   "admin",
				Phones: []string{"11231231234"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   validators.ErrInvalidMax,
				},
				ValidationError{
					Field: "Email",
					Err:   validators.ErrInvalidRegexp,
				},
			},
		},
		{
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "John Doe",
				Age:    1,
				Email:  "john@example.com",
				Role:   "admEn",
				Phones: []string{"11231231234"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   validators.ErrInvalidMin,
				},
				ValidationError{
					Field: "Role",
					Err:   validators.ErrInvalidIn,
				},
			},
		},
		{
			in: User{
				ID:     "12345678901234567890123456789012345", // ID length is not 36
				Name:   "John Doe",
				Age:    18,
				Email:  "johndoe@example.com",
				Role:   "admin",
				Phones: []string{"11231231234"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   validators.ErrInvalidLen,
				},
			},
		},
		{
			in: User{
				ID:     "12345678901234567890123456789012345", // ID length is not 36
				Name:   "John Doe",
				Age:    18,
				Email:  "johndoe@example.com",
				Role:   "admin",
				Phones: []string{"892", "11231231234"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   validators.ErrInvalidLen,
				},
				ValidationError{
					Field: "Phones",
					Err:   validators.ErrInvalidLen,
				},
			},
		},
		{
			in: App{
				Version: "1.0.0",
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "1.0",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   validators.ErrInvalidLen,
				},
			},
		},
		{
			in: Token{
				Header:    []byte("header"),
				Payload:   []byte("payload"),
				Signature: []byte("signature"),
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				Body: "OK",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 300,
				Body: "OK",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   validators.ErrInvalidIn,
				},
			},
		},
	}

	for i, testCase := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			testCase := testCase
			t.Parallel()

			err := Validate(testCase.in)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}
