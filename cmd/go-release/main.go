package main

import (
	"fmt"
	"go-release/internal/changelog"
	"go-release/internal/config"
	"go-release/internal/git"
	"go-release/internal/version"
	"os"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig("go-release.yaml")
	if err != nil {
		panic(err)
	}

	for _, project := range cfg.Projects {
		lastTag, err := git.GetLatestTag(project.TagPrefix)
		if err != nil {
			panic(err)
		}

		if lastTag == "" {
			fmt.Println("Nenhuma tag encontrada, iniciando do zero.")
		} else {
			fmt.Println("Última versão:", lastTag)
		}

		commitsAfterTag, err := git.GetCommitsSince(lastTag)
		if err != nil {
			panic(err)
		}
		fmt.Println("Commits encontrados:", len(commitsAfterTag))

		bumpType := detectBump(commitsAfterTag, project)
		newVersion := version.NextVersion(lastTag, bumpType, project.TagPrefix)
		fmt.Println("Nova versão:", newVersion)

		ch := changelog.Generate(commitsAfterTag, project.Bump, newVersion)
		if err := os.WriteFile("CHANGELOG.md", []byte(ch), 0644); err != nil {
			panic(err)
		}

		fmt.Printf("✅ Release %s → %s gerado com sucesso!\n", lastTag, newVersion)
	}
}

func detectBump(commits []string, project config.Project) string {
	for _, c := range commits {
		for _, kw := range project.Bump.MajorKeywords {
			if containsWord(c, kw) {
				return "major"
			}
		}
	}
	for _, c := range commits {
		for _, kw := range project.Bump.MinorKeywords {
			if containsWord(c, kw) {
				return "minor"
			}
		}
	}
	for _, c := range commits {
		for _, kw := range project.Bump.PatchKeywords {
			if containsWord(c, kw) {
				return "patch"
			}
		}
	}
	return "patch"
}

func containsWord(text, word string) bool {
	return strings.Contains(text, word)
}
