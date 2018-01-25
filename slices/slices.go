package slices

func ReverseStringSlice(oldString []string) []string {
	length := len(oldString)
	newString := make([]string, length)
	for i := 0; i < length; i++ {
		newString[i] = oldString[length-i-1]
	}
	return newString
}





