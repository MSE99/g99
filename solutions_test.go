package g99

import (
	"fmt"
	"testing"
)

func TestLastElement(t *testing.T) {
	scenarios := []struct {
		title        string
		l            list
		expectedOK   bool
		expectedItem int
	}{
		{
			title:        "nil list",
			l:            nil,
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "empty list",
			l:            makeList([]int{}),
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "list with 1 node",
			l:            makeList([]int{1}),
			expectedOK:   true,
			expectedItem: 1,
		},
		{
			title:        "list with 5 nodes",
			l:            makeList([]int{1, 2, 3, 4, 5}),
			expectedOK:   true,
			expectedItem: 5,
		},
		{
			title:        "list with 11 nodes",
			l:            makeList([]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 5}),
			expectedOK:   true,
			expectedItem: 5,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			item, ok := last(sce.l)
			assertEq(ok, sce.expectedOK, t)
			assertEq(item, sce.expectedItem, t)
		})
	}
}

func TestLastButOne(t *testing.T) {
	scenarios := []struct {
		title        string
		l            list
		expectedOK   bool
		expectedItem int
	}{
		{
			title:        "nil list",
			l:            nil,
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "empty list",
			l:            makeList([]int{}),
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "list with 1 node",
			l:            makeList([]int{1}),
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "list with 2 nodes",
			l:            makeList([]int{1, 2}),
			expectedOK:   true,
			expectedItem: 1,
		},
		{
			title:        "list with 5 nodes",
			l:            makeList([]int{1, 2, 3, 4, 5}),
			expectedOK:   true,
			expectedItem: 4,
		},
		{
			title:        "list with 3 nodes",
			l:            makeList([]int{1, 2, 3}),
			expectedOK:   true,
			expectedItem: 2,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			item, ok := lastButOne(sce.l)
			assertEq(item, sce.expectedItem, t)
			assertEq(ok, sce.expectedOK, t)
		})
	}
}

func TestNth(t *testing.T) {
	scenarios := []struct {
		title        string
		l            list
		pos          int
		expectedOK   bool
		expectedItem int
	}{
		{
			title:        "nil list",
			l:            nil,
			pos:          1,
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "empty list",
			l:            makeList([]int{}),
			pos:          1,
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "out of bounds",
			l:            makeList([]int{1}),
			pos:          2,
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "in bounds",
			l:            makeList([]int{1}),
			pos:          1,
			expectedOK:   true,
			expectedItem: 1,
		},
		{
			title:        "negative pos",
			l:            makeList([]int{1}),
			pos:          -1,
			expectedOK:   false,
			expectedItem: 0,
		},
		{
			title:        "10 items",
			l:            makeList([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}),
			pos:          10,
			expectedOK:   true,
			expectedItem: 100,
		},
		{
			title:        "middle of 10",
			l:            makeList([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}),
			pos:          5,
			expectedOK:   true,
			expectedItem: 50,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			item, ok := nth(sce.l, sce.pos)
			assertEq(item, sce.expectedItem, t)
			assertEq(ok, sce.expectedOK, t)
		})
	}
}

func TestCount(t *testing.T) {
	scenarios := []struct {
		title    string
		l        list
		expected int
	}{
		{
			title:    "nil list",
			l:        nil,
			expected: 0,
		},
		{
			title:    "empty list",
			l:        makeList([]int{}),
			expected: 0,
		},
		{
			title:    "list with 1 element",
			l:        makeList([]int{1}),
			expected: 1,
		},
		{
			title:    "list 5 elements",
			l:        makeList([]int{1, 2, 3, 4, 5}),
			expected: 5,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			result := count(sce.l)
			assertEq(result, sce.expected, t)
		})
	}
}

func TestReverse(t *testing.T) {
	scenarios := []struct {
		title    string
		l        list
		expected []int
	}{
		{
			title:    "nil list",
			l:        nil,
			expected: []int{},
		},
		{
			title:    "empty list",
			l:        makeList([]int{}),
			expected: []int{},
		},
		{
			title:    "three elements",
			l:        makeList([]int{1, 2, 3}),
			expected: []int{3, 2, 1},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			got := reverse(sce.l)
			assertSliceEq(listToSlice(got), sce.expected, t)
		})
	}
}

