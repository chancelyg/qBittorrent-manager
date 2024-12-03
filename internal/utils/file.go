package utils

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"torrent-manager/internal/models"
)

// SearchTorrentOnLocal 根据哈希值在本地文件中查找种子
func SearchTorrentOnLocal(hash, recordFile string) (models.Torrent, error) {
	var torrent models.Torrent
	var torrents []models.Torrent

	if _, err := os.Stat(recordFile); os.IsNotExist(err) {
		return torrent, errors.New("no existing torrents found")
	}

	data, err := os.ReadFile(recordFile)
	if err != nil {
		return torrent, err
	}

	if err := json.Unmarshal(data, &torrents); err != nil {
		return torrent, err
	}

	for _, torrent := range torrents {
		if torrent.Hash == hash {
			return torrent, nil
		}
	}

	return torrent, errors.New("hash not found")
}

// SaveTorrentToLocal 将种子记录保存到本地文件
func SaveTorrentToLocal(torrents []models.Torrent, recordFile string) error {
	data, err := json.MarshalIndent(torrents, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(recordFile, data, 0644); err != nil {
		return err
	}

	return nil
}

// IsWithinProtectionPeriod 检查是否在保护期内
func IsWithinProtectionPeriod(addTime time.Time, protectionPeriod int) bool {
	nowTime := time.Now()
	daysDiff := int(nowTime.Sub(addTime).Hours() / 24)
	return daysDiff < protectionPeriod
}
