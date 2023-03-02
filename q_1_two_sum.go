package main

//https://leetcode-cn.com/problems/two-sum/


func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int, len(nums)/2)
	for  i := 0; i < len(nums); i++ {
		index,ok:=tmpMap[target-nums[i]]
		if ok {
			return []int{index,i}
		}
		tmpMap[nums[i]]=i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for  i := 0; i < len(nums) - 1; i++ {
		for  j := i+1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i,j}
			}
		}
	}
	return nil
}


func main(){

}