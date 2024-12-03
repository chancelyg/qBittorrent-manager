package main

import (
	"flag"
	"fmt"
	"time"

	"torrent-manager/internal/client"
	"torrent-manager/internal/utils"
)

func main() {
	// 定义命令行参数
	qbURL := flag.String("url", "http://localhost:8080", "The URL of the qBittorrent Web UI")
	username := flag.String("username", "admin", "The username for qBittorrent Web UI")
	password := flag.String("password", "adminadmin", "The password for qBittorrent Web UI")
	recordFile := flag.String("recordFile", "torrent-records.json", "The file to save torrent records")
	ratioIncrease := flag.Float64("ratioIncrease", 0.5, "The ratio increases each time")
	protectionPeriod := flag.Int("protectionPeriod", 7, "The protection period for torrent")
	try := flag.Bool("try", false, "Display the deletion target without actually deleting the torrent")

	// 解析命令行参数
	flag.Parse()

	// 创建 qBittorrent 客户端
	qbClient, err := client.NewClient(*qbURL)
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	// 执行登录
	if err := qbClient.Login(*username, *password); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 获取种子信息
	torrents, err := qbClient.GetTorrents()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 检查上传率
	for _, torrent := range torrents {
		fmt.Printf("Processing torrent <%s>\n", torrent.Name)
		if record, err := utils.SearchTorrentOnLocal(torrent.Hash, *recordFile); err != nil {
			fmt.Printf("Torrent <%s> name not found\n", torrent.Name)
			continue
		} else {
			addTime := time.Unix(torrent.Addon, 0)
			if utils.IsWithinProtectionPeriod(addTime, *protectionPeriod) {
				fmt.Printf("Torrent <%s> not more than %d days\n", torrent.Name, *protectionPeriod)
				continue
			}

			fmt.Printf("Torrent <%s> local record ratio %.2f and torrent current ratio %.2f\n", record.Name, record.Ratio, torrent.Ratio)

			if torrent.Ratio-record.Ratio < *ratioIncrease {
				fmt.Printf("Deleting torrent <%s> due to insufficient share ratio increase (ratio increase %.2f).\n", torrent.Name, torrent.Ratio-record.Ratio)
				if err := qbClient.DeleteTorrent(torrent.Hash); err != nil && !*try {
					fmt.Printf("Failed to delete torrent %s error: %v\n", torrent.Name, err)
				}
			}
		}
	}

	// 保存本次查询结果，方便下次查询比对分享率的增长情况
	if err := utils.SaveTorrentToLocal(torrents, *recordFile); err != nil {
		fmt.Printf("Error saving torrents to local file: %v\n", err)
		return
	}

	fmt.Println("All records processed and updated.")
}
