package cmpmatcher

import (
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

type matcher struct {
	y    interface{}
	opts []cmp.Option
	diff string
}

func (m *matcher) Matches(x interface{}) bool {
	m.diff = cmp.Diff(m.y, x, m.opts...)
	return cmp.Equal(m.y, x, m.opts...)
}

func (m *matcher) String() string {
	return m.diff
}

func CMP(y interface{}, opts ...cmp.Option) gomock.Matcher {
	return &matcher{
		y:    y,
		opts: opts,
	}
}
