package main

import "fmt"

func main() {
	// 1. 5 == 5
	// true
	fmt.Println(5 == 5)

	// 2. 10 != 3
	// true
	fmt.Println(10 != 3)

	// 3. 7 > 12
	// false
	fmt.Println(7 > 12)

	// 4. 15 < 20
	// true
	fmt.Println(15 < 20)

	// 5. 8 >= 8
	// true
	fmt.Println(8 >= 8)

	// 6. 6 <= 4
	// false
	fmt.Println(6 <= 4)

	// 7. (10 > 5) && (3 < 1)
	// false
	fmt.Println((10 > 5) && (3 < 1))

	// 8. (10 > 5) || (3 < 1)
	// true
	fmt.Println((10 > 5) || (3 < 1))

	// 9. !(5 == 5)
	// false
	fmt.Println(!(5 == 5))

	// 10. !(7 < 3)
	// true
	fmt.Println(!(7 < 3))

	// 11. true && false
	// false
	fmt.Println(true && false)

	// 12. false || false
	// false
	fmt.Println(false || false)

	// 13. true || false
	// true
	fmt.Println(true || false)

	// 14. (4 + 6 == 10) && (9 > 2)
	// true
	fmt.Println((4+6 == 10) && (9 > 2))

	// 15. (12 / 3 == 4) || (8 < 5)
	// true
	fmt.Println((12/3 == 4) || (8 < 5))

	var age int = 20
	var hasTicket bool = true

	canEnter := (age >= 18) && hasTicket

	fmt.Println(canEnter)

	var isLoggedIn bool = true
	var isAdmin bool = true

	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)
	fmt.Println(hasAccess)
}

//
