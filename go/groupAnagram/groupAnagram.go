

func groupAnagram(strs [] string) [][] string {
	m := make(map[string][][] string)	

	for _, str := range strs {
		splitStr := strings.Split(str, "")
		sort.Strings(splitStr)
		sortedStr := strings.Join(splitStr, "")
		m[sortedStr] = append(m[sortedStr], str)
	}

	var result [][] string

	for _, keys :range m {
		result
	}
}