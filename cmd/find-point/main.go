/*
find-point は観測地点の一覧から目的の観測地点の ID を探します。

Usage:

	find-point [flags]

The flags are:

	-q
		絞り込むためのキーワードです。漢字、カタカナ、英語名を部分一致で検索します。
	-h
		false を指定すると、出力にヘッダーを表示しません。
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
	"strings"

	"github.com/mikan/go-amedas"
)

func main() {
	q := flag.String("q", "", "絞り込みキーワード (例: Island)")
	header := flag.Bool("h", true, "ヘッダーの表示・非表示")
	debug := flag.Bool("d", false, "デバッグログ出力有効・無効")
	flag.Parse()
	client := amedas.NewDefaultClient()
	if *debug {
		client = amedas.NewClient(http.DefaultClient, log.Default())
	}
	points, err := client.ListPoints(context.Background())
	if err != nil {
		log.Fatalf("観測地点の取得に失敗しました: %v", err)
	}
	if *header {
		fmt.Println("id,lat,lon,kjName,knName,enName")
	}
	for _, p := range points {
		if len(*q) > 0 && !strings.Contains(p.KjName, *q) && !strings.Contains(p.KnName, *q) && !strings.Contains(p.EnName, *q) {
			continue
		}
		fmt.Printf("%s,%.3f,%.3f,%s,%s,%s\n",
			p.ID, p.Latitude(), p.Longitude(), p.KjName, p.KnName, p.EnName)
	}
}
