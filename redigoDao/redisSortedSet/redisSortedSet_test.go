package redisSortedSet

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/SwathiAR/redigoDao"
)

func TestAddSortedSet(t *testing.T) {
	isSet, err := AddSortedSet("hackers", 1, "xyz")
	assert.True(t, isSet)
	assert.Nil(t, err)

	sortedSet, err := GetSortedSet("hackers")
	assert.Nil(t, err)
	assert.Equal(t, sortedSet, []string{"xyz"})

	redigoDao.DeleteAllKeys()

}

func TestGetSortedSet(t *testing.T) {
	AddSortedSet("hackers", 1, "xyz")
	AddSortedSet("hackers", 2, "abc")
	AddSortedSet("hackers", 3, "pqr")

	sortedSet, err := GetSortedSet("hackers")
	assert.Nil(t, err)
	assert.Equal(t, sortedSet, []string{"xyz", "abc", "pqr" })

	redigoDao.DeleteAllKeys()

}

func TestGetRangeOfSet(t *testing.T) {
	AddSortedSet("hackers", 1, "xyz")
	AddSortedSet("hackers", 2, "abc")
	AddSortedSet("hackers", 3, "pqr")
	sortedSet, err := GetRangeOfSet("hackers", 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, sortedSet, []string{"xyz", "abc"})

	sortedSet1, err1 := GetRangeOfSet("hackers", 1, 2)
	assert.Nil(t, err1)
	assert.Equal(t, sortedSet1, []string{"abc", "pqr"})

	sortedSet2, err2 := GetRangeOfSet("hackers", 0, 2)
	assert.Nil(t, err2)
	assert.Equal(t, sortedSet2, []string{"xyz", "abc", "pqr"})
	redigoDao.DeleteAllKeys()

}