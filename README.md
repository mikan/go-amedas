go-amedas
---------

Go 言語 (golang) で気象庁のアメダスの JSON データを利用するためのデータ定義とクライアントライブラリです。

## 機能

- 観測地点の一覧取得: `MapPoints()`, `ListPoints()`
- 全観測地点の最新または指定時刻の観測データ取得: `AllPointMeasurement()`, `LatestAllPointMeasurement()`
- 特定の観測地点の最新または指定時刻の観測データ取得: `SinglePointMeasurements()`, `LatestSinglePointMeasurement()`

観測地点を指定してデータを取得すると、その日のその時点までの最高気温、最低気温、最大瞬間風速の情報も併せて得られます。

## 使用例

```go
package main

import (
	"context"
	"fmt"

	"github.com/mikan/go-amedas"
)

func main() {
	m, _ := amedas.NewDefaultClient().LatestSinglePointMeasurement(context.TODO(), "44132")
	fmt.Println(m)
	// pressure=1005.8,normalPressure=1008.8,temp=7.8,humidity=55,...
}
```

備考: 44132 は東京の観測地点 ID

### サンプル main プログラム

詳しくは各プログラムのパッケージコメントをご覧ください。

- cmd/find-point - 観測地点の一覧から目的の観測地点の ID を検索
- cmd/all-point-data - 全観測地点の最新の観測データを取得して表示
- cmd/single-point-data - 指定した観測地点の最新の観測データを取得して表示

## 免責事項

- 完全に無保証です。気象庁の正式なサービスではなく、今後仕様変更等で予告なく使えなくなる可能性があります。
- メンテナンスや障害等でデータが欠損したり、データ取得自体ができないタイミングがある可能性があります。
- 観測データは後で修正・変更されることがあります。

## LICENSE

[BSD 3-clause](LICENSE)

## Author

- [mikan](https://github.com/mikan)
