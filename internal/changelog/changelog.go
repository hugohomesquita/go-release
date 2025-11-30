package changelog

import (
	"fmt"
	"go-release/internal/config"
	"strings"
)

func Generate(commits []string, bump config.Bump) string {
	var feats, fixes, breaking []string

	for _, c := range commits {
		isMajor := false
		for _, kw := range bump.MajorKeywords {
			if strings.Contains(c, kw) {
				breaking = append(breaking, c)
				isMajor = true
				break
			}
		}
		if isMajor {
			continue
		}

		isMinor := false
		for _, kw := range bump.MinorKeywords {
			if strings.Contains(c, kw) {
				feats = append(feats, c)
				isMinor = true
				break
			}
		}
		if isMinor {
			continue
		}

		for _, kw := range bump.PatchKeywords {
			if strings.Contains(c, kw) {
				fixes = append(fixes, c)
				break
			}
		}
	}

	return fmt.Sprintf(`# Changelog

## 🚀 Features
%s

## 🐛 Fixes
%s

## 💥 Breaking Changes
%s
`,
		formatList(feats),
		formatList(fixes),
		formatList(breaking),
	)
}

func formatList(items []string) string {
	if len(items) == 0 {
		return "- Nenhum"
	}
	return "- " + strings.Join(items, "\n- ")
}
