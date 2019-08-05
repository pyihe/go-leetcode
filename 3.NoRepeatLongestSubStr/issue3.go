package __NoRepeatLongestSubStr

/*
   @Create by GoLand
   @Author: hong
   @Time: 2018/7/20 14:13
   @File: issue3.go
*/

/*
Question Description:
Given a string, find the length of the longest substring without repeating characters.
*/

//my solution
func lengthOfLongestSubstring(s string) int {

	b := []byte(s)
	var length int
	index := 0
label:
	byteMap := make(map[byte]int)
	var temp []byte
	for i := index; i < len(b); i++ {
		if _, ok := byteMap[b[i]]; ok {
			if i-index > length {
				length = i - index
				temp = append(temp, b[index:i]...)
				//fmt.Println(fmt.Sprintf("b[index:i]= %v", b[index:i]))
				//fmt.Println(fmt.Sprintf("temp = %v", string(temp)))
				//fmt.Println(fmt.Sprintf("length = %v		index = %v		i = %v", length, index, i))
			}
			index += 1
			goto label
		} else {
			byteMap[b[i]] = i
			//temp = append(temp, b[i])
			if i == len(b)-1 && len(byteMap) == i-index+1 && len(byteMap) > length {
				length = len(byteMap)
			}
			//fmt.Println(fmt.Sprintf("byteMap = %v", byteMap))
		}
	}
	//fmt.Println(fmt.Sprintf("temp = %v", string(temp)))
	return length
}

//leetcode go solution
func lengthOfLongestSubstring2(str string) int {
	if len(str) == 0 {
		return 0
	}
	var (
		length int
		result int
	)

	hash := make(map[byte]int, len(str))

	for i, j := 0, 0; i < len(str); i++ {
		if v, ok := hash[str[i]]; ok && v+1 > j {
			j = v + 1
		}

		hash[str[i]] = i

		if length = i - j + 1; length > result {
			result = length
		}
	}

	return result
}
