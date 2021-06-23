//
// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maven

import (
	"context"
	"crypto/sha256"
	hex2 "encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Artifact struct {
	root     string
	fileName string
}

// this is useless since md5 broken?
func (a *Artifact) MD5(ctx context.Context) (string, error) {
	mu := a.root + "/" + a.fileName + ".md5"

	mb := make([]byte, 0, 32)

	r, err := retrieve(ctx, mu)
	if err != nil {
		return "", err
	}

	defer r.Close()
	if b, err := io.ReadAll(io.LimitReader(r, int64(cap(mb)))); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

// this is useless since sha1 broken?
func (a *Artifact) SHA1(ctx context.Context) (string, error) {
	su := a.root + "/" + a.fileName + ".sha1"

	sb := make([]byte, 0, 40)

	r, err := retrieve(ctx, su)
	if err != nil {
		return "", err
	}

	defer r.Close()
	if b, err := io.ReadAll(io.LimitReader(r, int64(cap(sb)))); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func (a *Artifact) CalculateSHA256(ctx context.Context) (string, error) {
	r, err := retrieve(ctx, a.root+"/"+a.fileName)
	if err != nil {
		return "", err
	}

	defer r.Close()
	sh := sha256.New()
	if _, err := io.Copy(sh, r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sb := sh.Sum(make([]byte, 0, sh.Size()))
	hex := hex2.EncodeToString(sb)

	return hex, nil
}

func retrieve(ctx context.Context, url string) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("could not retrieve remote artifact -- status code:%v", resp.StatusCode)
	}
	return resp.Body, nil
}
