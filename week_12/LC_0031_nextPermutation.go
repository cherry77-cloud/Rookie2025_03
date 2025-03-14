func nextPermutation(nums []int)  {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i + 1] {
        i -= 1
    }
    if i >= 0 {
        j := n - 1
        for j >= 0 && nums[i] >= nums[j] {
            j -= 1
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    reverse(nums[i+1:])
}

func reverse(a []int) {
    for i, n := 0, len(a); i < n / 2; i++ {
        a[i], a[n - 1 - i] = a[n - 1 - i], a[i]
    }
}
