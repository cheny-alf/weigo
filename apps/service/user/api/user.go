package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"wego/apps/service/user/api/internal/config"
	"wego/apps/service/user/api/internal/handler"
	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/model/ent"
	"wego/apps/service/user/model/ent/migrate"
)

var configFile = flag.String("f", "/Users/cheny/GolandProjects/githubProject/weigo/apps/service/user/api/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	client := initDb()

	ctx := svc.NewServiceContext(c, client)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func initDb() *ent.Client {
	var err error
	client, err := ent.Open("mysql", "root:128568chen@tcp(localhost:3306)/mydb")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	// defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
