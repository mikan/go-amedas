/*
single-point-data は指定した観測地点の最新の観測データを取得して表示します。

Usage:

	single-point-data [flags]

The flags are:

	-p
		観測地点 ID、フラグを指定しない場合はプロンプトを表示し入力を促します。
	-d
		true を指定すると、デバッグログを出力します。
*/
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mikan/go-amedas"
)

func main() {
	point := flag.String("p", "", "観測地点 ID (例: 44132)")
	debug := flag.Bool("d", false, "デバッグログ出力有効・無効")
	flag.Parse()
	if *point == "" {
		fmt.Printf("観測地点 ID (例: 44132) > ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			*point = scanner.Text()
			if _, err := strconv.Atoi(*point); err == nil {
				break // 数値に変換できるかどうか試し、パスしたら ID とみなす
			}
		}
	}
	client := amedas.NewDefaultClient()
	if *debug {
		client = amedas.NewClient(http.DefaultClient, log.Default())
	}
	m, err := client.LatestSinglePointMeasurement(context.Background(), *point)
	if err != nil {
		log.Fatalf("観測データの取得に失敗しました: %v", err)
	}
	fmt.Println(*m)
}
