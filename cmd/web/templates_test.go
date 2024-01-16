package main

import (
	"testing"
	"time"

	"github.com/skantay/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	testCases := []struct {
		name string
		time time.Time
		want string
	}{
		{
			name: "UTC",
			time: time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		},
		{
			name: "Empty",
			time: time.Time{},
			want: "",
		},
		{
			name: "CET",
			time: time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, humanDate(testCase.time), testCase.want)
		})
	}
}
