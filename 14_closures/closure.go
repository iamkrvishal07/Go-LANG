package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())

}

// A closure is a function that remembers and has access to variables from its outer function scope, even after the outer function has finished executing.

// count variable stack mein nahi, heap mein chala jata hai, kyunki inner function uska reference hold karta hai (closure ke through).

//  Normally:
//--------------------------------------------------------------------------
// Function ke local variables stack mein store hote hain.
// Jab function return ho jata hai, stack frame destroy ho jata hai.
// Un variables ka access khatam.

// But in this case:
// Yaha counter() function ek inner function return karta hai jo count variable ka access karta hai.
// Go compiler smart hai. Use pata chal jata hai ki count ka lifetime counter() function se zyada hai.
// Isliye Go us variable ko heap mein promote kar deta hai — so it survives even after counter() returns.

// Heap pe store hone ki wajah se count zinda rehta hai jab tak increment() function ka reference program mein kahin exist karta hai.
