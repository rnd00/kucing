package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const thiscatdoesnotexist string = "https://thiscatdoesnotexist.com"

func Cat() ([]byte, error) {
	var data []byte
	resp, err := http.Get(thiscatdoesnotexist)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return data, err
		}
		data = bodyBytes
		return data, nil
	}
	return []byte{}, errors.New(fmt.Sprintf("HTTP Status Code returned was not expected; %+v", resp.StatusCode))
}
