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
