package aftership

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

var (
	ErrUnexpectedResponseStatus = errors.New("unexpected response status")
)

type AfterShip struct {
	key string
}

func New(key string) *AfterShip {
	return &AfterShip{
		key: key,
	}
}

func (a *AfterShip) prepareAndSend(ctx context.Context, method, url string, body interface{}) (*http.Response, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, "https://api.aftership.com/v4"+url, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("aftership-api-key", a.key)
	req.Header.Set("Content-Type", "application/json")

	b, err := httputil.DumpRequest(req, true)
	fmt.Println(string(b))

	return http.DefaultClient.Do(req)
}

func formatError(source error, r *http.Response) error {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		var body struct {
			Meta struct {
				Code    int
				Message string
				Type    string
			}
		}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return err
		}

		return fmt.Errorf("%w: %s: %s", source, body.Meta.Type, body.Meta.Message)
	} else {
		return fmt.Errorf("%w: %s", source, r.Status)
	}
}
