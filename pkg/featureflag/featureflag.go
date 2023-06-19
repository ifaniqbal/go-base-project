package featureflag

import (
	"github.com/ifaniqbal/go-base-project/pkg/environment"
)

// OsEnvFeatureFlag is a struct for checking feature flag status from
// OS environment variables.
type OsEnvFeatureFlag struct {
	env environment.Environment
}

func NewOsEnv(env environment.Environment) *OsEnvFeatureFlag {
	return &OsEnvFeatureFlag{
		env: env,
	}
}

func Default() *OsEnvFeatureFlag {
	return NewOsEnv(environment.Default())
}

// CanBeSkipped checks if a feature flag can be skipped in a specific environment.
// It returns true if the environment is either "development" or "local".
func (f OsEnvFeatureFlag) CanBeSkipped(environment string) bool {
	return environment == "development" || environment == "local"
}

// IsActive checks if a feature flag is active in the current environment.
// If the environment is "development" or "local", it returns true.
// Otherwise, it calls IsExplicitlyActive to check the flag's status.
func (f OsEnvFeatureFlag) IsActive(flag string) bool {
	if f.CanBeSkipped(f.env.Get("ENVIRONMENT")) {
		return true
	}

	return f.IsExplicitlyActive(flag)
}

// IsExplicitlyActive checks if a feature flag is explicitly active.
// It returns true if the value of the flag is either "1" or "true".
func (f OsEnvFeatureFlag) IsExplicitlyActive(flag string) bool {
	val := f.env.Get(flag)
	return val == "1" || val == "true"
}
