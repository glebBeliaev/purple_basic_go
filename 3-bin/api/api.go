package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"purple_basic_go/3-bin/model"
)

type Client struct {
	http    *http.Client
	baseURL string
	key     string
}

func NewClient() *Client {
	base := os.Getenv("JSONBIN_BASE")
	if base == "" {
		base = "https://api.jsonbin.io/v3"
	}
	key := os.Getenv("JSONBIN_KEY")
	return &Client{
		http:    &http.Client{Timeout: 15 * time.Second},
		baseURL: base,
		key:     key,
	}
}

func (c *Client) CreateBin(ctx context.Context, binData interface{}, name string) (string, error) {
	if c.key == "" {
		return "", errors.New("JSONBIN_KEY not set")
	}
	url := c.baseURL + "/b"

	body := map[string]interface{}{
		"metadata": map[string]interface{}{"name": name},
		"record":   binData,
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.key)

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("create failed: %s: %s", resp.Status, string(b))
	}

	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}

	// jsonbin v3 обычно возвращает id в metadata.id
	if md, ok := out["metadata"].(map[string]any); ok {
		if id, ok := md["id"].(string); ok && id != "" {
			return id, nil
		}
		if id, ok := md["uuid"].(string); ok && id != "" {
			return id, nil
		}
	}
	if id, ok := out["id"].(string); ok && id != "" {
		return id, nil
	}

	// Fallback: вернём «сырое» тело (чтобы не потерять инфу)
	rawOut, _ := json.Marshal(out)
	return string(rawOut), nil
}

func (c *Client) GetBin(ctx context.Context, id string) (*model.Bin, error) {
	if c.key == "" {
		return nil, errors.New("JSONBIN_KEY not set")
	}
	url := c.baseURL + "/b/" + id + "/latest"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", c.key)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("get failed: %s: %s", resp.Status, string(b))
	}

	var wrapper map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}

	if recRaw, ok := wrapper["record"]; ok {
		var b model.Bin
		if err := json.Unmarshal(recRaw, &b); err != nil {
			return nil, err
		}
		return &b, nil
	}

	return nil, errors.New("unexpected response: no record field")
}

func (c *Client) UpdateBin(ctx context.Context, id string, binData interface{}) error {
	if c.key == "" {
		return errors.New("JSONBIN_KEY not set")
	}
	url := c.baseURL + "/b/" + id

	raw, err := json.Marshal(binData)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(raw))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.key)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("update failed: %s: %s", resp.Status, string(b))
	}
	return nil
}

func (c *Client) DeleteBin(ctx context.Context, id string) error {
	if c.key == "" {
		return errors.New("JSONBIN_KEY not set")
	}
	url := c.baseURL + "/b/" + id

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", c.key)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete failed: %s: %s", resp.Status, string(b))
	}
	return nil
}
