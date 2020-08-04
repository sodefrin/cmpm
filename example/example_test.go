package example_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	cmpmatcher "github.com/sodefrin/cmpm"
	"github.com/sodefrin/cmpm/example"
	mock_example "github.com/sodefrin/cmpm/example/mock"
)

func TestExmple(t *testing.T) {
	mockCtr := gomock.NewController(t)
	ex := mock_example.NewMockExample(mockCtr)

	ex.EXPECT().Go(cmpmatcher.CMP(example.AAA{A: "want a", B: "want b"}))
	ex.Go(example.AAA{A: "aaaa", B: "bbb"})
	//
	// --- FAIL: TestExmple (0.00s)
	// cmpmatcher_test.go:17: Unexpected call to *mock_example.MockExample.Go([{aaaa bbb}]) at cmpmatcher/example/mock/mock.go:39 because:
	// expected call at cmpmatcher/example/cmpmatcher_test.go:16 doesn't match the argument at index 0.
	//   Got: {aaaa bbb}
	//   Want:   example.AAA{
	//    - 	A: "want a",
	//    + 	A: "aaaa",
	//    - 	B: "want b",
	//    + 	B: "bbb",
	//  }
	//  FAIL
	//  exit status 1
	//
}