func TestPalindrome(t *testing.T) {
	scenarios := []struct {
		title    string
		in       list
		expected bool
	}{
		{
			title:    "nil list",
			in:       nil,
			expected: true,
		},
		{
			title:    "empty list",
			in:       makeList([]int{}),
			expected: true,
		},
		{
			title:    "one element",
			in:       makeList([]int{1}),
			expected: true,
		},
		{
			title:    "two different elements",
			in:       makeList([]int{1, 2}),
			expected: false,
		},
		{
			title:    "three dupes",
			in:       makeList([]int{1, 1, 1}),
			expected: true,
		},
		{
			title:    "five dupes",
			in:       makeList([]int{7, 7, 7, 7, 7}),
			expected: true,
		},
		{
			title:    "palindrome#1",
			in:       makeList([]int{1, 2, 2, 1}),
			expected: true,
		},
		{
			title:    "palindrome#2",
			in:       makeList([]int{5, 3, 5, 5, 5, 3, 5}),
			expected: true,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			result := isPalindrome(sce.in)
			assertEq(result, sce.expected, t)
		})
	}
}

func TestFlatten(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []any
		expected []int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: []int{},
		},
		{
			title:    "empty slice",
			in:       []any{},
			expected: []int{},
		},
		{
			title:    "slice with non slice elements",
			in:       []any{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			title:    "slice with one slice",
			in:       []any{[]any{1, 2, 3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			title:    "slice with embedded slices",
			in:       []any{[]any{1, []any{2, 3}}, []any{4, []any{5}}},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := flatten(sce.in)
			assertSliceEq(result, sce.expected, t)
		})
	}
}

func TestEliminate(t *testing.T) {
	scenarios := []struct {
		title    string
		in       list
		expected list
	}{
		{
			title:    "nil list",
			in:       nil,
			expected: makeList([]int{}),
		},
		{
			title:    "empty list",
			in:       makeList([]int{}),
			expected: makeList([]int{}),
		},
		{
			title:    "list with consecutive dupes",
			in:       makeList([]int{1, 1, 2, 2, 3, 3}),
			expected: makeList([]int{1, 2, 3}),
		},
		{
			title:    "list with non-consecutive dupes",
			in:       makeList([]int{1, 2, 3, 1, 2, 3}),
			expected: makeList([]int{1, 2, 3}),
		},
		{
			title:    "list with non-consecutive dupes #2",
			in:       makeList([]int{4, 5, 5, 4}),
			expected: makeList([]int{4, 5}),
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := eliminate(sce.in)
			assertSliceEq(listToSlice(result), listToSlice(sce.expected), t)
		})
	}
}

