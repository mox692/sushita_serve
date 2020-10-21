package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sushita_serve/db"
	"sushita_serve/response"

	"golang.org/x/xerrors"
)

type setRankingRequestBody struct {
	Name  string `json: "name"`
	Score int    `json: "score"`
}

type RankingHandler struct {
	token          string
	setRequestBody setRankingRequestBody
	db             *sql.DB
}

// リクエストからヘッダを取得し、構造体にセット
// Todo: この関数もメソッド化してテストしやすくしたい。。
// SetRankingじゃなくて、呼び出しもとから選択できるようにしたい。
func HandleSetRanking(w http.ResponseWriter, r *http.Request) {
	var requestBody setRankingRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		response.ErrResponse(w, xerrors.Errorf(": %w", err), 500)
		return
	}
	rh := &RankingHandler{
		token:          r.Header.Get("user-token"),
		setRequestBody: requestBody,
		db:             db.Conn,
	}
	rh.setRanking(w, r)
}

func HandleGetRanking(w http.ResponseWriter, r *http.Request) {
	var requestBody setRankingRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		response.ErrResponse(w, xerrors.Errorf(": %w", err), 500)
	}
	rh := &RankingHandler{
		token:          r.Header.Get("user-token"),
		setRequestBody: requestBody,
		db:             db.Conn,
	}
	rh.getRanking(w, r)
}

type rankingGetResponse struct {
	userName string `json: "user_name"`
	score    int    `json: "score"`
}

func (rh *RankingHandler) getRanking(w http.ResponseWriter, r *http.Request) {
	userRankings, err := rh.selectAllRankingData()
	if err != nil {
		response.ErrResponse(w, xerrors.Errorf(": %w", err), 500)
	}
	for i := 0; i < len(userRankings); i++ {
		fmt.Fprintf(w, "%d: %s\n", i+1, userRankings[i].UserName)
	}

	responseDatas := []*rankingGetResponse{}

	for _, v := range userRankings {
		responseData := rankingGetResponse{}
		responseData.userName = v.UserName
		responseData.score = v.Score
		responseDatas = append(responseDatas, &responseData)
	}

	response.SuccessResponse(w, responseDatas)
}

func (rh *RankingHandler) setRanking(w http.ResponseWriter, r *http.Request) {
	// Todo: logicを工夫して、db接続を減らす
	_, err := rh.selectUserRankingByToken(rh.token)
	if err != nil {
		if err == sql.ErrNoRows {
			// !insertRankingDataByTokenが正常のときはerrにはnilが代入される
			log.Println("doing insert...")
			err = rh.insertRankingDataByToken(rh.token)
		}
		fmt.Printf("%+v\n", err)
	} else {
		log.Println("doing update...")
		err = rh.updateUserRankingByToken(rh.token)
	}

	// errをまとめて処理
	if err != nil {
		response.ErrResponse(w, xerrors.Errorf(": %w", err), 500)
		return
	}

	// normal finish log
	log.Printf("Completed normally!!")

	fmt.Fprintln(w, "Your score just has been registered in the server!")

}

func (rh *RankingHandler) selectAllRankingData() ([]*db.UserRanking, error) {
	rows, err := rh.db.Query("SELECT * FROM .user_ranking;")
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return convertToRanking(rows)
}

func (rh *RankingHandler) insertRankingDataByToken(token string) error {
	stmt, err := rh.db.Prepare("INSERT INTO user_ranking (user_id, user_name, score) VALUES (?, ?, ?);")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	_, err = stmt.Exec(token, rh.setRequestBody.Name, rh.setRequestBody.Score)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}

func (rh *RankingHandler) updateUserRankingByToken(token string) error {
	stmt, err := rh.db.Prepare("update user_ranking SET score = ? where id = ?")
	if err != nil {
		return xerrors.Errorf("db.Conn.Prepare error: %w", err)
	}
	_, err = stmt.Exec(rh.setRequestBody, token)
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
			return userRankings, xerrors.Errorf(": %w", err)
		}
		userRankings = append(userRankings, &userRanking)
	}
	return userRankings, nil
}

func (rh *RankingHandler) selectUserRankingByToken(token string) (*db.UserRanking, error) {
	row := rh.db.QueryRow("select * from user_ranking where user_id = ?", rh.token)
	return convertToUserRanking(row)
}

func convertToUserRanking(row *sql.Row) (*db.UserRanking, error) {
	userRanking := db.UserRanking{}
	err := row.Scan(&userRanking.ID, &userRanking.UserID, &userRanking.UserName, &userRanking.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, xerrors.Errorf(": %w", err)
		}
		return nil, xerrors.Errorf("row.Scan error: %w", err)
	}
	return &userRanking, nil
}
