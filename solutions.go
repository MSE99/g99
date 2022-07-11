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

// func printList(l list) {
// 	if l == nil {
// 		return
// 	}

// 	result := ""
// 	current := l

// 	for current != nil {
// 		result += fmt.Sprintf("%v ", current.val)
// 		current = current.next
// 	}

// 	fmt.Println(result)
// }