func TestPack(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		expected [][]int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: [][]int{},
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: [][]int{},
		},
		{
			title:    "slice with two elements",
			in:       []int{1, 1, 2, 2},
			expected: [][]int{{1, 1}, {2, 2}},
		},
		{
			title:    "pairs",
			in:       []int{1, 1, 2, 2, 1, 1, 3, 3, 4, 4, 5, 5},
			expected: [][]int{{1, 1}, {2, 2}, {1, 1}, {3, 3}, {4, 4}, {5, 5}},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			result := pack(sce.in)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestLenEncoding(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		expected [][2]int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: [][2]int{},
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: [][2]int{},
		},
		{
			title:    "4 items",
			in:       []int{1, 1, 2, 2},
			expected: [][2]int{{2, 1}, {2, 2}},
		},
		{
			title:    "11 items",
			in:       []int{1, 1, 2, 2, 1, 1, 1, 4, 4, 4, 4},
			expected: [][2]int{{2, 1}, {2, 2}, {3, 1}, {4, 4}},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			result := lenEncoding(sce.in)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestModdedLenEncoding(t *testing.T) {
	scenarios := []struct {
		title string

		in       []int
		expected []any
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: []any{},
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: []any{},
		},
		{
			title:    "consecutive elements #1",
			in:       []int{1, 2, 2, 3, 3},
			expected: []any{1, [2]int{2, 2}, []int{2, 3}},
		},
		{
			title:    "consecutive elements #2",
			in:       []int{1, 2, 2, 3, 3, 4, 3},
			expected: []any{1, [2]int{2, 2}, []int{2, 3}, 4, 3},
		},
		{
			title:    "consecutive elements #3",
			in:       []int{1, 2, 3, 4, 5},
			expected: []any{1, 2, 3, 4, 5},
		},
		{
			title:    "consecutive elements #4",
			in:       []int{1, 2, 3, 4, 5, 5, 5},
			expected: []any{1, 2, 3, 4, []int{3, 5}},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			result := moddedLenEncoding(sce.in)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestDecodeList(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []any
		expected []int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: []int{},
		},
		{
			title:    "empty slice",
			in:       []any{},
			expected: []int{},
		},
		{
			title:    "slice with no dupes",
			in:       []any{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			title:    "slice with dupes",
			in:       []any{1, 2, 3, 4, [2]int{2, 5}},
			expected: []int{1, 2, 3, 4, 5, 5},
		},
		{
			title: "slice with ONLY dupes",
			in: []any{
				[2]int{2, 1},
				[2]int{2, 2},
				[2]int{2, 3},
				[2]int{2, 4},
				[2]int{2, 5},
			},
			expected: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
		{
			title: "slice with dupes in middle",
			in: []any{
				1, 1,
				[2]int{2, 2},
				[2]int{2, 3},
				[2]int{2, 4},
				5, 5,
			},
			expected: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			out := decodeList(sce.in)
			assertSerializedEq(out, sce.expected, t)
		})
	}
}

func TestDupeList(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		expected []int
	}{
		{
			title:    "nill slice",
			in:       nil,
			expected: []int{},
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: []int{},
		},
		{
			title:    "1 element slice",
			in:       []int{1},
			expected: []int{1, 1},
		},
		{
			title:    "2 element slice",
			in:       []int{1, 2},
			expected: []int{1, 1, 2, 2},
		},
		{
			title:    "consecutive elements #1",
			in:       []int{1, 1, 1},
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			title:    "consecutive elements #2",
			in:       []int{1, 1, 2, 2, 3, 3},
			expected: []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3},
		},
		{
			title: "consecutive elements #3",
			in: []int{
				4, 4, 4,
				5, 5,
				2,
			},
			expected: []int{
				4, 4, 4,
				4, 4, 4,

				5, 5,
				5, 5,

				2, 2,
			},
		},
		{
			title: "consecutive elements #4",
			in: []int{
				4,

				5, 5, 5,
				5, 5, 5,

				2,
			},
			expected: []int{
				4, 4,

				5, 5, 5, 5,
				5, 5, 5, 5,
				5, 5, 5, 5,

				2, 2,
			},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := dupe(sce.in)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestDupeX(t *testing.T) {
	scenarios := []struct {
		title string

		in       []int
		expected []int
		x        int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: []int{},
			x:        1,
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: []int{},
			x:        1,
		},
		{
			title:    "x <= 0",
			in:       []int{1, 2, 3},
			expected: []int{},
			x:        0,
		},
		{
			title:    "x <= 0",
			in:       []int{1, 2, 3},
			expected: []int{},
			x:        -5,
		},
		{
			title:    "x == 0",
			in:       []int{1, 2, 3},
			expected: []int{},
			x:        0,
		},
		{
			title:    "x is 2",
			in:       []int{1, 2, 3},
			expected: []int{1, 1, 2, 2, 3, 3},
			x:        2,
		},
		{
			title:    "x is 1",
			in:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
			x:        1,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := dupeX(sce.in, sce.x)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestDropNth(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		expected []int
		nth      int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: []int{},
			nth:      0,
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: []int{},
			nth:      3,
		},
		{
			title:    "1 element & nth = 1",
			in:       []int{1},
			expected: []int{},
			nth:      1,
		},
		{
			title:    "every 3rd",
			in:       []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 2, 4, 5},
			nth:      3,
		},
		{
			title:    "every 4th",
			in:       []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{1, 2, 3, 5, 6, 7},
			nth:      4,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := dropNth(sce.in, sce.nth)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestExtract(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		expected [2][]int
		at       int
	}{
		{
			title:    "nil slice",
			in:       nil,
			expected: [2][]int{{}, {}},
			at:       0,
		},
		{
			title:    "empty slice",
			in:       []int{},
			expected: [2][]int{{}, {}},
			at:       0,
		},
		{
			title:    "three elements",
			in:       []int{1, 2, 3},
			expected: [2][]int{{1, 2, 3}, {}},
			at:       3,
		},
		{
			title:    "four elements",
			in:       []int{1, 2, 3, 4},
			expected: [2][]int{{1, 2, 3}, {4}},
			at:       3,
		},
		{
			title:    "five elements",
			in:       []int{1, 2, 3, 4, 5},
			expected: [2][]int{{1, 2}, {3, 4, 5}},
			at:       2,
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()

			left, right := extract(sce.in, sce.at)
			assertSerializedEq([2][]int{left, right}, sce.expected, t)
		})
	}
}

func TestRotate(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		expected []int
		rotation int
	}{
		{
			title:    "nil slice",
			in:       nil,
			rotation: 0,
			expected: []int{},
		},
		{
			title:    "empty slice",
			in:       []int{},
			rotation: 1,
			expected: []int{},
		},
		{
			title:    "slice with one element",
			in:       []int{1},
			rotation: 1,
			expected: []int{1},
		},
		{
			title:    "multiple elements #1",
			in:       []int{1, 2, 3, 4, 5, 6},
			rotation: 3,
			expected: []int{4, 5, 6, 1, 2, 3},
		},
		{
			title:    "multiple elements #2",
			in:       []int{1, 2, 3, 4, 5, 6},
			rotation: 5,
			expected: []int{6, 1, 2, 3, 4, 5},
		},
		{
			title:    "multiple elements #3",
			in:       []int{1, 2, 3, 4, 5, 6},
			rotation: 2,
			expected: []int{3, 4, 5, 6, 1, 2},
		},
		{
			title:    "multiple elements #4",
			in:       []int{1, 2, 3, 4, 5, 6},
			rotation: 4,
			expected: []int{5, 6, 1, 2, 3, 4},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := rotate(sce.in, sce.rotation)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestRemoveByIdx(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		idx      int
		expected []int
	}{
		{
			title:    "nil slice",
			in:       nil,
			idx:      0,
			expected: []int{},
		},
		{
			title:    "empty slice",
			in:       []int{},
			idx:      0,
			expected: []int{},
		},
		{
			title:    "negative idx",
			in:       []int{1, 2, 3, 4},
			idx:      -50,
			expected: []int{1, 2, 3, 4},
		},
		{
			title:    "element at start",
			in:       []int{1, 2, 3, 4},
			idx:      0,
			expected: []int{2, 3, 4},
		},
		{
			title:    "element at end",
			in:       []int{1, 2, 3, 4},
			idx:      3,
			expected: []int{1, 2, 3},
		},
		{
			title:    "element at middle",
			in:       []int{1, 2, 3, 4},
			idx:      1,
			expected: []int{1, 3, 4},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := removeByIdx(sce.in, sce.idx)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func TestInsert(t *testing.T) {
	scenarios := []struct {
		title    string
		in       []int
		pos      int
		item     int
		expected []int
	}{
		{
			title:    "nil slice",
			in:       nil,
			pos:      0,
			item:     99,
			expected: []int{99},
		},
		{
			title:    "empty slice",
			in:       []int{},
			pos:      0,
			item:     99,
			expected: []int{99},
		},
		{
			title:    "empty slice #2",
			in:       []int{},
			pos:      2,
			item:     99,
			expected: []int{0, 0, 99},
		},
		{
			title:    "insert at beginning",
			in:       []int{1, 2, 3},
			pos:      0,
			item:     99,
			expected: []int{99, 1, 2, 3},
		},
		{
			title:    "insert at end",
			in:       []int{1, 2, 3},
			pos:      3,
			item:     99,
			expected: []int{1, 2, 3, 99},
		},
		{
			title:    "insert at middle",
			in:       []int{1, 2, 3},
			pos:      2,
			item:     99,
			expected: []int{1, 2, 99, 3},
		},
	}

	for _, scenario := range scenarios {
		sce := scenario

		t.Run(sce.title, func(t *testing.T) {
			t.Parallel()
			result := insert(sce.in, sce.pos, sce.item)
			assertSerializedEq(result, sce.expected, t)
		})
	}
}

func assertEq[T comparable](got, want T, t *testing.T) {
	if got != want {
		t.Error("assertion failed")
		t.Errorf("got: %v", got)
		t.Errorf("want: %v", want)
	}
}

func assertSliceEq[T comparable](got, want []T, t *testing.T) {
	if len(got) != len(want) {
		t.Error("assertion failed")
		t.Errorf("got: %v", got)
		t.Errorf("want: %v", want)
		return
	}

	for idx, item := range got {
		if want[idx] != item {
			t.Error("assertion failed")
			t.Errorf("got: %v", got)
			t.Errorf("want: %v", want)
		}
	}
}

func assertSerializedEq[T any](got, want T, t *testing.T) {
	serializedGot := fmt.Sprintf("%v", got)
	serializedWant := fmt.Sprintf("%v", want)

	if serializedGot != serializedWant {
		t.Error("assertion failed")
		t.Errorf("got: %v", serializedGot)
		t.Errorf("want: %v", serializedWant)
	}
}
