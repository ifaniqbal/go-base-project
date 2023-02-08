package featureflag

import (
	"reflect"
	"testing"

	mocks "bitbucket.org/ifan-moladin/base-project/mocks/utils/environment"
	"bitbucket.org/ifan-moladin/base-project/utils/environment"
)

func TestNew(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		want := &OsEnvFeatureFlag{env: environment.Default()}
		if got := Default(); !reflect.DeepEqual(got, want) {
			t.Errorf("Default() = %v, want %v", got, want)
		}
	})
}

func TestEnvFeatureFlag_CanBeSkipped(t *testing.T) {
	type args struct {
		environment string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "local",
			args: args{environment: "local"},
			want: true,
		},
		{
			name: "development",
			args: args{environment: "development"},
			want: true,
		},
		{
			name: "staging",
			args: args{environment: "staging"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := OsEnvFeatureFlag{}
			if got := f.CanBeSkipped(tt.args.environment); got != tt.want {
				t.Errorf("CanBeSkipped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvFeatureFlag_IsExplicitlyActive(t *testing.T) {
	type fields struct {
		env environment.Environment
	}
	type args struct {
		flag string
	}
	mockEnv := mocks.NewEnvironment(t)
	mockEnv.EXPECT().Get("valid_flag").Return("1").Once()
	mockEnv.EXPECT().Get("valid_flag").Return("true").Once()
	mockEnv.EXPECT().Get("valid_flag").Return("false").Once()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: fields{env: mockEnv},
			args:   args{flag: "valid_flag"},
			want:   true,
		},
		{
			name:   "true",
			fields: fields{env: mockEnv},
			args:   args{flag: "valid_flag"},
			want:   true,
		},
		{
			name:   "false",
			fields: fields{env: mockEnv},
			args:   args{flag: "valid_flag"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := OsEnvFeatureFlag{
				env: tt.fields.env,
			}
			if got := f.IsExplicitlyActive(tt.args.flag); got != tt.want {
				t.Errorf("IsExplicitlyActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvFeatureFlag_IsActive(t *testing.T) {
	type fields struct {
		env environment.Environment
	}
	type args struct {
		flag string
	}
	flag := "valid_flag"
	mockEnv := mocks.NewEnvironment(t)
	mockEnv.EXPECT().Get("ENVIRONMENT").Return("local").Once()
	mockEnv.EXPECT().Get("ENVIRONMENT").Return("staging").Once()
	mockEnv.EXPECT().Get(flag).Return("false").Once()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "skipped",
			fields: fields{env: mockEnv},
			args:   args{flag: flag},
			want:   true,
		},
		{
			name:   "checked",
			fields: fields{env: mockEnv},
			args:   args{flag: flag},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := OsEnvFeatureFlag{
				env: tt.fields.env,
			}
			if got := f.IsActive(tt.args.flag); got != tt.want {
				t.Errorf("IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}
