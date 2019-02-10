package process

import (
	"bytes"
	"fmt"
	"github.com/mono83/charlie/model"
	"github.com/mono83/charlie/db"
)

// ReleaseConsumer accepts Release and performs implementation specific action on it
type ReleaseConsumer func(release model.Release) error

// PrintToConsole is an example of ReleaseConsumer that prints information about release into console
func PrintToConsole(release model.Release) error {

	var buffer bytes.Buffer

	buffer.WriteString("Release - ")
	buffer.WriteString(release.Version.String())
	buffer.WriteString(" has ")

	size := len(release.SummaryType())
	counter := 0
	for t, n := range release.SummaryType() {
		buffer.WriteString(fmt.Sprintf("%d", n))
		buffer.WriteString(" issues of type ")
		buffer.WriteString(t.String())
		if counter < size-1 {
			buffer.WriteString(", ")
		}
		counter++
	}

	fmt.Println(buffer.String())

	return nil
}

func DbSaver(releaseRepo db.ReleaseRepository) (func(release model.Release) error) {
	return func(release model.Release) error {
		_, err := releaseRepo.Store(release)
		return err
	}
}