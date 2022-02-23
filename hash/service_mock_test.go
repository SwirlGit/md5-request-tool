package hash

import "context"

// can't use gomock, because it's beyond Go standard library
// so implement small mock

type responseGetterMock struct {
	ret1 []byte
	ret2 error
}

func newResponseGetterMock(ret1 []byte, ret2 error) *responseGetterMock {
	return &responseGetterMock{ret1: ret1, ret2: ret2}
}

func (m *responseGetterMock) GetResponseBody(_ context.Context, _ string) ([]byte, error) {
	return m.ret1, m.ret2
}
