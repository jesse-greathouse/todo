package todo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const ITEM_TABLE_CREATE_SQL = "item-table-create.sql"
const ITEM_TABLE_SELECT_ALL_SQL = "item-table-select-all.sql"
const ITEM_TABLE_SELECT_BY_ID_SQL = "item-table-select-by-id.sql"
const ITEM_TABLE_MAX_PRIORITY_SQL = "item-table-max-priority.sql"
const ITEM_TABLE_INSERT_SQL = "item-table-insert.sql"
const ITEM_TABLE_UPDATE_SQL = "item-table-update.sql"
const ITEM_TABLE_DELETE_SQL = "item-table-delete.sql"

type Database struct {
	Env *Environment
}

type Response struct {
	Result string
}

type FileQuery struct {
	File  string
	Query string
}

func (d Database) Init() {
	d.InitItemTable()
}

func (d Database) InitItemTable() {
	db, err := sql.Open("mysql", d.Env.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(d.GetFileQuery(ITEM_TABLE_CREATE_SQL))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func (d Database) GetFileQuery(fileName string) string {
	helpers := &Helpers{}
	return helpers.ReadFile(d.Env.SQL_PATH + "/" + fileName)
}

func (d Database) UpdateItem(i Item) Item {
	db, err := sql.Open("mysql", d.Env.DSN)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(d.GetFileQuery(ITEM_TABLE_UPDATE_SQL))
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(i.Description, i.Priority, i.Completed, i.Id)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated item ID = %d , affected = %d\n", i.Id, rowCnt)

	return d.GetItem(int(i.Id))
}

func (d Database) CreateItem(i Item) Item {
	var priority int

	db, err := sql.Open("mysql", d.Env.DSN)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow(d.GetFileQuery(ITEM_TABLE_MAX_PRIORITY_SQL)).Scan(&priority)

	stmt, err := db.Prepare(d.GetFileQuery(ITEM_TABLE_INSERT_SQL))
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(i.Description, (priority + 1))

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return i
	} else {
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		tx.Commit()

		log.Printf("Created item ID = %d, affected = %d\n", lastId, rowCnt)

		return d.GetItem(int(lastId))
	}
}

func (d Database) GetItem(itemId int) Item {
	var (
		id          int
		description string
		priority    int
		completed   int
	)
	db, err := sql.Open("mysql", d.Env.DSN)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.QueryRow(d.GetFileQuery(ITEM_TABLE_SELECT_BY_ID_SQL), itemId).Scan(&id, &description, &priority, &completed)
	if err != nil {
		log.Fatal(err)
	}

	item := Item{
		Id:          id,
		Description: description,
		Priority:    priority,
		Completed:   completed,
	}

	return item
}

func (d Database) GetAllItems() []Item {
	var (
		items       []Item
		id          int
		description string
		priority    int
		completed   int
	)
	db, err := sql.Open("mysql", d.Env.DSN)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query(d.GetFileQuery(ITEM_TABLE_SELECT_ALL_SQL))
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &description, &priority, &completed)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, Item{
			Id:          id,
			Description: description,
			Priority:    priority,
			Completed:   completed,
		})
	}

	return items
}

func (d Database) DeleteItem(itemId int) Response {
	db, err := sql.Open("mysql", d.Env.DSN)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(d.GetFileQuery(ITEM_TABLE_DELETE_SQL))
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(itemId)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	resultStr := fmt.Sprintf("Deleted item ID = %d, affected = %d", itemId, rowCnt)

	log.Print(resultStr + "\n")

	return Response{
		Result: resultStr,
	}

}
