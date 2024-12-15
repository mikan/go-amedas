package amedas

// Point は観測地点の属性を示します。
type Point struct {
	Type   string    `json:"type"`   // A=管区・地方・沖縄気象台, B=測候所・特別地域気象観測所, C=アメダス, D=父島, E=南鳥島, F=富士山
	Elems  string    `json:"elems"`  // [0]=気, [1]=降水量, [2]=風向, [3]=風速, [4]=日照時間, [5]=積雪深, [6]=湿度, [7]=気圧
	Lat    []float64 `json:"lat"`    // 緯度 ([度,分])
	Lon    []float64 `json:"lon"`    // 経度 ([度,分])
	Alt    int       `json:"alt"`    // 標高 (m)
	KjName string    `json:"kjName"` // 漢字
	KnName string    `json:"knName"` // カタカナ
	EnName string    `json:"enName"` // English
}

// PointWithID は観測地点の属性に観測地点 ID を追加した本ライブラリ独自のデータ構造を示します。
type PointWithID struct {
	ID string `json:"id,omitempty"`
	Point
}

// Latitude はこの観測地点の緯度の十進表現を返却します。
func (p *Point) Latitude() float64 {
	return p.Lat[0] + p.Lat[1]/60
}

// Longitude はこの観測地点の経度の十進表現を返却します。
func (p *Point) Longitude() float64 {
	return p.Lon[0] + p.Lon[1]/60
}
