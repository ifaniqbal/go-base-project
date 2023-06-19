package api

import (
	"errors"
	"testing"

	"github.com/ifaniqbal/go-base-project/pkg/catcher"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var expectedPanic = func(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("success is not expected")
	}
}
var expectedNotPanic = func(t *testing.T) {
	if r := recover(); r != nil {
		t.Errorf("panic is not expected")
	}
}

func TestDefault(t *testing.T) {
	type funcs struct {
		initDbConn  initDbConnFunc
		initCatcher initCatcherFunc
	}
	type testCase struct {
		name  string
		funcs funcs
		want  func(t *testing.T)
	}
	var errorInitCatcher = func() (catcher.Catcher, error) {
		return nil, errors.New("e")
	}
	var successInitCatcher = func() (catcher.Catcher, error) {
		return nil, nil
	}
	var errorInitDbConn = func(dsn string) (*gorm.DB, error) {
		return nil, errors.New("e")
	}
	var successInitDbConn = func(dsn string) (*gorm.DB, error) {
		return nil, nil
	}
	var testCases = []testCase{
		{
			name: "initCatcher error",
			funcs: funcs{
				initCatcher: errorInitCatcher,
				initDbConn:  nil,
			},
			want: expectedPanic,
		},
		{
			name: "initDbConn error",
			funcs: funcs{
				initCatcher: successInitCatcher,
				initDbConn:  errorInitDbConn,
			},
			want: expectedPanic,
		},
		{
			name: "success",
			funcs: funcs{
				initCatcher: successInitCatcher,
				initDbConn:  successInitDbConn,
			},
			want: expectedNotPanic,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.want(t)
			initCatcher = tt.funcs.initCatcher
			initDbConn = tt.funcs.initDbConn
			Default()
		})
	}
}

func Test_defaultInitCatcher(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		_, err := defaultInitCatcher()()
		assert.Nil(t, err, "defaultInitCatcher")
	})
}

func Test_defaultInitDbConn(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		_, err := defaultInitDbConn()("")
		assert.NotNil(t, err, defaultInitDbConn)
	})
}
