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

func TestUnique(t *testing.T) {
	s := []string{"a", "b", "c", "c", "c", "d"}
	expected := []string{"a", "b", "c", "d"}
	assert.Equal(t, expected, Unique(s))
	assert.Equal(t, true, Contains(s, "c"))
	assert.Equal(t, false, Contains(s, ""))
	assert.Equal(t, true, ContainsIgnorecase(s, "C"))
	f, _ := FindIndex(s, "c")
	assert.Equal(t, 2, f)
	f, _ = FindIndex(s, "C")
	assert.Equal(t, -1, f)
}

func TestInsert(t *testing.T) {
	s := []string{"a", "b", "c", "c", "c", "d"}
	expected := []string{"a", "b", "c", "E", "c", "c", "d"}
	s = Insert(s, 3, "E")
	assert.Equal(t, expected, s)
	s = Remove(s, "c")
	expected = []string{"a", "b", "E", "c", "c", "d"}
	assert.Equal(t, expected, s)
	s = Remove(s, "c")
	expected = []string{"a", "b", "E", "c", "d"}
	assert.Equal(t, expected, s)
	s = RemoveFrom(s, 2)
	assert.Equal(t, []string{"a", "b", "c", "d"}, s)
}
func TestSwitchIf(t *testing.T) {
	a, b := SwitchIf(true, "a", "b")
	assert.Equal(t, "b", a)
	assert.Equal(t, "a", b)
	a, b = SwitchIf(false, "a", "b")
	assert.Equal(t, "a", a)
	assert.Equal(t, "b", b)
}