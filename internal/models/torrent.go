package models

// Torrent 结构体定义
type Torrent struct {
	Hash  string  `json:"hash"`
	Name  string  `json:"name"`
	Ratio float64 `json:"ratio"`
	Addon int64   `json:"added_on"`
}
