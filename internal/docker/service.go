package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type DockerService struct {
	basePath string
}

type DockerOrganization struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func NewService(basePath string) *DockerService {
	return &DockerService{basePath: basePath}
}

func (s *DockerService) Start(org DockerOrganization) error {
	return s.runComposeCommand(org, "up", "-d")
}

func (s *DockerService) Stop(org DockerOrganization) error {
	return s.runComposeCommand(org, "down")
}

func (s *DockerService) Restart(org DockerOrganization, services ...string) error {
	args := []string{"restart"}
	if len(services) > 0 && services[0] != "" && services[0] != "all" {
		args = append(args, services[0])
	}
	return s.runComposeCommand(org, args...)
}

func (s *DockerService) Logs(org DockerOrganization, service string, tail int) (string, error) {
	composeFile, err := s.getComposeFilePath(org)
	if err != nil {
		return "", err
	}

	args := []string{"-f", composeFile, "logs", fmt.Sprintf("--tail=%d", tail), service}
	output, err := s.runCommand(filepath.Dir(composeFile), "docker-compose", args...)
	if err != nil {
		return "", fmt.Errorf("failed to fetch logs for service %q: %w\n%s", service, err, output)
	}
	return output, nil
}

func (s *DockerService) Status(org DockerOrganization) (map[string]string, error) {
	composeFile, err := s.getComposeFilePath(org)
	if err != nil {
		return nil, err
	}

	args := []string{"-f", composeFile, "ps", "--format", "json"}
	output, err := s.runCommand(filepath.Dir(composeFile), "docker-compose", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service status: %w\n%s", err, output)
	}

	status := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		switch {
		case strings.Contains(line, "nextcloud-app"):
			status["nextcloud"] = parseStatus(line)
		case strings.Contains(line, "mariadb"):
			status["database"] = parseStatus(line)
		case strings.Contains(line, "redis"):
			status["redis"] = parseStatus(line)
		}
	}

	return status, nil
}

// Internal Helpers

func (s *DockerService) runComposeCommand(org DockerOrganization, args ...string) error {
	composeFile, err := s.getComposeFilePath(org)
	if err != nil {
		return err
	}

	fullArgs := append([]string{"-f", composeFile}, args...)
	output, err := s.runCommand(filepath.Dir(composeFile), "docker-compose", fullArgs...)
	if err != nil {
		return fmt.Errorf("compose command failed: %w\n%s", err, output)
	}
	return nil
}

func (s *DockerService) getComposeFilePath(org DockerOrganization) (string, error) {
	dir := filepath.Join(s.basePath, org.Name)
	file := filepath.Join(dir, "docker-compose.yml")

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return "", fmt.Errorf("compose file not found for organization %q", org.Name)
	}
	return file, nil
}

func (s *DockerService) runCommand(dir, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func parseStatus(line string) string {
	if strings.Contains(line, "Up") {
		return "running"
	}
	return "stopped"
}
