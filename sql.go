package main

import (
	"database/sql"
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}

	defer db.Close()
	insert(db)
	//update(db)
	//delete(db)
	findByAll(db)
	//findByPK(db)
}

/*
*主キー検索
*引数：db
*戻り値：なし
*/
func findByPK(db *sql.DB) {
	var (
		deptno int
		dname string
		loc string
	)
	stmt, err := db.Prepare("SELECT * FROM dept WHERE deptno = ?")
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(10).Scan(&deptno, &dname, &loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deptno, dname, loc)
}



/*
*全件取得
*引数：db
*戻り値：なし
*/
func findByAll(db *sql.DB) {
	var (
		deptno int
		dname string
		loc string
	)

	rows, err := db.Query("SELECT * FROM dept")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&deptno, &dname, &loc)
		if err != nil {
			log.Fatal(err)
		}
		//結果
		fmt.Println(deptno, dname, loc)
	}
}

/*
*追加
*引数：db
*戻り値：なし
*/
func insert(db *sql.DB) {
	tr,err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO dept(deptno, dname, loc) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		 tr.Rollback()
	}
	res, err := stmt.Exec(50, "WATATAKU", "JAPAN")
	if err != nil {
		log.Fatal(err)
		 tr.Rollback()
	}
	rowCnt, err := res.RowsAffected()  // 影響を受けた行数
	if err != nil {
		log.Fatal(err)
		 tr.Rollback()
	}
	log.Printf("%d件追加しました", rowCnt)
	tr.Commit()
}

/*
*更新
*引数：db
*戻り値：なし
*/
func update(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE dept SET dname=?,loc=? WHERE deptno=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("KAWAAI", "KORIA", 50)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()  // 影響を受けた行数
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d件更新しました", rowCnt)
}

/*
*削除
*引数：db
*戻り値：なし
*/
func delete(db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM dept WHERE deptno=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(50)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()  // 影響を受けた行数
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d件削除しました", rowCnt)
}