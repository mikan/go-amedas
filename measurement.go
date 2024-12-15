package amedas

import (
	"fmt"
)

var windDirectionLabels = []string{
	"北北東", "北東", "東北東", "東", "東南東", "南東", "南南東", "南", // [0]=1 [7]=8
	"南南西", "南西", "西南西", "西", "西北西", "北西", "北北西", "北", // [8]=9 [15]=16
}

// HourMinute は計測データ中に出現する時分を格納するデータの構造を示します。
type HourMinute struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

// Measurement は計測データの構造を示します。
type Measurement struct {
	Pressure          []float64   `json:"pressure,omitempty"`          // 現地気圧 (hPa)
	NormalPressure    []float64   `json:"normalPressure,omitempty"`    // 海面気圧 (hPa)
	Temp              []float64   `json:"temp,omitempty"`              // 気温 (℃)
	Humidity          []int       `json:"humidity,omitempty"`          // 湿度 (%)
	Visibility        []float64   `json:"visibility,omitempty"`        // 視程 (km)
	Snow              []int       `json:"snow,omitempty"`              // 積雪 (cm)
	Snow1H            []int       `json:"snow1h,omitempty"`            // 1時間降雪量 (cm)
	Snow6H            []int       `json:"snow6h,omitempty"`            // 6時間降雪量 (cm)
	Snow12H           []int       `json:"snow12h,omitempty"`           // 12時間降雪量 (cm)
	Snow24H           []int       `json:"snow24h,omitempty"`           // 24時間降雪量 (cm)
	Sun10M            []float64   `json:"sun10m,omitempty"`            // 10分間の日照時間 (m, 0-10)
	Sun1H             []float64   `json:"sun1h,omitempty"`             // 1時間の日照時間 (h, 0-1)
	Precipitation10M  []float64   `json:"precipitation10m,omitempty"`  // 10分間降水量 (mm)
	Precipitation1H   []float64   `json:"precipitation1h,omitempty"`   // 1時間間降水量 (mm)
	Precipitation3H   []float64   `json:"precipitation3h,omitempty"`   // 3時間降水量 (mm)
	Precipitation24H  []float64   `json:"precipitation24h,omitempty"`  // 24時間降水量 (mm)
	WindDirection     []int       `json:"windDirection,omitempty"`     // 風向 (1=北北東, 2=北東 ... 16=北)
	Wind              []float64   `json:"wind,omitempty"`              // 風速 (m/s)
	PrefNumber        *int        `json:"prefNumber,omitempty"`        // 都道府県 JIS コード
	ObservationNumber *int        `json:"observationNumber,omitempty"` // 観測地点番号
	MaxTempTime       *HourMinute `json:"maxTempTime,omitempty"`       // 最高気温を記録した時刻
	MaxTemp           *[]float64  `json:"maxTemp,omitempty"`           // 最高気温
	MinTempTime       *HourMinute `json:"minTempTime,omitempty"`       // 最低気温を記録した時刻
	MinTemp           *[]float64  `json:"minTemp,omitempty"`           // 最低気温
	GustTime          *HourMinute `json:"gustTime,omitempty"`          // 当日中の最大瞬間風速の時刻
	GustDirection     *[]int      `json:"gustDirection,omitempty"`     // 当日中の最大瞬間風速の風向
	Gust              *[]float64  `json:"gust,omitempty"`              // 当日中の最大瞬間風速
}

