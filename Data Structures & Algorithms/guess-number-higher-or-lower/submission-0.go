/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    low, high := 1, n
	for low <= high {
		mid := (low + high) / 2
		switch {
			case guess(mid) > 0:
				low = mid + 1
			case guess(mid) < 0:
				high = mid -1
			default:
				return mid
		}
	}
	return -1
}
