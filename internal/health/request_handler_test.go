package health

import (
	"errors"
	"fmt"
	"github.com/ifaniqbal/go-base-project/internal/utils"
	mocks "github.com/ifaniqbal/go-base-project/test/mocks/internal_/health"
	mocks2 "github.com/ifaniqbal/go-base-project/test/mocks/internal_/utils"
	mocks3 "github.com/ifaniqbal/go-base-project/test/mocks/internal_/utils"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestRequestHandler_check(t *testing.T) {
	type fields struct {
		repo    VersionGetterRepository
		catcher utils.Catcher
	}
	type args struct {
		ctx utils.Context
	}
	type testCase struct {
		name   string
		fields fields
		args   args
	}
	var tests []testCase

	name := "error"
	mockRepo := mocks.NewVersionGetterRepository(t)
	mockCatcher := mocks3.NewCatcher(t)
	mockContext := mocks2.NewContext(t)
	mockCurrentCatcher := mocks3.NewCatcher(t)
	mockCatcher.EXPECT().WithContext(mockContext).Return(mockCurrentCatcher).Once()
	err := errors.New("e")
	mockRepo.EXPECT().GetVersion().Return("", err).Once()
	mockCurrentCatcher.EXPECT().Catch(fmt.Errorf("repo.GetVersion: %w", err)).Once()
	mockContext.EXPECT().JSON(http.StatusInternalServerError, mock.Anything).Once()
	f := fields{repo: mockRepo, catcher: mockCatcher}
	a := args{ctx: mockContext}
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
	})

	name = "success"
	version := "8"
	mockCatcher.EXPECT().WithContext(mockContext).Return(mockCurrentCatcher).Once()
	mockRepo.EXPECT().GetVersion().Return(version, nil).Once()
	mockContext.EXPECT().JSON(http.StatusOK, gin.H{
		"message":      "System is healthy",
		"mysqlVersion": version,
	}).Once()
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := RequestHandler{
				repo:    tt.fields.repo,
				catcher: tt.fields.catcher,
			}
			h.check(tt.args.ctx)
		})
	}
}
