package path_provider

import (
	"fmt"
	"path/filepath"
)

// PathProvider - a structure that should be used to get the path to a file/folder in the project
// from which projectlinter was actually launched
//
// For example, I have a project auth-sv, and I run projectlinter in it
// pathProvider.PathInCaller(".helm/templates") will return me the absolute path to the templates folder in the auth-sv project
// For example, this would be /var/www/auth-sv/.helm/.templates
type PathProvider struct {
	customServicePath string
}

func NewPathProvider(pathToService string) *PathProvider {
	return &PathProvider{pathToService}
}

func (p *PathProvider) PathToCaller() string {
	if p.customServicePath != "" {
		return p.customServicePath
	}

	currentPath, _ := filepath.Abs(".")

	return currentPath
}

func (p *PathProvider) PathInCaller(relativePath string) string {
	return fmt.Sprintf("%s/%s", p.PathToCaller(), relativePath)
}
