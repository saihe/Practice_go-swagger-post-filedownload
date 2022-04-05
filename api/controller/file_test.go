package controller

import (
	"app/restapi/operations"
	"fmt"
	"io/ioutil"
	"testing"

	mock_service "app/mock/service"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestNewFileController(t *testing.T) {
	assert.NotNil(t, NewFileController(nil), "Fail initialize.")
}

func TestFileControllerImpl_Generate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mock_service.NewMockFileService(ctrl)
	c := NewFileController(s)

	id := "9999"
	name := "unittest"
	n := fmt.Sprintf("%s_%s.csv", name, id)
	ef, _ := ioutil.TempFile("", n)
	defer ef.Close()
	ef.Seek(0, 0)
	b, _ := ioutil.ReadAll(ef)

	s.EXPECT().Generate(gomock.Any(), gomock.Any()).Return(n, b)

	want := operations.NewPostFileOK().WithPayload(&operations.PostFileOKBody{
		Name: n,
		Data: b,
	})

	got := c.Generate(operations.PostFileParams{
		File: operations.PostFileBody{
			ID:   &id,
			Name: &name,
		},
	})

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismach -want +got %s", diff)
	}
}
