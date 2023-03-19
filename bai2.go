package main

import "fmt"

// input "{}"==>true còn "{)" false
func check(s string) bool {
	switch s {
	case "{}", "}{":
		return true
	case "()", ")(":
		return true
	case "[]", "][":
		return true
	default:
		return false
	}

}
func bai2() {
	b := check("[]")
	fmt.Println(b)
}
