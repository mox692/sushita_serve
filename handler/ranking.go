package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sushita_serve/db"

	"golang.org/x/xerrors"
)

func GetRanking(w http.ResponseWriter, r *http.Request) {
	log.Println("アクセスが来ました！")
	fmt.Fprintln(w, "サーバーからの書き込みです!!")

	// ヘッダから取得したtokenを用いてuser認証
	token := r.Header.Get("user-token")
	_, err := selectUserRankingByToken(token)
	if err != nil {
		log.Fatal("%s", err)
		fmt.Fprintf(w, "invalide token.\n", token)
		return
	}
	fmt.Fprintf(w, "token: %s\n", token)

	userRankings, err := selectAllRankingData()
	if err != nil {
		log.Fatal("%s", err)
	}

	for i := 0; i < len(userRankings); i++ {
		fmt.Fprintf(w, "%d: %s\n", i+1, userRankings[i].UserName)
	}
}

func SetRanking(w http.ResponseWriter, r *http.Request) {

	// リクエストを確認
	log.Printf("リクエストの表示です!!: %#v\n\n", r)
	log.Printf("リクエストbodyの表示です!!: %#v\n\n", r.Body)

	// デコード
	var requestBody setRankingRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Fatalln("decode err")
	}
	log.Printf("デコード後のリクエストbodyの表示です!!: %#v", requestBody)

	token := r.Header.Get("user-token")

	// connection test
	w.Write([]byte("dfsalj"))
	fmt.Fprintf(w, token)

	// tokenを使用して、dbに接続、対象userのレコードがすでに存在するか確認
	_, err = selectUserRankingByToken(token)
	// もし存在していなかったらinsert
	if err != nil {
		if err == sql.ErrNoRows {
			requestBody.insertRankingDataByToken(token, &requestBody)
		}
		fmt.Errorf(" :%w", err)
	} else {
		requestBody.updateUserRankingByToken(token, &requestBody)
	}
	// もし存在していたらupdate
}

func selectAllRankingData() ([]*db.UserRanking, error) {
	rows, err := db.Conn.Query("SELECT * FROM .user_ranking;")
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return convertToRanking(rows)
}

func (rb *setRankingRequestBody) insertRankingDataByToken(token string, requestBody *setRankingRequestBody) {

}

func (rb *setRankingRequestBody) updateUserRankingByToken(token string, requestBody *setRankingRequestBody) {

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

func selectUserRankingByToken(token string) (*db.UserRanking, error) {
	row := db.Conn.QueryRow("select * from user_ranking where user_id = ?", token)
	return convertToUserRanking(row)
}

func convertToUserRanking(row *sql.Row) (*db.UserRanking, error) {
	userRanking := db.UserRanking{}
	err := row.Scan(&userRanking.ID, &userRanking.UserID, &userRanking.UserName, &userRanking.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, xerrors.Errorf("row.Scan error: %w", err)
	}
	return &userRanking, err
}

type setRankingRequestBody struct {
	Name  string `json: "name"`
	Score int    `json: "score"`
}
