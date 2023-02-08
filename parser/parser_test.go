package parser

import (
	"testing"

	"github.com/radoslavboychev/gophercises-link/parsingerrors"
	"github.com/stretchr/testify/assert"
)

var filename string = "../.././src/ex1.html"

func TestParseHTML(t *testing.T) {

	t.Run("CASE_FAILED_NO_FILE_FOUND", func(t *testing.T) {
		_, err := ParseHTML("")

		assert.ErrorIs(t, err, parsingerrors.ErrOpeningFile)
	})

	t.Run("CASE_FAILED_END_OF_FILE", func(t *testing.T) {
		_, err := ParseHTML(".././src/ex4.html")

		assert.ErrorIs(t, err, parsingerrors.ErrEOF)

	})

	t.Run("CASE_SUCCESS_HTML_PARSED", func(t *testing.T) {
		res, err := ParseHTML(".././src/ex1.html")

		assert.NotNil(t, res)
		assert.NoError(t, err)
 
	})
}
