package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var h = NewHandler()

func TestValidateIBANNumber(t *testing.T) {
	tableTest := []struct {
		name               string
		ibanNumber         string
		expectedStatusCode int
		expectedOutput     string
	}{
		{
			name:               "happy-path",
			ibanNumber:         "GB33BUKB20201555555555",
			expectedStatusCode: http.StatusOK,
			expectedOutput:     `{"isValid":true}`,
		},
		{
			name:               "empty iban number",
			ibanNumber:         "",
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `iban number not given`,
		},
		{
			name:               "invalid country code",
			ibanNumber:         "LK33BUKB20201555555555",
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedOutput:     `information not exists for given country code`,
		},
		{
			name:               "invalid iban number",
			ibanNumber:         "GB33BUKB2020155555555",
			expectedStatusCode: http.StatusOK,
			expectedOutput:     `{"isValid":false}`,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", fmt.Sprintf("/validation?iban_number=%s", testCase.ibanNumber), nil)
			if err != nil {
				t.Fatal(err)
			}

			// we create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			// our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			h.ValidateIBANNumber(rr, req)

			// check the status code is what we expect.
			if rr.Code != testCase.expectedStatusCode {
				t.Errorf("expected %v, but got %v", testCase.expectedStatusCode, rr.Code)
			}

			// check the response body is what we expect.
			if rr.Body.String() != testCase.expectedOutput {
				t.Errorf("expected %v, but got %v", testCase.expectedOutput, rr.Body.String())
			}
		})
	}
}
