func minEatingSpeed(piles []int, h int) int {
	maxpile := piles[0]
	for _, pile := range piles {
		if maxpile < pile {
			maxpile = pile
		}
	}

	l, r := 1, maxpile
	for l < r {
       m := (l+r) / 2
	   hours := 0
	   for _, s:= range piles {
		 hours += (s + m - 1)/m
	   }

	   if hours > h  {
		 l = m + 1
	   } else {
		 r = m 
	   }
	}

	return r
}
  