package dictionary

import "github.com/maxzhirnov/utm-builder/internal/config"

// Sources returns the list of available UTM sources
func Sources() []string {
	config.Mu.RLock()
	defer config.Mu.RUnlock()
	return config.Config.Sources
}

// Mediums returns the list of available UTM mediums
func Mediums() []string {
	config.Mu.RLock()
	defer config.Mu.RUnlock()
	return config.Config.Mediums
}
