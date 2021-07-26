package renamer

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// Option is a renamer option.
type Option func(id, preferred string, annotations map[string]string) (string, bool)

// WithPackageName adds the package name to the start of the name.
var WithPackageName = func(id, preferred string, annotations map[string]string) (string, bool) {
	pkg, ok := annotations[AnnotationPackageKey]

	if !ok {
		path, ok := annotations[AnnotationPackagePathKey]

		if !ok {
			return "", false
		}

		split := strings.Split(path, "/")

		if len(split) == 0 {
			return "", false
		}

		pkg = split[len(split)-1]
	}

	return strcase.ToCamel(pkg) + strcase.ToCamel(preferred), true
}
