package listparameter

import "strings"

type Parameters struct {
	Paginate
	Sorts  []Sort
	Search Search
	Params []Param
}

func (listParams *Parameters) ToStringOrder() string {
	var order []string
	for i := 0; i < len(listParams.Sorts); i++ {
		order = append(order, listParams.Sorts[i].ToString())
	}
	return strings.Join(order, ", ")
}
