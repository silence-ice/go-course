package example

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"time"
)

// ///////////////////////////////////////声明数据库结构///////////////////////////////////////////////////////
// 红包商品表,默认有符号类型字段
type GoodsSigned struct {
	Goods
}

// 红包商品表,无符号类型字段
type GoodsUnsigned struct {
	Goods
}

// 红包商品
type Goods struct {
	RemainAmount   decimal.Decimal `db:"remain_amount"`        //红包剩余金额额
	RemainQuantity int             `db:"remain_quantity"`      //红包剩余数量
	EnvelopeNo     string          `db:"envelope_no,uni"`      //红包编号,红包唯一标识
	CreatedAt      time.Time       `db:"created_at,omitempty"` //创建时间
	UpdatedAt      time.Time       `db:"updated_at,omitempty"` //更新时间
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

var db *dbx.Database

func init() {
	settings := dbx.Settings{
		DriverName: "mysql",
		Host:       "127.0.0.1:3306",
		User:       "root",
		Password:   "huangbin0236",
		Database:   "resk",
		Options: map[string]string{
			"parseTime": "true",
		},
	}
	var err error
	db, err = dbx.Open(settings)
	if err != nil {
		fmt.Println(err)
	}
	db.RegisterTable(&GoodsSigned{}, "goods")
	db.RegisterTable(&GoodsUnsigned{}, "goods_unsigned")
}

// 事务锁方案
var query = "select * from goods " +
	"where envelope_no=? " +
	"for update"
var update = "update goods " +
	"set remain_amount=?,remain_quantity=? " +
	"where envelope_no=? "

func UpdateForLock(g Goods) {
	//通过db.tx函数构建事务锁代码块
	err := db.Tx(func(runner *dbx.TxRunner) error {
		//第一步：锁定需要修改的资源，也就是需要修改的数据行
		//编写事务锁查询语句，使用for update子句来锁定资源
		out := &GoodsSigned{}
		ok, err := runner.Get(out, query, g.EnvelopeNo)
		if !ok || err != nil {
			return err
		}

		//第二部：计算剩余金额和剩余数量
		subAmount := decimal.NewFromFloat(0.01)
		remainAmount := out.RemainAmount.Sub(subAmount)
		remainQuantity := out.RemainQuantity - 1

		//第三步：执行更新
		_, row, err := runner.Execute(update, remainAmount, remainQuantity, g.EnvelopeNo)
		if err != nil {
			return err
		}
		if row < 1 {
			return errors.New("库存扣减失败")
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

// 数据库无符号类型+直接更新方案
func UpdateForUnsigned(g Goods) {
	update := "update goods_unsigned " +
		"set remain_amount=remain_amount-?, " +
		"remain_quantity=remain_quantity-? " +
		"where envelope_no=?"
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}

// 乐观锁方案（重点是where条件）
func UpdateForOptimistic(g Goods) {
	update := "update goods " +
		"set remain_amount=remain_amount-?, " +
		" remain_quantity=remain_quantity-? " +
		" where envelope_no=? " +
		" and remain_amount>=? " +
		" and remain_quantity>=? "
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}

// 乐观锁+无符号字段双保险方案
func UpdateForOptimisticAndUnsigned(g Goods) {
	update := "update goods_unsigned " +
		"set remain_amount=remain_amount-?, " +
		" remain_quantity=remain_quantity-? " +
		" where envelope_no=? " +
		" and remain_amount>=? " +
		" and remain_quantity>=? "
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}
