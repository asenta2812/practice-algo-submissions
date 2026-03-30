func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS, countT := make(map[rune]int), make(map[rune]int)

    for i, ch := range s {
        countS[ch]++
        countT[rune(t[i])]++
    }

    for key, value := range countS {
        if value != countT[key] {
            return false
        }
    }

    return true
}
