package milrabliars

// DressedUpInt is an int represented as 2^(Pow2) * Simplified + 1.
type DressedUpInt struct {
	Pow2       int
	Simplified int
}

// DressUp the given number by representing it as 2^(Pow2) * Simplified + 1.
func DressUp(n int) DressedUpInt {
	var pow2 int = 1
	var simplified int = (n - 1) / 2

	for simplified%2 == 0 {
		pow2++
		simplified /= 2
	}

	return DressedUpInt{
		Pow2:       pow2,
		Simplified: simplified,
	}
}

// CalculateAPowXModN calculates a^x (mod n)
func CalculateAPowXModN(a, x, n int) int {
	var result int = 1
	for i := 0; i < x; i++ {
		result = (result * a) % n
	}
	return result
}

// IsStrongProbablePrimeToBaseADirect checks if the given number `n` is a strong
// probably prime to base `a`, i.e., if the witness a testifies that the number
// n is prime. This must be given `s` and `d` such that `n = 2^s * d + 1`.
//
// A number n is a strong probably prime to a iff any of the following hold:
// - a^d ≡ 1 (mod n)
// - there exists some r integer where 0 ≤ r < s and a^(2^r * d) ≡ -1 (mod n)
func IsStrongProbablePrimeToBaseADirect(n, a, s, d int) bool {
	for r := 0; r < s; r++ {
		res := CalculateAPowXModN(a, (1<<r)*d, n)
		if res == n-1 || (r == 0 && res == 1) {
			return true
		}
	}

	return false
}

// IsStrongProbablePrimeToBaseA checks if the given number is a strong probable
// prime to base A, ie., if the witness a testifies that the number n is prime.
// If this returns true, then n is probably prime. Otherwise, n is composite.
func IsStrongProbablePrimeToBaseA(n, a int) bool {
	dressedUp := DressUp(n)
	return IsStrongProbablePrimeToBaseADirect(n, a, dressedUp.Pow2, dressedUp.Simplified)
}

var StarWitnessesInt32 = [...]int{2, 3, 5, 7}

// CalculateIsPrimeAndLiars both determines if n is prime deterministically and
// any lying witnesses. n must be odd and >= 3. If the number is prime then it's
// not possible for a witness to lie and so the slice is nil rather than empty.
func CalculateIsPrimeAndLiars(n int) (bool, []int) {
	var firstTruthfulStarWitnessIndex int = -1
	for idx, witness := range StarWitnessesInt32 {
		if !IsStrongProbablePrimeToBaseA(n, witness) {
			firstTruthfulStarWitnessIndex = idx
			break
		}

		if n < 2047 {
			// checking just 2 is enough for n < 2047
			return true, nil
		}

		if idx == 1 && n < 1373653 {
			return true, nil
		}

		if idx == 2 && n < 25326001 {
			return true, nil
		}
	}

	if firstTruthfulStarWitnessIndex == -1 {
		return true, nil
	}

	liars := make([]int, 0, 64)
	var idxHint int = 0
	for a := 2; a < StarWitnessesInt32[firstTruthfulStarWitnessIndex]; a++ {
		if a == StarWitnessesInt32[idxHint] {
			// already checked
			liars = append(liars, a)
			liars = append(liars, n-a)
			idxHint++
		} else if IsStrongProbablePrimeToBaseA(n, a) {
			liars = append(liars, a)
			liars = append(liars, n-a)
		}
	}

	halfN := n / 2
	for a := StarWitnessesInt32[firstTruthfulStarWitnessIndex] + 1; a <= halfN; a++ {
		if IsStrongProbablePrimeToBaseA(n, a) {
			liars = append(liars, a)
			liars = append(liars, n-a)
		}
	}

	return false, liars
}
