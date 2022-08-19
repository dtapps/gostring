package gostring

import (
	"math/rand"
	"time"
)

// GenerateRandom 生成count个长度length不重复的随机数
func GenerateRandom(length, count int) []int {

	fI := 0
	startStr := "1"
	endStr := "9"

	for {
		if fI+2 > length {
			break
		}
		startStr += "0"
		endStr += "9"

		fI = fI + 1
	}

	return GenerateRandomNumber(ToInt(startStr), ToInt(endStr), count)
}

// GenerateRandomNumber 生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start, end, count int) []int {

	// 范围检查
	if end < start || (end-start) < count {
		return nil
	}

	// 存放结果的slice
	nums := make([]int, 0)

	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {

		// 生成随机数
		num := r.Intn(end-start) + start

		// 查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
