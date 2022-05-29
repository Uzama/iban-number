package validator

import (
	"errors"
	"testing"
)

var v = NewValidator()

func TestValidator(t *testing.T) {
	tableTest := []struct {
		name           string
		ibanNumber     string
		expectedOutput bool
		expectedError  error
	}{
		{
			name:           "happy path",
			ibanNumber:     "GB33BUKB20201555555555",
			expectedOutput: true,
			expectedError:  nil,
		},
		{
			name:           "not provided country",
			ibanNumber:     "LK33BUKB20201555555555",
			expectedOutput: false,
			expectedError:  errors.New("information not exists for given country code"),
		},
		{
			name:           "not matched length",
			ibanNumber:     "GB33BUKB202015555555557",
			expectedOutput: false,
			expectedError:  nil,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := v.Validate(testCase.ibanNumber)

			// check the error whether what we expect
			if err != nil && testCase.expectedError == nil {
				t.Errorf("expected %v, but got %v", testCase.expectedError, err)
			}

			if err == nil && testCase.expectedError != nil {
				t.Errorf("expected %v, but got %v", testCase.expectedError, err)
			}

			if err != nil && testCase.expectedError != nil {
				if err.Error() != testCase.expectedError.Error() {
					t.Errorf("expected %s, but got %s", testCase.expectedError.Error(), err.Error())
				}
			}

			// check the result whether what we expect
			if result != testCase.expectedOutput {
				t.Errorf("expected %v, but got %v", testCase.expectedOutput, result)
			}
		})
	}
}
