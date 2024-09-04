package utils

import (
	"encoding/json"
	"io"
	"os"
	"wordma/config"
	"wordma/server/dto"
)

var EmojiJson []dto.EmojiJson

// InitEmoji 初始化Emoji
func InitEmoji() {
	for _, path := range config.EmojiPaths {
		emojiData, err := loadEmojiJson(path)
		if err != nil {
			panic("载入Emoji错误：" + err.Error())
		}
		EmojiJson = append(EmojiJson, *emojiData)
	}
}

func loadEmojiJson(filePath string) (*dto.EmojiJson, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 使用值类型初始化 emojiData
	var emojiData dto.EmojiJson
	if err := json.Unmarshal(byteValue, &emojiData); err != nil {
		return nil, err
	}

	// 返回指向值类型的指针
	return &emojiData, nil
}
