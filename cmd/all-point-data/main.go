/*
all-point-data は全観測地点の最新の観測データを取得して表示します。

Usage:

	all-point-data [flags]

The flags are:

	-p
		観測地点 ID、画面に表示する観測地点を1箇所に制限したい場合に指定します。
	-d
		true を指定すると、デバッグログを出力します。
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mikan/go-amedas"
)

func main() {
	point := flag.String("p", "", "観測地点 ID (例: 44132)")
	debug := flag.Bool("d", false, "デバッグログ出力有効・無効")
	flag.Parse()
	client := amedas.NewDefaultClient()
	if *debug {
		client = amedas.NewClient(http.DefaultClient, log.Default())
	}
	mp, err := client.MapPoints(context.Background())
	if err != nil {
		log.Fatalf("観測地点の取得に失敗しました: %v", err)
	}
	m, err := client.LatestAllPointMeasurement(context.Background())
	if err != nil {
		log.Fatalf("観測データの取得に失敗しました: %v", err)
	}
	for p, measurement := range m {
		if *point != "" && p != *point {
			continue
		}
		fmt.Printf("%s(%s): %s\n", p, mp[p].KjName, measurement)
	}
}
