package example

import (
	"testing"

	mock "./mock_example"

	"github.com/golang/mock/gomock"
)

func TestExample1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExample := mock.NewMockExample(ctrl)
	mockExample.EXPECT().Method("hoge").Return(1)

	t.Log("result:", mockExample.Method("hoge"))
}
