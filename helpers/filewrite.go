package helpers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WriteToFile(data []byte) (string, error) {
	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fnFound := false
	var fn, fp string
	for !fnFound {
		// generate filename, 10 in length
		fn = generateFilename(10)
		// check
		fp = fmt.Sprintf("%s/%s", wd, fn)
		if err := checkFilename(fp); err != nil {
			// go search another filename
			continue
		}
		fnFound = true
	}

	if err := os.WriteFile(fp, data, 0644); err != nil {
		return "", err
	}

	return fp, nil
}

func generateFilename(l int) string {
	// use time(unixnano as seed)
	var sr *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate string using predetermined base for name
	const cs = "aiueo24680"
	b := make([]byte, l)
	for i := range b {
		b[i] = cs[sr.Intn(len(cs))]
	}

	// return string
	return fmt.Sprintf("kucing[%s].jpg", string(b))
}

func checkFilename(x string) error {
	a, _ := os.Stat(x)
	if a != nil {
		return errors.New("This filename has already existed")
	}
	return nil
}
