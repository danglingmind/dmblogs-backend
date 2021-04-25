package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlStore struct {
	db *sql.DB
}

func NewMySqlStore() *MySqlStore {
	return &MySqlStore{}
}

var _ Database = &MySqlStore{}

func (ms *MySqlStore) Open(ctx context.Context, host, username, password, dbname string, port int) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, fmt.Sprintf("%d", port), dbname)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		ms.db = nil
		return err
	}
	ms.db = conn
	fmt.Printf("MySql is runnign at %s", dsn)
	return nil
}

func (ms *MySqlStore) Close(ctx context.Context) error {
	return ms.db.Close()
}

func (ms *MySqlStore) QueryRow(ctx context.Context, q string, params ...interface{}) (Row, error) {
	res, err := ms.Query(ctx, q, params)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errors.New("one row expected from query got many")
	}
	return res[0], nil
}

func (ms *MySqlStore) Query(ctx context.Context, q string, params ...interface{}) ([]Row, error) {
	stmt, err := ms.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := make([]Row, 0)
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(Row)
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		result = append(result, m)
	}
	return result, nil
}
func (ms *MySqlStore) Save(ctx context.Context, q string, params ...interface{}) error {
	stmt, err := ms.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// TODO: find a proper way to do this
	args := params
	_, err = stmt.ExecContext(ctx, args...)
	return err
}
