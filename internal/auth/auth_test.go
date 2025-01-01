package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input      http.Header
		wantString string
		wantError  bool
	}{
		"Simple": {
			input: http.Header{
				"Authorization": []string{"ApiKey test"},
			},
			wantString: "test",
			wantError:  false,
		},
		"Missing Header": {
			input:      http.Header{},
			wantString: "",
			wantError:  true,
		},
		"Bad Header": {
			input: http.Header{
				"Authorization": []string{"Bad header test"},
			},
			wantString: "",
			wantError:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotString, gotError := GetAPIKey(tc.input)
			if (gotError != nil) != tc.wantError {
				if tc.wantError {
					t.Fatalf("Expected an error, did not get one")
				} else {
					t.Fatalf("Unexpected error: %s", gotError.Error())
				}
			}
			if gotString != tc.wantString {
				t.Fatalf("Wanted %s, got %s", tc.wantString, gotString)
			}
		})
	}
}
