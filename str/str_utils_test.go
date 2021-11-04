package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test code

// test Cleanup
func TestCleanup(t *testing.T) {
	assert.Equal(t, "hello", CleanUp("  hello"))
	assert.Equal(t, "hello", CleanUp("  hello  "))
	assert.Equal(t, "hello foo", CleanUp(`  hello  

	foo
	`))
	assert.Equal(t, "hello foo bar oi", CleanUp(`  hello  

	foo       bar
				oi
	`))
}
