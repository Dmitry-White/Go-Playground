package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func isValidID(cid string) bool {
	return len(cid) == 12 || len(cid) == 64
}

func killContainer(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	defer os.Remove(fileName)

	cid := strings.TrimSpace(string(data))
	if !isValidID(cid) {
		return fmt.Errorf("bad container id: %q", cid)
	}

	log.Printf("Stopping container %s\n", cid)

	command := exec.Command("docker", "rm", "-f", cid)
	err = command.Run()
	if err != nil {
		return fmt.Errorf("failed to stop %s: %w", cid, err)
	}

	return nil
}

func runContainer(filename string) error {
	image := "docker/getting-started"
	ports := "80:80"

	log.Printf("Running container from \"%s\"\n", image)

	command := exec.Command("docker", "run", "--cidfile", filename, "-d", "-p", ports, image)
	err := command.Run()
	if err != nil {
		return fmt.Errorf("failed to run %s: %w", image, err)
	}

	return nil
}

func setupErrors(fileName string) error {
	err := runContainer(fileName)
	if err != nil {
		return err
	}

	return nil
}

func checkErrors(fileName string) error {
	err := killContainer(fileName)
	if err != nil {
		return err
	}

	return nil
}
