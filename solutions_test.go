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

func assertEq[T comparable](got, want T, t *testing.T) {
	if got != want {
		t.Error("assertion failed")
		t.Errorf("got: %v", got)
		t.Errorf("want: %v", want)
	}
}
