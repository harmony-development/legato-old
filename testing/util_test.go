package testing

import (
	"database/sql"
	"harmony-server/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var u64tsTests = []struct {
	input    uint64
	expected string
}{
	{0, "0"},
	{8, "8"},
	{999, "999"},
	{1001, "1001"},
}

var u64tsaTests = []struct {
	input    []uint64
	expected []string
}{
	{[]uint64{2, 4, 6, 8}, []string{"2", "4", "6", "8"}},
	{[]uint64{18446744073709551615, 18446744073709551614}, []string{"18446744073709551615", "18446744073709551614"}},
}

var timeTSTests = []struct {
	input    time.Time
	expected string
}{
	{
		time.Date(2009, 1, 9, 12, 0, 0, 0, time.UTC),
		"1231502400",
	},
}

var nullTimeTest1 = "1231502400"

var nullTimeTSTests = []struct {
	input    sql.NullTime
	expected *string
}{
	{
		sql.NullTime{
			Time:  time.Date(2009, 1, 9, 12, 0, 0, 0, time.UTC),
			Valid: true,
		},
		&nullTimeTest1,
	},
	{
		sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
		nil,
	},
}

func TestU64TS(t *testing.T) {
	for _, test := range u64tsTests {
		assert.Equal(t, util.U64TS(test.input), test.expected)
	}
}

func TestU64TSA(t *testing.T) {
	for _, test := range u64tsaTests {
		assert.Equal(t, util.U64TSA(test.input), test.expected)
	}
}

func TestTimeTS(t *testing.T) {
	for _, test := range timeTSTests {
		assert.Equal(t, util.TimeTS(test.input), test.expected)
	}
}

func TestNullTimeTS(t *testing.T) {
	for _, test := range nullTimeTSTests {
		s := util.NullTimeTS(test.input)
		if s != nil {
			assert.Equal(t, *s, *test.expected)
		} else {
			assert.Equal(t, s, test.expected)
		}
	}
}
