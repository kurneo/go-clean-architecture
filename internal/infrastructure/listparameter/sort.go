package listparameter

type Sort struct {
	Field string
	Dir   string
}

func (sort *Sort) ToString() string {
	if sort.Field == "" {
		return ""
	}
	return sort.Field + " " + sort.Dir
}
