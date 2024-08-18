package player_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"player"
	"testing"
  "errors"
)
var actual_error = errors.New("Stats can't be negative")

func TestHadAGoodGame(t *testing.T) {
    // Define test cases
    tests := []struct {
        name string
        stats player.Stats
        expectedBool bool
        expectedErr  error
    }{
        // Test case 1: Good game
        {   name: "Good Game",
            stats:        player.Stats{Name: "John", Minutes: 30, Points: 25, Assists: 7, TurnOver: 3, Rebounds: 12},
            expectedBool: true,
            expectedErr:  nil,
        },
        // Test case 2: Negative stats
        {
            stats:        player.Stats{Name: "Alice", Minutes: -10, Points: 15, Assists: 5, TurnOver: 2, Rebounds: 8},  name: "negative Stats",
            expectedBool: false,
            expectedErr:  actual_error,
        },
        // Add more test cases as needed
    }

    // Iterate over test cases
    for _, tc := range tests {
        // Call HadAGoodGame function with test stats
        gotBool, gotErr := player.HadAGoodGame(tc.stats)
        assert.Equal(t,tc.expectedBool,gotBool)

        if gotErr == nil{
            t.Skip("Error is nil for", tc.name)
        }
        // Assert the returned boolean value
        if assert.Error(t,gotErr,"Test Name:  %s",tc.name){
            assert.Equal(t, tc.expectedErr, gotErr)
        }

    }
}

