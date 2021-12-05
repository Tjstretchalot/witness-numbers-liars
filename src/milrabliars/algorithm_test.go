package milrabliars_test

import (
	"fmt"
	"milrabliars/milrabliars"
	"sort"
)

func ExampleDressUp() {
	dressedUp := milrabliars.DressUp(747)
	fmt.Printf("747 = 2^%d * %d + 1", dressedUp.Pow2, dressedUp.Simplified)
	// Output:
	// 747 = 2^1 * 373 + 1
}

func ExampleDressUp_pow2Plus1() {
	dressedUp := milrabliars.DressUp(9)
	fmt.Printf("9 = 2^%d * %d + 1", dressedUp.Pow2, dressedUp.Simplified)
	// Output:
	// 9 = 2^3 * 1 + 1
}

func Example_CalculateAPowXModN() {
	fmt.Printf("3^4 ≡ %d (mod 7)", milrabliars.CalculateAPowXModN(3, 4, 7))
	// Output:
	// 3^4 ≡ 4 (mod 7)
}

func Example_IsStrongProbablePrimeToBaseA_91_10() {
	fmt.Printf("91 is a strong probable prime to base 10: %t", milrabliars.IsStrongProbablePrimeToBaseA(91, 10))
	// Output:
	// 91 is a strong probable prime to base 10: true
}

func Example_IsStrongProbablePrimeToBaseA_97_2() {
	fmt.Printf("97 is a strong probable prime to base 2: %t", milrabliars.IsStrongProbablePrimeToBaseA(97, 2))
	// Output:
	// 97 is a strong probable prime to base 2: true
}

func Example_IsStrongProbablePrimeToBaseA_2() {
	fmt.Printf("747 is a strong probable prime to base 23: %t", milrabliars.IsStrongProbablePrimeToBaseA(747, 23))
	// Output:
	// 747 is a strong probable prime to base 23: false
}

func Example_IsStrongProbablePrimeToBaseA_65and8() {
	fmt.Printf("65 is a strong probable prime to base 8: %t", milrabliars.IsStrongProbablePrimeToBaseA(65, 8))
	// Output:
	// 65 is a strong probable prime to base 8: true
}

func Example_IsStrongProbablePrimeToBaseA_21and8() {
	fmt.Printf("21 is a strong probable prime to base 8: %t", milrabliars.IsStrongProbablePrimeToBaseA(21, 8))
	// Output:
	// 21 is a strong probable prime to base 8: false
}

func Example_IsStrongProbablePrimeToBaseA_compositesFor8() {
	for i := 11; i < 100; i += 2 {
		// 2 does not lie for these numbers
		if !milrabliars.IsStrongProbablePrimeToBaseA(i, 2) && milrabliars.IsStrongProbablePrimeToBaseA(i, 8) {
			fmt.Printf("%d is a strong probable prime to base 8\n", i)
		}
	}
	// Output:
	// 65 is a strong probable prime to base 8
}

func Example_CalculateIsPrimeAndLiars_9() {
	prime, arr := milrabliars.CalculateIsPrimeAndLiars(9)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Printf("%t: %v", prime, arr)
	// Output:
	// false: []
}

func Example_CalculateIsPrimeAndLiars_91() {
	prime, arr := milrabliars.CalculateIsPrimeAndLiars(91)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Printf("%t: %v", prime, arr)
	// Output:
	// false: [9 10 12 16 17 22 29 38 53 62 69 74 75 79 81 82]
}

func Example_CalculateIsPrimeAndLiarsAsSlice_97() {
	prime, arr := milrabliars.CalculateIsPrimeAndLiars(97)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Printf("%t: %v", prime, arr)
	// Output:
	// true: []
}
