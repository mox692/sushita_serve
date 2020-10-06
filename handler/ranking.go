package handler

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"sushita_serve/db"

	"golang.org/x/xerrors"
)

type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}
type UserRanking struct {
	ID       int
	UserID   string
	UserName string
	Score    int
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

	// ヘッダから取得したtokenを用いてuser認証
	token := r.Header.Get("user-token")
	_, err := selectUserByToken(token)
	if err != nil {
		log.Fatal("%s", err)
		return
	}
	fmt.Printf("token: %s\n", token)

	userRankings, err := selectAllRankingData()
	if err != nil {
		log.Fatal("%s", err)
		return
	}

	fmt.Fprintf(w, "1位の名前: %s\n", userRankings[0].UserName)
	fmt.Fprintf(w, "2位の名前: %s\n", userRankings[1].UserName)
}

func selectAllRankingData() ([]*UserRanking, error) {
	rows, err := db.Conn.Query("SELECT * FROM .user_ranking;")
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return convertToRanking(rows)
}

func convertToRanking(rows *sql.Rows) ([]*UserRanking, error) {
	userRankings := []*UserRanking{}
	for rows.Next() {
		userRanking := UserRanking{}
		err := rows.Scan(&userRanking.ID, &userRanking.UserID, &userRanking.UserName, &userRanking.Score)
		if err != nil {
			return userRankings, fmt.Errorf(": %w", err)
		}
		userRankings = append(userRankings, &userRanking)
	}
	return userRankings, nil
}

func selectUserByToken(token string) (*UserRanking, error) {
	row := db.Conn.QueryRow("select * from user_ranking where uesr_id = ?", token)
	return convertToUserRanking(row)
}

func convertToUserRanking(row *sql.Row) (*UserRanking, error) {
	userRanking := UserRanking{}
	err := row.Scan(&userRanking.ID, &userRanking.UserID, &userRanking.UserName, &userRanking.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("row.Scan error: %w", err)
	}
	return &userRanking, nil
}
