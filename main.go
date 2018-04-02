package main

import "fmt"
import "time"

func main() {
	// 実行すると
	// 必要な情報をいろいろインタラクティブに聞いてきて
	// (初期値) で
	// _postsにファイルを吐き出してくれるだけ で良い
	// title, description, date, filename
	// とりあえずこれだけでいい

	// 必要なもの
	// 入力待ち ok
	// ファイル生成
	// ファイルの重複検知
	// とりあえずこんなもん？

	//h := header{
	//	title:       "foo",
	//	description: "brah brah",
	//	date:        "2018-04-02",
	//	filename:    "test",
	//}

	h := map[string]string{
		"title":       "foo",
		"description": "brah brah",
		"date":        "2018-04-02",
		"filename":    "test",
	}

	year, month, day := time.Now().Date()
	m := fmt.Sprintf("%02d", month)
	d := fmt.Sprintf("%02d", day)

	fmt.Printf("Today is %d-%s-%s\n", year, m, d)

	scans := map[string]string{}

	var temp string
	for k, v := range h {
		fmt.Printf("%s? (%s): ", k, v)
		temp = v
		fmt.Scan(&temp)
		fmt.Println(temp)
		scans[k] = temp
	}

	fmt.Println(scans)

}
