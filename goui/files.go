package goui

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

//var resLock

type resource struct {
	err      error
	lastUsed time.Time
}

func init() {
	//go func() {
	//	for {
	//		select {
	//
	//		}
	//	}
	//}()
}

var loadedResources = map[string]resource{}

func loadFile(path string) (io.ReadCloser, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}

	file, err := os.Open(path)
	return file, err
}
