package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 模型结构体
type User struct {
	gorm.Model
	Name string `gorm:"column:name;UNIQUE"`
	Age  int    `gorm:"column:age"`
}

func test01() {
	// 数据库连接信息
	dsn := "host=localhost user=postgres password=666666 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("无法连接数据库：", err)
		return
	}

	// 自动迁移创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("无法创建表：", err)
		return
	}

	// 创建用户
	user := User{Name: "John Doe", Age: 30}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("无法创建用户：", result.Error)
		return
	}

	// 查询用户
	var retrievedUser User
	result = db.First(&retrievedUser, user.ID)
	if result.Error != nil {
		fmt.Println("无法查询用户：", result.Error)
		return
	}

	fmt.Printf("ID: %d, Name: %s, Age: %d\n", retrievedUser.ID, retrievedUser.Name, retrievedUser.Age)
}

func test02() {
	pgConfig := "host=localhost user=postgres password=666666 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	//打开数据库
	db, err := gorm.Open(postgres.Open(pgConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to open postgres database.")
		return
	}

	//First查找并更新
	u := User{}
	db.First(&u)
	fmt.Printf("find value:%#v\n", u)
	db.Model(&u).Update("name", "小狐狸")

	//根据字段值查找
	uu := new(User)
	db.Find(uu, "name=?", "小狐狸")
	db.Model(uu).Update("age", "18") // Update修改指定字段

	//Find 实际传入的参数为切片 可以查出多个数据
	us := []User{}
	db.Find(&us)
	fmt.Printf("find value:%#v\n", us)
	//db.Model(&us).Update("age", "16")

	//利用Where做查询条件并更新
	db.Model(&us).Where("name = ?", "小狐狸").Update("age", 16)                    //Update
	db.Model(&us).Where("name = ?", "小狐狸").Updates(User{Name: "小小狐狸", Age: 14}) // Updates
	m1 := map[string]interface{}{
		"name": "小狐狸",
		"age":  17,
	}
	db.Model(&us).Where("name = ?", "John Doe").Updates(m1) // 传入map,相当于可以全量更新也可以更新部分值
	m1["age"] = 20
	db.Model(&us).Where("name = ?", "John Doe").Select("age").Updates(m1) // 传入map,相当于可以全量更新也可以更新部分值,仅更新age
	m1["name"] = "Jack"
	db.Model(&us).Where("name = ?", "John Doe").Omit("age").Updates(m1)                                         // 传入map,相当于可以全量更新也可以更新部分值,除age外都更新
	rows := db.Model(&us).Where("name = ?", "John Doe").UpdateColumns(User{Name: "小宇航员", Age: 20}).RowsAffected //仅会更新这两列，不会更新其他的updated_at之类
	fmt.Printf("affected rows = %#v", rows)
	db.Table("users").Where("id in (?)", []int{2}).Updates(map[string]interface{}{"name": "hello", "age": 18}) //批量更新
	fmt.Printf("us:%#v\n", us)

	//全量更新
	u1 := User{
		Name: "小王子",
		Age:  18,
	}
	db.Save(&u1)
	u1.Name = "玫瑰"
	db.Debug().Save(&u1) //Debug会打印出SQL语句
	u1.Age = 16
	db.Model(&u1).Where("name = ?", "玫瑰").Save(&u1)

	//使用SQL语句表达更新
	var u2 User
	resp := []User{}
	db.First(&u2)
	db.Model(&u2).Update("age", gorm.Expr("age + ?", 2)).Scan(&resp)
	db.Model(&u2).Updates(map[string]interface{}{"age": gorm.Expr("age - ?", 1)}).Scan(&resp)
	db.Model(&u2).UpdateColumn("age", gorm.Expr("age - ?", 10)).Scan(&resp)
	db.Model(&u2).Where("age < 15 ").UpdateColumn("age", gorm.Expr("age + ?", 15)).Scan(&resp)
	fmt.Printf("resp:%#v\n", resp)

	//删除，注意确保主键字段有值
	var u3 = User{}
	u3.ID = 1
	//UPDATE "users" SET "deleted_at"='2023-12-24 21:05:06.181' WHERE "users"."id" = 1 AND "users"."deleted_at" IS NULL
	db.Debug().Delete(&u3) //仅仅会软删除，记录deleted_at的时间
	// 为删除 SQL 添加额外的 SQL 操作
	//db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&u3)
	//// DELETE from emails where id=10 OPTION (OPTIMIZE FOR UNKNOWN);

	//批量删除:先确定条件再删除
	//db.Debug().Where("name LIKE ? ", "J%").Delete([]User{}) //因为是批量删除，所以Delete()中是Slice
	//db.Debug().Delete([]User{}, "age = ?", 18)

	//软删除后，查询记录，一般查不到
	u4 := []User{}
	//SELECT * FROM "users" WHERE name LIKE '%J' AND "users"."deleted_at" IS NULL
	db.Debug().Where("name LIKE ?", "%J").Find(&u4)
	//SELECT * FROM "users" WHERE name LIKE '%J'
	db.Debug().Unscoped().Where("name LIKE ?", "%J").Find(&u4) //利用Unscoped可以查询被软删除的记录
	fmt.Printf("u4=%v\n", u4)
	//真正的物理删除
	db.Debug().Unscoped().Where("name LIKE ?", "J%").Delete([]User{})
}

func main() {
	test01()

	test02()
}
