func twoSum(nums []int, target int) []int {
    
	arrLength := len(nums)

    for i := 0; i < arrLength; i++ {
		for j := i+1; j < count; j++ {
			if(nums[i] + nums[j] == target) return [i,j]
		}
	}
}