package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Envelope map[string]interface{}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body must only contain a single JSON value")

	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data Envelope, header http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range header {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// Task 1
func LongestSubstring(text string) string {
	maxSubstring := ""
	currentSubstring := ""

	set := make(map[rune]bool)

	for _, letter := range text {
		if set[letter] {
			for i, char := range currentSubstring {
				if char == letter {
					currentSubstring = currentSubstring[i+1:]
					break
				}
				delete(set, char)
			}
		}

		currentSubstring += string(letter)
		set[letter] = true

		if len(currentSubstring) > len(maxSubstring) {
			maxSubstring = currentSubstring
		}
	}

	return maxSubstring
}

// Task 2
func EmailFinder(emails string) []string {
	regex := `Email:\s*([^\s@]+@[^\s@]+\.[^\s@]+)`
	// /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(emails, -1)

	var checkedEmails []string

	for _, match := range matches {
		email := match[1]
		checkedEmails = append(checkedEmails, email)

	}
	return checkedEmails
}

func ReadParam(r *http.Request, paramName string) (int, error) {
	params := httprouter.ParamsFromContext(r.Context())
	i, err := strconv.ParseInt(params.ByName(paramName), 10, 64)
	text := fmt.Sprintf("invalid %s parameter", paramName)
	if err != nil {
		return -1, errors.New(text)
	}
	return int(i), nil
}
