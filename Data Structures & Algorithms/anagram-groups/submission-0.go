func groupAnagrams(strs []string) [][]string {
   // a-z : 26 characters 
   m := make(map[[26]int][]string)

   for _, str := range strs {
        fcount := [26]int{}
        for _, s := range str {
            fcount[s - 'a']++
        }
        m[fcount] = append(m[fcount], str)
   }

   result := [][]string{}
   for _, ss := range m {
        result = append(result, ss)
   }
   return result
}
