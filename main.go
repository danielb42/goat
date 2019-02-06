package goat

import (
	"errors"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// AddJob adds a command to the at(1) execution queue. It will be run at atTime.
func AddJob(command string, atTime time.Time) (int, error) {
	atCmd := exec.Command("at", "-t", atTime.Format("200601021504"))
	atStdin, err := atCmd.StdinPipe()
	if err != nil {
		return -1, errors.New("could not talk to at")
	}

	io.WriteString(atStdin, command)
	atStdin.Close()

	atStdout, err := atCmd.CombinedOutput()
	if err != nil {
		return -1, errors.New("could not add job")
	}

	outputTokens := regexp.MustCompile(`job (\d+) at`).FindStringSubmatch(string(atStdout))
	if outputTokens == nil {
		return -1, errors.New("job has been added, but did not get job id")
	}

	jobID, _ := strconv.Atoi(outputTokens[1])

	return jobID, nil
}

// RemoveJob removes the job specified by jobID from the at(1) execution queue.
func RemoveJob(jobID int) error {
	if err := exec.Command("atrm", strconv.Itoa(jobID)).Run(); err != nil {
		return errors.New("could not delete job")
	}

	return nil
}
