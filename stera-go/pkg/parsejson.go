package pkg

import (
	"encoding/json"
	"log"
)

// GeoJSON is Convert JSON to Go struct
type GeoJSON struct {
	Type       string `json:"type"`
	Properties struct {
		ID          string      `json:"id"`
		Fid         string      `json:"fid"`
		SeibiData   int         `json:"整備データ"`
		SeibiA      int         `json:"整備デーA"`
		SeibiFinish string      `json:"整備完了日"`
		OrgGILvl    string      `json:"orgGILvl"`
		OrgMDID     interface{} `json:"orgMDId"`
		HyoujiKubun string      `json:"表示区分"`
		KIND        string      `json:"種別"`
		NAME        interface{} `json:"名称"`
		Median      float64     `json:"median標高"`
	} `json:"properties"`
	Geometry struct {
		Type string `json:"type"`
		// Coordinates [][][][]float64 `json:"coordinates"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
}

// ParseJSON はGeoJSONデータの読み込み処理をする．
func ParseJSON(jStr string) (id string, elv float64, cordnts [][][]float64) {

	// GeoJson構造体の変数stcDataを宣言
	var stcData GeoJSON

	// エラー処理のためjsonStrを[]byte型に変換？
	// Unmarshalで[]byte型で受け取ったJSON形式のファイルをポインタに保存
	if err := json.Unmarshal([]byte(jStr), &stcData); err != nil {
		log.Println(err)
		// return nil
	}

	// pj := &stcData
	// log.Printf("pointerJSON:%p\n", pj)
	// log.Println("stcData", unsafe.Sizeof(stcData))

	id = stcData.Properties.ID
	elv = stcData.Properties.Median
	// nullの場合は0になる
	cordnts = stcData.Geometry.Coordinates

	return id, elv, cordnts
}
