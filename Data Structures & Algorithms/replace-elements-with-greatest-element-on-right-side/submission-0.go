func replaceElements(arr []int) []int {
	// Define maxRight for last element
	maxRight := -1
	// Loop through arr from last to first
	for i:= len(arr)-1; i>=0; i-- {
		// Store the current element
		current:= arr[i]

		// Assign the max to the current element
		arr[i] = maxRight

		// Check if the current element is bigger than the maxRight
		if current > maxRight{
			// Assign the current ele to maxRight if it is big
			maxRight = current
		}
	}
	return arr
}
