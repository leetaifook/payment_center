package mysql

import (
    "database/sql"
    "github.com/mikespook/mymysql/godrv"
    "os"
    "path"
    "payment_center/yaml"
)

var (
    Database string
    PreTable string
    Db       *sql.DB
)

func init() {
    config, _ := yaml.ParseFile(path.Dir(os.Args[0]) + "/config/mysql.config.yaml")
    Database = config["Name"]
    PreTable = config["Pretable"]
    godrv.Register("SET NAMES " + config["Charset"])
    dns := "tcp://" + config["Host"] + ":" + config["Port"] + "/" + config["Name"] + "/" + config["User"] + "/" + config["Password"] //+ "?charset=" + config["Charset"]
    db, err := sql.Open("mymysql", dns)

    if err != nil {
        panic(err)
    }

    Db = db
}
