package utils

type FeatureFlag interface {
	CanBeSkipped(environment string) bool
	IsActive(flag string) bool
	IsExplicitlyActive(flag string) bool
}
