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

type setRankingRequestBody struct {
	Name  string `json: "name"`
	Score int    `json: "score"`
}

func SetRanking(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("user-token")

	var requestBody setRankingRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Fatalln("decode err")
	}

	// Todo: logicを工夫して、db接続を減らす
	_, err = selectUserRankingByToken(token)
	if err != nil {
		if err == sql.ErrNoRows {
			// !insertRankingDataByTokenが正常のときはerrにはnilが代入される
			log.Println("doing insert...")
			err = requestBody.insertRankingDataByToken(token)
		}
		fmt.Errorf(" :%w", err)
	} else {
		log.Println("doing update...")
		err = requestBody.updateUserRankingByToken(token)
	}

	// errをまとめて処理
	if err != nil {
		log.Fatalf(": %w", err)
		return
	}

	// normal finish log
	log.Printf("Completed normally!!")

	// Todo: レスポンスに何が必要か。レスポンス用の構造体を作成
	fmt.Fprintln(w, "サーバーからの書き込みです!!")

}

func selectAllRankingData() ([]*db.UserRanking, error) {
	rows, err := db.Conn.Query("SELECT * FROM .user_ranking;")
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return convertToRanking(rows)
}

func (rb *setRankingRequestBody) insertRankingDataByToken(token string) error {
	stmt, err := db.Conn.Prepare("INSERT INTO user_ranking (user_id, user_name, score) VALUES (?, ?, ?);")
	if err != nil {
		return fmt.Errorf("db.Conn.Prepare err : %w", err)
	}
	_, err = stmt.Exec(token, rb.Name, rb.Score)
	if err != nil {
		return fmt.Errorf("stmt.Exec err : %w", err)
	}
	return nil
}

func (rb *setRankingRequestBody) updateUserRankingByToken(token string) error {
	stmt, err := db.Conn.Prepare("update user_ranking SET score = ? where id = ?")
	if err != nil {
		return xerrors.Errorf("db.Conn.Prepare error: %w", err)
	}
	_, err = stmt.Exec(rb.Score, token)
	if err != nil {
		return xerrors.Errorf("stmt.Exec error: %w", err)
	}
	return nil
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
