package handler

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// setrankingでerrが発生しない。
// writerに正しくdbの情報が書き込まれる。
// Todo: 書き込み先のdbをモック化する。
func TestSetRanking(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rh := RankingHandler{
		token:          "test",
		setRequestBody: setRankingRequestBody{"motoyuki", 22},
		db:             db,
	}
	columns := []string{"id", "user_id", "user_name", "score"}
	mock.ExpectQuery("select").
		WithArgs(rh.token).
		WillReturnRows(sqlmock.NewRows(columns).AddRow("sfd", "1", "test title", 11))

	mock.ExpectExec("INSERT INTO user_ranking").WillReturnResult(sqlmock.NewResult(1, 1))

	var w http.ResponseWriter
	var r *http.Request
	rh.SetRanking(w, r)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %+v", err)
	}

}