func (m Measurement) String() string {
	var s string
	delim := ""
	if len(m.Pressure) > 0 {
		s += fmt.Sprintf("%spressure=%v", delim, m.Pressure[0])
		delim = ","
	}
	if len(m.Pressure) > 0 {
		s += fmt.Sprintf("%snormalPressure=%v", delim, m.NormalPressure[0])
		delim = ","
	}
	if len(m.Temp) > 0 {
		s += fmt.Sprintf("%stemp=%v", delim, m.Temp[0])
		delim = ","
	}
	if len(m.Humidity) > 0 {
		s += fmt.Sprintf("%shumidity=%v", delim, m.Humidity[0])
		delim = ","
	}
	if len(m.Visibility) > 0 {
		s += fmt.Sprintf("%svisibility=%v", delim, m.Visibility[0])
		delim = ","
	}
	if len(m.Snow) > 0 {
		s += fmt.Sprintf("%ssnow=%v", delim, m.Snow[0])
		delim = ","
	}
	if len(m.Snow1H) > 0 {
		s += fmt.Sprintf("%ssnow1h=%v", delim, m.Snow1H[0])
		delim = ","
	}
	if len(m.Snow6H) > 0 {
		s += fmt.Sprintf("%ssnow6h=%v", delim, m.Snow6H[0])
		delim = ","
	}
	if len(m.Snow12H) > 0 {
		s += fmt.Sprintf("%ssnow12h=%v", delim, m.Snow12H[0])
		delim = ","
	}
	if len(m.Snow24H) > 0 {
		s += fmt.Sprintf("%ssnow24h=%v", delim, m.Snow24H[0])
		delim = ","
	}
	if len(m.Sun10M) > 0 {
		s += fmt.Sprintf("%ssun10m=%v", delim, m.Sun10M[0])
		delim = ","
	}
	if len(m.Sun1H) > 0 {
		s += fmt.Sprintf("%ssun1h=%v", delim, m.Sun1H[0])
		delim = ","
	}
	if len(m.Precipitation10M) > 0 {
		s += fmt.Sprintf("%sprecipication10m=%v", delim, m.Precipitation10M[0])
		delim = ","
	}
	if len(m.Precipitation1H) > 0 {
		s += fmt.Sprintf("%sprecipication1h=%v", delim, m.Precipitation1H[0])
		delim = ","
	}
	if len(m.Precipitation3H) > 0 {
		s += fmt.Sprintf("%sprecipication3h=%v", delim, m.Precipitation3H[0])
		delim = ","
	}
	if len(m.Precipitation24H) > 0 {
		s += fmt.Sprintf("%sprecipication24h=%v", delim, m.Precipitation24H[0])
		delim = ","
	}
	if len(m.WindDirection) > 0 {
		s += fmt.Sprintf("%swindDirection=%d(%s)", delim, m.WindDirection[0], m.WindDirectionLabel())
		delim = ","
	}
	if len(m.Wind) > 0 {
		s += fmt.Sprintf("%swind=%v", delim, m.Wind[0])
		delim = ","
	}
	if m.PrefNumber != nil {
		s += fmt.Sprintf("%sprefNumber=%d", delim, *m.PrefNumber)
	}
	if m.ObservationNumber != nil {
		s += fmt.Sprintf("%sobservationNumber=%d", delim, *m.ObservationNumber)
	}
	if m.MaxTemp != nil && len(*m.MaxTemp) > 0 {
		s += fmt.Sprintf("%smaxTemp=%v", delim, (*m.MaxTemp)[0])
	}
	if m.MaxTempTime != nil {
		s += fmt.Sprintf("%smaxTempTime=%d:%02d", delim, (*m.MaxTempTime).Hour, (*m.MaxTempTime).Minute)
	}
	if m.MinTemp != nil && len(*m.MinTemp) > 0 {
		s += fmt.Sprintf("%sminTemp=%v", delim, (*m.MinTemp)[0])
	}
	if m.MinTempTime != nil {
		s += fmt.Sprintf("%sminTempTime=%d:%02d", delim, (*m.MinTempTime).Hour, (*m.MinTempTime).Minute)
	}
	if m.Gust != nil && len(*m.Gust) > 0 {
		s += fmt.Sprintf("%sgust=%v", delim, (*m.Gust)[0])
	}
	if m.GustDirection != nil && len(*m.GustDirection) > 0 {
		s += fmt.Sprintf("%sgustDirection=%d(%s)", delim, (*m.GustDirection)[0], m.GustDirectionLabel())
	}
	if m.GustTime != nil {
		s += fmt.Sprintf("%sgustTime=%d:%02d", delim, (*m.GustTime).Hour, (*m.GustTime).Minute)
	}
	return s
}

// WindDirectionLabel は風向の日本語表記を出力します。
func (m Measurement) WindDirectionLabel() string {
	if len(m.WindDirection) == 0 || m.WindDirection[0] <= 0 || m.WindDirection[0] > len(windDirectionLabels) {
		return ""
	}
	return windDirectionLabels[m.WindDirection[0]-1]
}

// GustDirectionLabel は最大瞬間風速の風向の日本語表記を出力します。
func (m Measurement) GustDirectionLabel() string {
	if m.GustDirection == nil || len(*m.GustDirection) == 0 || (*m.GustDirection)[0] <= 0 || (*m.GustDirection)[0] > len(windDirectionLabels) {
		return ""
	}
	return windDirectionLabels[(*m.GustDirection)[0]-1]
}
