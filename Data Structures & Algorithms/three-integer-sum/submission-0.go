import "maps"
import "slices"

func threeSum(nums []int) [][]int {
	n := len(nums)
	myMap := make(map[[3]int]bool, 0)
	for ix := 0; ix < n; ix++ {
		for iy := ix+1; iy < n; iy++ {
			for iz := iy+1; iz < n; iz++ {
				if nums[ix] + nums[iy] + nums[iz] == 0 {
					out := [3]int{nums[ix],nums[iy],nums[iz]}
					sort.Ints(out[:])
					myMap[out] = true
				}
			}
		}
	}
	keys := slices.Collect(maps.Keys(myMap))

	// Como tu función pide [][]int (slice de slices), 
	// hay que convertir cada [3]int a []int
	results := make([][]int, 0, len(keys))
	for _, k := range keys {
		results = append(results, k[:])
	}

	return results
}
