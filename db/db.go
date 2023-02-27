package db

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/boltdb/bolt"
)

const (
	Todopath   = "todo.db"
	TodoBucket = "todos"
)

type Todo struct {
	Name      string
	Details   string
	Category  string
	Deadline  time.Time
	Create_at time.Time
	Clear     bool
}

func Conn() (*bolt.DB, error) {
	conn, err := bolt.Open(Todopath, 0600, nil)
	return conn, err
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"[rand.Intn(len("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))]
	}
	return string(b)
}
func RandomID() []byte {
	return []byte(RandStringBytes(4))
}

func AddToDo(todo Todo) error {
	db, err := Conn()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(TodoBucket))
		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		b.Put(RandomID(), buf)
		return nil
	})
}
func TodoClear(id string) bool {
	db, err := Conn()
	if err != nil {
		return false
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TodoBucket))
		buf := b.Get([]byte(id))
		if buf == nil {
			return fmt.Errorf("Todo not found")
		}

		var todo Todo
		json.Unmarshal(buf, &todo)
		todo.Clear = true
		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		b.Put([]byte(id), buf)
		return nil
	})
	if err != nil {
		return false
	}
	return true
}

func throwErr(s string) {
	panic("unimplemented")
}
func UpdateTodo(id string, todo Todo) error {
	db, err := Conn()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TodoBucket))
		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		b.Put([]byte(id), buf)
		return nil
	})
}
func DeleteTodo(id string) bool {
	db, err := Conn()
	if err != nil {
		return false
	}
	defer db.Close()
	if db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TodoBucket))
		return b.Delete([]byte(id))
	}) == nil {
		return true
	}
	return false
}

func GetToDo() (map[string]Todo, error) {
	db, err := Conn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var todo map[string]Todo = make(map[string]Todo, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TodoBucket))
		b.ForEach(func(k, v []byte) error {
			res := Todo{}
			if json.Unmarshal(v, &res) != nil {
				return json.Unmarshal(v, &res)
			}
			todo[string(k)] = res
			return nil
		})
		return nil
	})
	return todo, err
}
