package main

import (
	"encoding/json"
	"testing"
)

type Test struct {
	RowOne string
	RowTwo string
	Want	int
}

func TestMain(t *testing.T) {
	tests := []Test{
		{
			RowOne: "[1,1,3,1,1]",
			RowTwo: "[1,1,5,1,1]",
			Want: 1,
		},
		{
			RowOne: "[[1],[2,3,4]]",
			RowTwo: "[[1],4]",
			Want: 1,
		},
		{
			RowOne: "[9]",
			RowTwo: "[[8,7,6]]",
			Want: -1,
		},
		{
			RowOne: "[[4,4],4,4]",
			RowTwo: "[[4,4],4,4,4]",
			Want: 1,
		},
		{
			RowOne: "[7,7,7,7]",
			RowTwo: "[7,7,7]",
			Want: -1,
		},
		{
			RowOne: "[]",
			RowTwo: "[3]",
			Want: 1,
		},
		{
			RowOne: "[[[]]]",
			RowTwo: "[[]]",
			Want: -1,
		},
		{
			RowOne: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			RowTwo: "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			Want: -1,
		},
	}

	for i, test := range tests {
		jsonArr := make([]interface{}, 0)
		json.Unmarshal([]byte(test.RowOne), &jsonArr)
		jsonArrTwo := make([]interface{}, 0)
		json.Unmarshal([]byte(test.RowTwo), &jsonArrTwo)

		result := listComparisonHelper(jsonArr, jsonArrTwo)
		if result != test.Want {
			t.Errorf("Test %d failed: Expected %d but got %d", i+1, test.Want, result)
		} else {
			t.Logf("Test %d success: Got %d", i+1, test.Want)
		}
	}

}
