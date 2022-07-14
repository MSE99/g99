package g99

type list *node

type node struct {
	val  int
	next *node
}

func makeList(s []int) list {
	if len(s) == 0 {
		return nil
	} else if len(s) == 1 {
		return &node{s[0], nil}
	}

	return &node{s[0], makeList(s[1:])}
}

func last(l list) (int, bool) {
	if l == nil {
		return 0, false
	} else if l.next == nil {
		return l.val, true
	}

	return last(l.next)
}

func lastButOne(l list) (int, bool) {
	if l == nil {
		return 0, false
	} else if l.next != nil && l.next.next == nil {
		return l.val, true
	}

	return lastButOne(l.next)
}

func nth(l list, pos int) (int, bool) {
	current := l
	i := 1

	for current != nil && i <= pos {
		if i == pos {
			return current.val, true
		}

		i++
		current = current.next
	}

	return 0, false
}

func count(l list) int {
	current := l
	count := 0

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func reverse(l list) list {
	items := []int{}
	traverseInRev(l, func(val int) {
		items = append(items, val)
	})
	return makeList(items)
}

func isPalindrome(l list) bool {
	items := []int{}
	traverseInRev(l, func(item int) {
		items = append(items, item)
	})

	for i := 0; i < len(items)/2; i++ {
		item := items[i]
		opposite := items[len(items)-1-i]

		if item != opposite {
			return false
		}
	}

	return true
}

// TODO: needs more work!, Solution to problem #7, not a perfect solution
// uses too much runtime type assertion and probably
// very slow
func flatten(l []any) []int {
	result := []int{}

	for _, item := range l {
		switch item := item.(type) {
		case int:
			result = append(result, item)
		case []any:
			result = append(result, flatten(item)...)
		}
	}

	return result
}

func eliminate(l list) list {
	var helper func(list, []int) list

	helper = func(l list, items []int) list {
		if l == nil {
			return makeList(items)
		}

		for _, item := range items {
			if l.val == item {
				return helper(l.next, items)
			}
		}

		return helper(l.next, append(items, l.val))
	}

	return helper(l, []int{})
}

func lenEncoding(in []int) [][2]int {
	packed := pack(in)
	result := [][2]int{}

	for _, p := range packed {
		result = append(result, [2]int{len(p), p[0]})
	}

	return result
}

func moddedLenEncoding(l []int) []any {
	encoded := lenEncoding(l)
	result := []any{}

	for _, item := range encoded {
		if item[0] == 1 {
			result = append(result, item[1])
		} else {
			result = append(result, item)
		}
	}

	return result
}

func decodeList(l []any) []int {
	result := []int{}

	for _, item := range l {
		switch t := item.(type) {
		case int:
			result = append(result, t)
		case [2]int:
			for i := 0; i < t[0]; i++ {
				result = append(result, t[1])
			}
		}
	}

	return result
}

func pack(in []int) [][]int {
	result := [][]int{}
	current := []int{}

	for _, item := range in {
		if len(current) == 0 || current[0] == item {
			current = append(current, item)
		} else if current[0] != item {
			result = append(result, current)
			current = []int{item}
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}

func dupe(in []int) []int {
	encoded := lenEncoding(in)
	result := []int{}

	for _, reps := range encoded {
		for i := 0; i < reps[0]*2; i++ {
			result = append(result, reps[1])
		}
	}

	return result
}

func dupeX(in []int, x int) []int {
	if x <= 0 {
		return []int{}
	}

	result := []int{}

	for _, item := range in {
		for i := 0; i < x; i++ {
			result = append(result, item)
		}
	}

	return result
}

func dropNth(in []int, nth int) []int {
	result := []int{}

	for i, j := 0, 1; i < len(in); i++ {
		if j < nth {
			j++
		} else {
			j = 1
			continue
		}

		result = append(result, in[i])
	}

	return result
}

func extract(in []int, at int) ([]int, []int) {
	left := []int{}
	right := []int{}

	for _, item := range in {
		if len(left) < at {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	return left, right
}

func rotate(in []int, rot int) []int {
	left, right := extract(in, rot)
	result := append(make([]int, 0, len(in)), right...)
	result = append(result, left...)
	return result
}

func removeByIdx(in []int, idx int) []int {
	if !(idx >= 0 && idx < len(in)) {
		return in
	}

	result := make([]int, 0, len(in))

	for pos, item := range in {
		if pos != idx {
			result = append(result, item)
		}
	}

	return result
}

func insert(in []int, pos, item int) []int {
	if pos+1 > len(in) {
		result := make([]int, len(in)+(pos+1-len(in)))
		copy(result, in)
		result[pos] = item
		return result
	}

	result := make([]int, 0, len(in)+1)

	for idx, original := range in {
		if idx == pos {
			result = append(result, item)
		}

		result = append(result, original)
	}

	return result
}

func traverseInRev(l list, cb func(int)) {
	if l == nil {
		return
	}

	traverseInRev(l.next, cb)
	cb(l.val)
}

func listToSlice(l list) []int {
	if l == nil {
		return []int{}
	}

	result := []int{}
	current := l

	for current != nil {
		result = append(result, current.val)
		current = current.next
	}

	return result
}
