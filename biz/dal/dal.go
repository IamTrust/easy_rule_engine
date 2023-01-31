// 初始化数据库配置

package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDefaultDSN 使用 Docker 的话不要改这个, 直接用 docker-compose up 命令即可.
// 不使用 Docker 的话改为自己的数据库即可, 且要用 ./sql/init.sql 创建记录表.
const MySQLDefaultDSN = "gorm:gorm@tcp(10.211.55.7:9910)/gorm?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	err = DB.Find(&Expression{}).Error
	if err != nil {
		panic(err)
	}
}
