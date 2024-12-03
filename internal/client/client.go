package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"torrent-manager/internal/models"
)

// Client 结构体定义，包含 HTTP 客户端和 qBittorrent URL
type Client struct {
	httpClient *http.Client
	qbURL      string
}

// NewClient 创建一个新的 Client 实例
func NewClient(qbURL string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		httpClient: &http.Client{Jar: jar},
		qbURL:      qbURL,
	}, nil
}

// 登录并获取会话 cookie
func (c *Client) Login(username, password string) error {
	loginData := fmt.Sprintf("username=%s&password=%s", username, password)
	req, err := http.NewRequest("POST", c.qbURL+"/api/v2/auth/login", bytes.NewBufferString(loginData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to login, status code: %d", resp.StatusCode)
	}

	fmt.Println("Logged in successfully.")
	return nil
}

// 获取当前所有种子的状态
func (c *Client) GetTorrents() ([]models.Torrent, error) {
	req, err := http.NewRequest("GET", c.qbURL+"/api/v2/torrents/info", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch torrents info, status code: %d", resp.StatusCode)
	}

	var torrents []models.Torrent
	if err := json.NewDecoder(resp.Body).Decode(&torrents); err != nil {
		return nil, err
	}

	fmt.Println("Torrents information fetched successfully.")
	return torrents, nil
}

// 删除种子
func (c *Client) DeleteTorrent(hash string) error {
	data := fmt.Sprintf("hashes=%s", hash)
	req, err := http.NewRequest("POST", c.qbURL+"/api/v2/torrents/delete", bytes.NewBufferString(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete torrent, status code: %d", resp.StatusCode)
	}

	return nil
}
