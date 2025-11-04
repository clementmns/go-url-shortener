package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "123456"

func TestGenerateShortUrl(t *testing.T) {
	initialLink1 := "https://www.google.com"
	shortLink1 := GenerateShortUrl(initialLink1, UserId)

	initialLink2 := "https://clementomnes.dev/en"
	shortLink2 := GenerateShortUrl(initialLink2, UserId)

	initialLink3 := "https://www.youtube.com/watch?v=MFOVtWgKznc"
	shortLink3 := GenerateShortUrl(initialLink3, UserId)

	assert.Equal(t, shortLink1, "C3P7qmCXQ6f")
	assert.Equal(t, shortLink2, "AUyWkGgtySu")
	assert.Equal(t, shortLink3, "FgCmCpC9Evv")
}
