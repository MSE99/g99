package g99

import "testing"

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

func assertEq[T comparable](got, want T, t *testing.T) {
	if got != want {
		t.Error("assertion failed")
		t.Errorf("got: %v", got)
		t.Errorf("want: %v", want)
	}
}
