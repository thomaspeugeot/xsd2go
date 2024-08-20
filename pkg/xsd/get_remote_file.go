package xsd

import (
	"fmt"
	"io"
	"net/http"
)

func getRemoteFile(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)

	}
	return resp.Body, nil
}
