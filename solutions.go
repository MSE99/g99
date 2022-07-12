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
