package slicedata

type SliceMapStringWrapper struct {
	Data []map[string]interface{}
	by func(p, q map[string]interface{}) bool
}

func (wrapper SliceMapStringWrapper) Len() int {
	return len(wrapper.Data)
}

func (wrapper SliceMapStringWrapper) Swap(i, j int) {
	wrapper.Data[i], wrapper.Data[j] = wrapper.Data[j], wrapper.Data[i]
}

func (wrapper SliceMapStringWrapper) Less(i, j int) bool {
	return wrapper.by(wrapper.Data[i], wrapper.Data[j])
}