package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetLatestTag(prefix string) (string, error) {
	out, err := exec.Command("git", "tag", "--list", prefix+"*", "--sort=-creatordate").Output()
	if err != nil {
		return "", fmt.Errorf("erro ao obter tags: %w", err)
	}
	tags := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(tags) == 0 || tags[0] == "" {
		return "", nil
	}
	return tags[0], nil
}

func GetCommitsSince(tag string) ([]string, error) {
	var cmd *exec.Cmd
	if tag == "" {
		cmd = exec.Command("git", "log", "HEAD", "--pretty=%s")
	} else {
		cmd = exec.Command("git", "log", fmt.Sprintf("%s..HEAD", tag), "--pretty=%s")
	}

	out, err := cmd.Output()

	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			stderr := string(ee.Stderr)
			if strings.Contains(stderr, "bad revision") ||
				strings.Contains(stderr, "unknown revision") ||
				strings.Contains(stderr, "not a git repository") {
				return []string{}, nil
			}
		}
		return nil, fmt.Errorf("erro ao obter commits: %w", err)
	}

	output := strings.TrimSpace(string(out))
	if output == "" {
		return []string{}, nil
	}

	lines := strings.Split(output, "\n")
	return lines, nil
}
