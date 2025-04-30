package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(postgres.Open("host=127.0.0.1 user=directus password=directus dbname=directus port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 创建 generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "D:\\develop\\src\\menul-service\\service\\dao",   // 生成到项目里的 "gen" 文件夹
		ModelPkgPath: "D:\\develop\\src\\menul-service\\service\\model", // 生成到项目里的 "gen" 文件夹

	})

	g.UseDB(db)

	user := g.GenerateModel("user")
	g.ApplyBasic(user)

	// 执行生成
	g.Execute()
}
