package addcountries

import (
	"testing"

	"gotest.tools/assert"
)

func TestRequest_isValid(t *testing.T) {
	t.Run("should return false when company is blank", func(t *testing.T) {
		req := Request{
			Source:  "amazon",
			Company: "",
			Country: "us",
		}
		assert.Equal(t, req.isValid(), false)
	})

	t.Run("should return true when company and country is present", func(t *testing.T) {
		req := Request{
			Source:  "amazon",
			Company: "google",
			Country: "us",
		}
		assert.Equal(t, req.isValid(), true)
	})
}
