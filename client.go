/*
Package amedas は気象庁のアメダスのデータの取得を容易にするためのクライアントライブラリです。
*/
package amedas

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"time"
)

// DefaultEndpoint はデータ取得 URL の共通部分を示します。テスト用に変更したい場合は SetEndpoint メソッドで変更できます。
const DefaultEndpoint = "https://www.jma.go.jp/bosai/amedas"

// Client はアメダスのデータ取得クライアントを提供します。
type Client struct {
	endpoint   string
	httpClient *http.Client
	logger     *log.Logger
}

// NewDefaultClient はデータ取得クライアントを構築します。デフォルトの HTTP クライアントを使い、デバッグログは出力しないように構成します。
func NewDefaultClient() *Client {
	return NewClient(http.DefaultClient, log.New(io.Discard, "", log.LstdFlags))
}

// NewClient はデータ取得クライアントを構築します。
func NewClient(httpClient *http.Client, logger *log.Logger) *Client {
	return &Client{endpoint: DefaultEndpoint, httpClient: httpClient, logger: logger}
}

// SetEndpoint はデータ取得 URL の共通部分を指定されたものに変更します。
// 先頭は http://, https:// から始まる必要があります。末尾は / を取り除いたものを指定します。
func (c *Client) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	c.logger.Printf("[go-amedas] %s %s", http.MethodGet, path)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer c.safeClose(resp.Body, filepath.Base(path)+" response body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %s: %s", resp.Status, string(body))
	}
	return body, nil
}

func (c *Client) safeClose(closer io.Closer, name string) {
	if closer != nil {
		if err := closer.Close(); err != nil {
			c.logger.Printf("[go-amedas] failed to close %s: %v", name, err)
		}
	}
}

// LatestTime は最終データ時刻を取得します。
func (c *Client) LatestTime(ctx context.Context) (*time.Time, error) {
	body, err := c.get(ctx, c.endpoint+"/data/latest_time.txt")
	if err != nil {
		return nil, err
	}
	t, err := time.Parse(time.RFC3339, string(body))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// MapPoints は観測地点一覧を ID をキーとする map で取得します。
func (c *Client) MapPoints(ctx context.Context) (map[string]Point, error) {
	body, err := c.get(ctx, c.endpoint+"/const/amedastable.json")
	if err != nil {
		return nil, err
	}
	var m map[string]Point
	if err = json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// ListPoints は観測地点一覧をスライスで取得します。各要素は観測地点 ID 順でソートされます。
func (c *Client) ListPoints(ctx context.Context) ([]PointWithID, error) {
	m, err := c.MapPoints(ctx)
	if err != nil {
		return nil, err
	}
	s := make([]PointWithID, 0, len(m))
	for id, point := range m {
		s = append(s, PointWithID{id, point})
	}
	sort.Slice(s, func(i, j int) bool { return s[i].ID < s[j].ID })
	return s, nil
}

// AllPointMeasurement は全観測地点の観測データを取得します。
func (c *Client) AllPointMeasurement(ctx context.Context, target time.Time) (map[string]Measurement, error) {
	body, err := c.get(ctx, fmt.Sprintf("%s/data/map/%d%02d%02d%02d%02d00.json",
		c.endpoint, target.Year(), target.Month(), target.Day(), target.Hour(), target.Minute()))
	if err != nil {
		return nil, err
	}
	var m map[string]Measurement
	if err = json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// LatestAllPointMeasurement は全観測地点の最新の観測データを取得します。
func (c *Client) LatestAllPointMeasurement(ctx context.Context) (map[string]Measurement, error) {
	latestTime, err := c.LatestTime(ctx)
	if err != nil {
		return nil, err
	}
	return c.AllPointMeasurement(ctx, *latestTime)
}

// SinglePointMeasurements は指定した観測地点の3時間刻みの時間枠内の10分ごとの観測データを取得します。
func (c *Client) SinglePointMeasurements(ctx context.Context, point string, target time.Time) (map[string]Measurement, error) {
	targetHour := int(target.Hour()/3) * 3 // 1時間刻みを3時間刻みに変更
	body, err := c.get(ctx, fmt.Sprintf("%s/data/point/%s/%d%02d%02d_%02d.json",
		c.endpoint, point, target.Year(), target.Month(), target.Day(), targetHour))
	if err != nil {
		return nil, err
	}
	var m map[string]Measurement
	if err = json.Unmarshal(body, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// LatestSinglePointMeasurement　は指定した観測地点の最新の観測データを取得します。
func (c *Client) LatestSinglePointMeasurement(ctx context.Context, point string) (*Measurement, error) {
	latestTime, err := c.LatestTime(ctx)
	if err != nil {
		return nil, err
	}
	m, err := c.SinglePointMeasurements(ctx, point, *latestTime)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	latest := m[keys[len(keys)-1]]
	return &latest, nil
}
