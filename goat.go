package goat

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// AddJob adds a command to an at(1) execution queue. It will be run at atTime.
// Optionally, the job can be written to specific queue toQueue, default is "a".
// Queues must be adressed by a single letter.
func AddJob(command string, atTime time.Time, toQueue ...string) (int, error) {
	atCmd := exec.Command("at", "-t", atTime.Format("200601021504"))

	if len(toQueue) > 0 {
		atCmd.Args = append(atCmd.Args, "-q", fmt.Sprintf("%c", toQueue[0][0]))
	}

	atStdin, _ := atCmd.StdinPipe()
	_, _ = io.WriteString(atStdin, command)
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

// RemoveJob removes the job specified by jobID from the at(1) execution queues.
func RemoveJob(jobID int) error {
	if err := exec.Command("atrm", strconv.Itoa(jobID)).Run(); err != nil {
		return errors.New("could not delete job")
	}

	return nil
}

// ClearQueue removes all jobs from the at queue specified by queueLetter.
func ClearQueue(queueLetter ...string) error {
	if len(queueLetter) == 0 {
		return errors.New("no queue letter given")
	}

	atqCmd := exec.Command("atq", "-q", fmt.Sprintf("%c", queueLetter[0][0]))
	atqStdout, err := atqCmd.Output()
	if err != nil {
		return errors.New("could not get job IDs")
	}

	jobIDs := regexp.MustCompile(`(?m:^\d+)`).FindAllStringSubmatch(string(atqStdout), -1)

	for _, idToken := range jobIDs {
		jobID, _ := strconv.Atoi(idToken[0])

		if err := RemoveJob(jobID); err != nil {
			return err
		}
	}

	return nil
}
