package main

import (
	"encoding/csv"
	"log"
	"net/url"
	"os"

	"github.com/nakabonne/netsurfer"
	"github.com/spf13/cast"
)

func main() {
	keyWords := []string{
		"漫画 無料",
		"無料 漫画",
		"マンガ 無料",
		"無料 マンガ",
	}

	urlList := []string{
		"https://manga-zero.com/",
		"https://www.sukima.me/",
		"https://www.cmoa.jp/freecontents/",
		"https://www.cmoa.jp/freecontents/title/girl/",
		"https://sokuyomi.jp/free.html",
		"https://www.ebookjapan.jp/ebj/free/",
		"https://renta.papy.co.jp/renta/sc/frm/page/topics/c_freeall.htm",
		"https://plus.comicess.com/",
		"https://bookstore.yahoo.co.jp/free/",
		"https://comic.k-manga.jp/",
		"https://dokusho-ojikan.jp/free_list/page_type=all/sort=series_rank",
		"https://www.mangaz.com/mens/",
		"https://comic.pixiv.net/categories/ファンタジー",
		"https://dokuha.jp/",
		"https://comic-walker.com/contents/list/",
		"https://manga-bang.com/",
	}

	var writer *csv.Writer
	// ファイルを書き込みモードでオープン(ファイルがなかったら作成する)
	writeFile, _ := os.OpenFile("./ranking.csv", os.O_WRONLY|os.O_CREATE, 0600)
	// Writerを書き込みモードでオープン
	writer = csv.NewWriter(writeFile)

	for _, list := range urlList {
		u, _ := url.Parse(list)
		//fmt.Println("URL:", list)
		csvURL := []string{list}
		for _, words := range keyWords {
			rank, _ := netsurfer.GetRank(u, words, 2)
			//fmt.Println("KeyWord:", words, "Rank:", rank)
			csvURL = append(csvURL, cast.ToString(rank))
		}
		writer.Write(csvURL)
	}
	writer.Flush()
}
func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
