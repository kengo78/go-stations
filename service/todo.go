package service

import (
	"context"
	"database/sql"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	// 「INSERT INTO テーブル名 (カラム1, カラム2, ...) VALUES (?, ?, ...)」のように定義します。ここでは、subject と description を保存するため、カラムはそれぞれの名前を指定し、値の代わりに ? を指定します。これは後から設定するためのプレースホルダーとして利用します。
// 	次に、ExecContext メソッドを利用して SQL 文を実行します。ExecContext メソッドは、引数に渡された SQL 文を実行し、実行結果を返します。Insert 文の場合は、実行結果として ID を返すため、変数 result に代入します。

// その後、result.LastInsertId() メソッドを利用して、保存した TODO の ID を取得します。

// 最後に、QueryRowContext メソッドを利用して、保存した TODO を取得し、model.TODO 型のインスタンスを作成して戻り値として返します。Scan メソッドを利用して、取得したデータを model.TODO 型のインスタンスに設定します。このメソッドは、Scan(dest ...interface{}) error のように引数にデータを設定する先のポインタを指定します。今回の場合は、todo の各フィールドを指定します。

// なお、サブミットする前に、実際に動作確認するためには、依存パッケージのインストールや DB の準備が必要です
	result, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, confirm, id)
	todo := &model.TODO{}
	err = row.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return todo, nil
	// return nil, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
