package milrabliars

import (
	"sort"
)

// RunningTotalLiars stores how many times any number has lied so far.
type RunningTotalLiars struct {
	// HighestNumberChecked is the highest number that this data includes. We
	// start checking at 3 and we only check odd numbers, so if this is 7, for
	// example, then we have checked 3, 5, and 7.
	HighestNumberChecked int

	// TimesLied stores how many times the number at the corresponding index
	// has lied. We don't bother checking the number 1, so index 0 and 1 are
	// always 0.
	TimesLied []int

	// Top10Liars contains the the 10 highest liar numbers at this total, in
	// descending order of number of lies. We can more efficiently update this
	// list than we could recalculate it every time we add a number. If the
	// number at index i is 0, than we have seen fewer than 10 liars total and
	// all j for j > i are also 0.
	//
	// Ties are generally broken toward the lower liar.
	Top10Liars []int
}

// NewRunningTotalLiars initializes a new blank running total liars.
func NewRunningTotalLiars() *RunningTotalLiars {
	return &RunningTotalLiars{
		HighestNumberChecked: 1,
		TimesLied:            make([]int, 256*1024), // start with 1MB
		Top10Liars:           make([]int, 10),
	}
}

// Top10Movement describes a single movement within the top 10 liars ranking
// after adding a number
type Top10Movement struct {
	// Witness is the number whose placement changed
	Witness int

	// WitnessNumLies is the number of times the witness has lied up to this
	// point
	WitnessNumLies int

	// OldIndex is the old index of the number in the top 10 liars ranking, or -1
	// if the number was not previously in the top 10 (and just joined)
	OldIndex int

	// NewIndex is the new index of the number in the top 10 liars ranking, or -1
	// if the number is no longer in the top 10 (and just left)
	NewIndex int
}

// AddNextNumber adds the data for the next number, which should be the current
// highest number checked + 2 (this will verify) to this data and returns the
// movements in the top 10 liars.
func (rtl *RunningTotalLiars) AddNextNumber(n int, liars []int) []Top10Movement {
	if n != rtl.HighestNumberChecked+2 {
		panic("n != rtl.HighestNumberChecked+2")
	}

	if len(liars) == 0 {
		rtl.HighestNumberChecked = n
		return make([]Top10Movement, 0)
	}

	if n > len(rtl.TimesLied) {
		newLengthByDoubling := len(rtl.TimesLied) * 2
		for n > newLengthByDoubling {
			newLengthByDoubling *= 2
		}

		newTimesLied := make([]int, newLengthByDoubling)
		copy(newTimesLied, rtl.TimesLied)
		rtl.TimesLied = newTimesLied
	}

	rtl.HighestNumberChecked = n
	possibleNewTop10Candidates := make([]int, 0, 10)
	numberOfLiesToBeCandidate := 0
	for _, top10Witness := range rtl.Top10Liars {
		if top10Witness == 0 {
			break
		}

		numberOfLiesToBeCandidate = rtl.TimesLied[top10Witness]
		possibleNewTop10Candidates = append(possibleNewTop10Candidates, top10Witness)
	}

	for _, liar := range liars {
		newNumTimesLied := rtl.TimesLied[liar] + 1
		rtl.TimesLied[liar] = newNumTimesLied
		if newNumTimesLied >= numberOfLiesToBeCandidate {
			// we have a top 10 liar if it's not already in the list
			found := false
			for _, top10Witness := range rtl.Top10Liars {
				if top10Witness == liar {
					found = true
					break
				}
			}

			if !found {
				possibleNewTop10Candidates = append(possibleNewTop10Candidates, liar)
			}
		}
	}

	sort.SliceStable(possibleNewTop10Candidates, func(i, j int) bool {
		return rtl.TimesLied[possibleNewTop10Candidates[i]] > rtl.TimesLied[possibleNewTop10Candidates[j]]
	})

	changes := make([]Top10Movement, 0)
	for newIndex, liar := range possibleNewTop10Candidates {
		if newIndex >= 10 {
			break
		}

		if rtl.Top10Liars[newIndex] == liar {
			continue
		}

		oldIndex := -1
		for i, oldLiar := range rtl.Top10Liars {
			if oldLiar == liar {
				oldIndex = i
				break
			}
		}

		changes = append(changes, Top10Movement{
			Witness:        liar,
			WitnessNumLies: rtl.TimesLied[liar],
			OldIndex:       oldIndex,
			NewIndex:       newIndex,
		})
	}

	for index, oldLiar := range rtl.Top10Liars {
		if oldLiar == 0 {
			break
		}

		newIndex := -1
		for i, liar := range possibleNewTop10Candidates {
			if i >= 10 {
				break
			}

			if oldLiar == liar {
				newIndex = i
				break
			}
		}

		if newIndex == -1 {
			changes = append(changes, Top10Movement{
				Witness:        oldLiar,
				WitnessNumLies: rtl.TimesLied[oldLiar],
				OldIndex:       index,
				NewIndex:       -1,
			})
		}
	}

	copy(rtl.Top10Liars, possibleNewTop10Candidates[:10])
	return changes
}
