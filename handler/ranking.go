package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"sushita_serve/db"

	"golang.org/x/xerrors"
)

type UserRanking struct {
	ID       int
	UserID   string
	UserName string
	Score    int
}

func GetRanking(w http.ResponseWriter, r *http.Request) {
	log.Println("アクセスが来ました！")
	fmt.Fprintln(w, "サーバーからの書き込みです!!")

	// ヘッダから取得したtokenを用いてuser認証
	token := r.Header.Get("user-token")
	_, err := selectUserByToken(token)
	if err != nil {
		log.Fatal("%s", err)
	}
	fmt.Printf("token: %s\n", token)

	userRankings, err := selectAllRankingData()
	if err != nil {
		log.Fatal("%s", err)
	}

	for i := 0; i < len(userRankings); i++ {
		fmt.Fprintf(w, "%d: %s\n", i+1, userRankings[0].UserName)
	}
	// fmt.Fprintf(w, "1位の名前: %s\n", userRankings[0].UserName)
	// fmt.Fprintf(w, "2位の名前: %s\n", userRankings[1].UserName)
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
	row := db.Conn.QueryRow("select * from user_ranking where user_id = ?", token)
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
