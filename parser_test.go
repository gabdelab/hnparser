package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processLine_fails_on_line_with_two_elements(t *testing.T) {
	logs := logs{}
	line := []string{"2018-12-05", "11:08:14"}
	assert.NotNil(t, processLine(logs, line))
}

func Test_processLine_fails_on_line_with_four_elements(t *testing.T) {
	logs := logs{}
	line := []string{"2018-12-05", "11:08:14", "myurl", "another thing"}
	assert.NotNil(t, processLine(logs, line))
}

func Test_processLine_fails_on_badly_formatted_line(t *testing.T) {
	logs := logs{}
	line := []string{"2018-12-05-00", "11:08:14", "myurl"}
	assert.NotNil(t, processLine(logs, line))
}

func Test_processLine_succeeds_on_correct_line(t *testing.T) {
	logs := logs{}
	line := []string{"2018-12-05", "11:08:14", "myurl"}
	assert.Nil(t, processLine(logs, line))
	assert.Equal(t, logs[2018][12][05][11][8][14]["myurl"], 1)
}

func Test_processLine_can_add_higher_counter(t *testing.T) {
	logs := logs{}
	line := []string{"2018-12-05", "11:08:14", "myurl"}
	assert.Nil(t, processLine(logs, line))
	assert.Nil(t, processLine(logs, line))
	assert.Equal(t, logs[2018][12][05][11][8][14]["myurl"], 2)
}

func Test_processLine_can_add_two_entries(t *testing.T) {
	logs := logs{}
	line1 := []string{"2018-12-05", "11:08:14", "myurl"}
	line2 := []string{"2018-12-05", "11:08:14", "mysecondurl"}
	assert.Nil(t, processLine(logs, line1))
	assert.Nil(t, processLine(logs, line2))
	assert.Equal(t, len(logs[2018][12][05][11][8][14]), 2)
}
