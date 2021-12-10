package main

type TopVotedCandidate struct {
	winner, ctime []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	winner := []int{}
	ctimes := []int{}
	personsv := make([]int, len(persons))
	max := 0
	for i := 0; i < len(times); i++ {
		personsv[persons[i]]++
		if personsv[persons[i]] >= max {
			winner = append(winner, persons[i])
			ctimes = append(ctimes, times[i])
			max = personsv[persons[i]]
		}
	}
	return TopVotedCandidate{winner: winner, ctime: ctimes}
}

func binarySearch(a []int, tar int) int {
	h, t := 0, len(a)
	for h < t {
		m := (h + t) / 2
		if a[m] == tar {
			return m
		}
		if a[m] < tar {
			h = m + 1
		} else {
			t = m
		}
	}
	return h
}
func (this *TopVotedCandidate) Q(t int) int {
	return this.winner[binarySearch(this.ctime, t)]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

func main() {
	Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
}
