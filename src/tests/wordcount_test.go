package apitest

import (
	"testing"
	"slackwc/api"
	"github.com/stretchr/testify/assert"
)

func TestEmptyWords(t *testing.T) {
	m := api.WordCounter("")
	assert.Equal(t, len(m), 0, "Blank list should return blank map.")
}
	
func TestMapCounts(t *testing.T) {
	m := api.WordCounter("aaa bbb bbb ccc ccc")
	assert.Equal(t, m["aaa"], 1, "map returned has invalid count")
	assert.Equal(t, m["bbb"], 2, "map returned has invalid count")
	assert.Equal(t, m["ccc"], 2, "map returned has invalid count")
}	
