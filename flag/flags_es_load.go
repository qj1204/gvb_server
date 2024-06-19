package flag

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/service/es_service/index"
	"os"
)

func ESLoad(jsonPath string) {
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		logrus.Fatalf("%s err: %s", jsonPath, err.Error())
	}

	var response ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		logrus.Fatalf("%s err: %s", string(byteData), err.Error())
	}

	// 创建索引
	index.CreateIndexByJson(response.Index, response.Mapping)

	// 批量导入数据
	bulk := global.ESClient.Bulk().Index(response.Index).Refresh("true")
	for _, model := range response.Data {

		var mapData map[string]any
		_ = json.Unmarshal(model.Row, &mapData)
		row, _ := json.Marshal(mapData)
		// 插入的数据，不能有换行
		req := elastic.NewBulkCreateRequest().Id(model.ID).Doc(string(row))
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Errorf("数据添加失败 err:%s", err.Error())
		return
	}
	logrus.Infof("数据添加成功， 共添加 %d 条", len(res.Succeeded()))
}
