package main

//目标把 " " 替换成 "%20" 字符
//字符串底层实现为数组、替换时会发生数组扩容
//数组扩容导致内容会一次 copy 影响效率

//如何在不数组扩容的前提下完成字符串替换？
//提前遍历计算出空格的数量、最终算出字符串的长度

func Replace(data string) (result string) {

	count := calculateBlank(data)
	// "%" ----> 37
	// "0" ----> 48
	// "2" ----> 50
	var rData = []rune(data)
	var dData = make([]rune, count*2+len(rData))
	for i := len(rData) - 1; i >= 0; i-- {
		j := i + count*2
		if rData[i] == 32 {
			dData[j] = 37
			dData[j-1] = 48
			dData[j-2] = 50
			count--
		} else {
			dData[j] = rData[i]
		}
	}
	return string(dData)
}

//遍历计算字符串中的空格数
func calculateBlank(data string) int {
	count := 0
	for _, v := range data {
		if v == 32 {
			count++
		}
	}
	return count
}
