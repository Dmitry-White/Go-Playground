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

	log.Printf("stopping container %s", cid)

	err = exec.Command("docker", "rm", "-f", cid).Run()
	if err != nil {
		return fmt.Errorf("failed to stop %s: %w", cid, err)
	}

	return nil
}

func setupErrors(fileName string) string {
	fmt.Println("Not Implemented")
	return ""
}

func checkErrors(fileName string) error {
	err := killContainer(fileName)
	if err != nil {
		return err
	}

	return nil
}
