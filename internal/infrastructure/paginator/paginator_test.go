package paginator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaginatorWithParamPage0(t *testing.T) {
	offset, err := ResolveOffset(0, 10)
	assert.Equal(t, 0, offset)
	assert.NotNil(t, err)
	assert.Equal(t, "page must greater than 0", err.Error())
}

func TestPaginatorResolveOffset(t *testing.T) {
	offset, err := ResolveOffset(2, 10)
	assert.Equal(t, 10, offset)
	assert.Nil(t, err)
}

func TestResolveTotalPagesWithLimit0(t *testing.T) {
	tp, e := ResolveTotalPages(100, 0)
	assert.Equal(t, 0, tp)
	assert.NotNil(t, e)
	assert.Equal(t, "limit must greater than 0", e.Error())
}

func TestResolveTotalPagesWithTotalRecords0(t *testing.T) {
	tp, e := ResolveTotalPages(0, 10)
	assert.Equal(t, 0, tp)
	assert.Nil(t, e)
}

func TestResolveTotalPagesWithTotalRecordSmallerThanLimit(t *testing.T) {
	tp, e := ResolveTotalPages(10, 100)
	assert.Equal(t, 1, tp)
	assert.Nil(t, e)
}

func TestResolveTotalPages(t *testing.T) {
	tp, e := ResolveTotalPages(100, 10)
	assert.Equal(t, 10, tp)
	assert.Nil(t, e)
}
