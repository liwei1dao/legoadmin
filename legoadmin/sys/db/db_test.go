package db_test

import (
	"fmt"
	"testing"

	"github.com/liwei1dao/lego/sys/mysql"
)

func TestRa(t *testing.T) {
	if _, err := mysql.NewSys(
		mysql.SetMySQLDsn("root:li13451234@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"),
	); err != nil {
		fmt.Printf("err:%v", err)
		return
	} else {
		// if err = sys.DropTable(comm.TableAgent); err != nil {
		// 	fmt.Printf("err:%v", err)
		// 	return
		// }
		// if err = sys.CreateTable(comm.TableAgent, &pb.DBAgent{}); err != nil {
		// 	fmt.Printf("err:%v", err)
		// 	return
		// }

		// if err = sys.Insert(comm.TableAgent, &pb.DBAgent{
		// 	Agentid:  "1000",
		// 	Agentkey: "D3PX7iaNEZN5FdkG2wfb0w==",
		// 	Currency: "BBL",
		// 	Addrurl:  "https://192.168.110.203:9200/api",
		// }); err != nil {
		// 	fmt.Printf("err:%v", err)
		// 	return
		// }
		// if err = sys.Update(comm.TableGamelist, db.M{"gameid": "dj001"}, db.M{"gamename": "dj001"}); err != nil {
		// 	fmt.Printf("err:%v", err)
		// 	return
		// }
		return
	}
}
