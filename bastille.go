package main

import (
	"fmt"
	"log"
	"os/exec"
)

func runBastilleCommands(args ...string) (string, error) {
	cmd := exec.Command("bastille", args...)
	log.Println("runBastilleCommands:", cmd)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bastille: %s ,failed: %v\n %s", cmd, err, result)
	}

	// result := "ok"
	return string(result), nil
}

func runOsCommands(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	log.Println("runOsCommands:", cmd)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("os: %v ,failed: %v\n %s", args, err, result)
	}

	return string(result), nil
}

func osInfo() (string, error) {
	args := []string{"-a"}
	return runOsCommands("uname", args...)
}

func osMemInfo() (string, error) {
	args := []string{"-h", "hw.physmem", "hw.usermem"}
	return runOsCommands("sysctl", args...)
}

func bastilleListAllJson() (string, error) {
	args := []string{"list", "-j", "all"}
	return runBastilleCommands(args...)
}

func bastilleVersion() (string, error) {
	args := []string{"-v"}
	return runBastilleCommands(args...)
}

func bastilleBootstrap(options, releasetemplate, updatearch string) (string, error) {
	args := []string{"bootstrap"}
	if options != "" {
		args = append(args, options)
	}

	if releasetemplate != "" {
		args = append(args, releasetemplate)
	}

	if updatearch != "" {
		args = append(args, updatearch)
	}

	return runBastilleCommands(args...)
}
