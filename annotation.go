package renamer

import "strings"

const (
	// AnnotationPackageKey is the key to specify the package
	AnnotationPackageKey = "package"
	// AnnotationPackagePathKey is the key to specify the package path
	AnnotationPackagePathKey = "package_path"
	// AnnotationStructKey is the key to specify the struct name
	AnnotationStructKey = "struct"
)

// FromPackage is an option to specify the package name.
func FromPackage(pkg string) Option {
	return func(id, preferred string, annotations map[string]string) (string, bool) {
		if pkg != "" {
			annotations[AnnotationPackageKey] = pkg
		}
		return "", false
	}
}

// FromPackage is an option to specify the full package path like example.com/user/repo/pkg.StructName
func FromFullPath(path string) Option {
	return func(id, preferred string, annotations map[string]string) (string, bool) {
		if path != "" {
			split := strings.Split(path, ".")

			annotations[AnnotationPackagePathKey] = strings.Join(split[:len(split)-1], ".")

			if len(split) >= 2 {
				annotations[AnnotationStructKey] = split[len(split)-1]
			}
		}
		return "", false
	}
}
