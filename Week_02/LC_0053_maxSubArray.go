func maxSubArray(nums []int) int {
    ans := math.MinInt    // 注意答案可以是负数，不能初始化成 0
    f := 0
    for _, x := range nums {
        f = max(f, 0) + x
        ans = max(ans, f)
    }
    return ans
}
