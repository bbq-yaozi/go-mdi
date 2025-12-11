package generate

import (
	"encoding/json"
	"path/filepath"

	"github.com/bbq-yaozi/go-mdi/internal/assets"
)

var metaFilePath = "mdi/meta.json"
var svgPath = "mdi/svg"

type MetaInfo struct {
	ID         string   `json:"id"`         // 唯一标识符
	BaseIconID string   `json:"baseIconId"` // 基础图标ID
	Name       string   `json:"name"`       // 图标名称
	Codepoint  string   `json:"codepoint"`  // Unicode码点
	Aliases    []string `json:"aliases"`    // 别名列表
	Tags       []string `json:"tags"`       // 分类标签
	Styles     []string `json:"styles"`     // 可用样式
	Author     string   `json:"author"`     // 作者
	Version    string   `json:"version"`    // 版本号
	Deprecated bool     `json:"deprecated"` // 是否弃用

	Data []byte `json:"-"`
}

func LoadMetaList() ([]*MetaInfo, error) {
	metaData, err := assets.ReadFile(metaFilePath)
	if err != nil {
		return nil, err
	}

	var metaList []*MetaInfo
	err = json.Unmarshal(metaData, &metaList)
	if err != nil {
		panic(err)
	}

	for _, meta := range metaList {
		svgData, err := loadSvgFile(meta.Name)
		if err != nil {
			return nil, err
		}

		meta.Data = svgData
	}
	return metaList, nil
}

func loadSvgFile(name string) ([]byte, error) {
	filePath := filepath.Join(svgPath, name+".svg")

	svgData, err := assets.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return svgData, nil
}
