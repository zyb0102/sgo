package orm

import (
	"fmt"
	"testing"
)

func TestNewGorm(t *testing.T) {
	gormOpt := GormOpt{
		Host:     "47.105.53.133",
		Port:     "3306",
		DBName:   "orm_test_v1",
		Username: "orm_test_v1",
		Password: "f2yARKBps5Xe7iHN",
		Charset:  "utf8mb4",
		Loc:      "Asia%2FShanghai",
	}
	db, err := NewGorm(&gormOpt)
	if err != nil {
		fmt.Println(err)
	}
	type user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var u []user
	u = append(u, user{Name: "张三"})
	u = append(u, user{Name: "李四"})
	u = append(u, user{Name: "王五"})
	u = append(u, user{Name: "杨柳"})

	tx := db.Table("user").Create(u)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(u)
}
