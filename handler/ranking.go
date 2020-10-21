package handler

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"sushita_serve/db"
)

type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

var TemplateOfLog = `
Remote address:   {{.RemoteAddr}}
Content-Type:     {{.ContentType}}
HTTP method:      {{.Method}}

path:
{{.Path}}

query string:
{{.Query}}

body:             
{{.Body}}
`

// rankingエンドポイントを司るHandler
// rにリクエスト情報を詰めて、処理の中で生じたerr等はwに書き込む。
func RankingGet(w http.ResponseWriter, r *http.Request) {

	// wの動作確認。
	fmt.Fprintf(w, "Hello, World\n")
	fmt.Fprintf(w, "Request: %v\n", r)

	// リクエストbodyの内容を整形してwに書き込み
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	log.Println(bufbody, "\n=====================")
	body := bufbody.String()
	log.Println(body, "\n=====================")

	line := LineOfLog{
		r.RemoteAddr,
		r.Header.Get("Content-Type"),
		r.URL.Path,
		r.URL.RawQuery,
		r.Method, body,
	}
	tmpl, err := template.New("line").Parse(TemplateOfLog)
	if err != nil {
		panic(err)
	}

	bufline := new(bytes.Buffer)
	err = tmpl.Execute(bufline, line)
	if err != nil {
		panic(err)
	}

	log.Printf(bufline.String())
}

func GetRanking(w http.ResponseWriter, r *http.Request) {
	log.Println("アクセスが来ました！")
	fmt.Fprintln(w, "サーバーからの書き込みです!!")

	userRankings, err := selectAllRankingData()
	if err != nil {
		log.Fatal("%s", err)
	}

	fmt.Fprintf(w, "一位の名前: %s", userRankings[0].UserName)
}

func selectAllRankingData() ([]*db.UserRanking, error) {
	rows, err := db.Conn.Query("SELECT * FROM .user_ranking;")
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return convertToRanking(rows)
}

func convertToRanking(rows *sql.Rows) ([]*db.UserRanking, error) {
	userRankings := []*db.UserRanking{}

	for rows.Next() {
		userRanking := db.UserRanking{}
		err := rows.Scan(&userRanking.ID, &userRanking.UserID, &userRanking.UserName, &userRanking.Score)
		if err != nil {
			return userRankings, fmt.Errorf(": %w", err)
		}
		userRankings = append(userRankings, &userRanking)
	}

	return userRankings, nil
}
