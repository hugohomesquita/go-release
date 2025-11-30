package version

import (
	"fmt"
	"strconv"
	"strings"
)

func NextVersion(current, bump, tagPrefix string) string {
	// Remove o prefixo da tag atual
	v := strings.TrimPrefix(current, tagPrefix)

	// Se não houver versão atual, começar do zero
	if v == "" {
		switch bump {
		case "major":
			return fmt.Sprintf("%s1.0.0", tagPrefix)
		case "minor":
			return fmt.Sprintf("%s0.1.0", tagPrefix)
		default:
			return fmt.Sprintf("%s0.0.1", tagPrefix)
		}
	}

	parts := strings.Split(v, ".")
	if len(parts) < 3 {
		return fmt.Sprintf("%s0.1.0", tagPrefix)
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch bump {
	case "major":
		major++
		minor, patch = 0, 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	}

	return fmt.Sprintf("%s%d.%d.%d", tagPrefix, major, minor, patch)
}
