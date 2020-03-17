package main

/*
给定一个仅包含数字 2-9.电话号码的字母组合 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

func letterCombinations(digits string) []string {
	var table = make(map[byte][]byte)
	table['2'] = []byte{'a', 'b', 'c'}
	table['3'] = []byte{'d', 'e', 'f'}
	table['4'] = []byte{'g', 'h', 'i'}
	table['5'] = []byte{'j', 'k', 'l'}
	table['6'] = []byte{'m', 'n', 'o'}
	table['7'] = []byte{'p', 'q', 'r', 's'}
	table['8'] = []byte{'t', 'u', 'v'}
	table['9'] = []byte{'w', 'x', 'y', 'z'}

	var result []string
	if len(digits) == 0 {
		return result
	}
	var tr = []string{""}
	for _, d := range digits {
		for _, db := range table[byte(d)] {
			for _, t := range tr {
				result = append(result, t+string(db))
			}
		}
		tr = result
		result = []string{}
	}
	return tr
}
