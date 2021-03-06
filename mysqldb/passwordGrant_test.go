package mysqldb

import (
	"fmt"
	"testing"
	"time"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbPg db.Database
var odbPg odb.Oauth2DB
var cidPg int64 = 1
var spIDPg int64
var spID2Pg int64

func TestMySQLOauthDBPg_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()
}

func TestMySQLOauthDBPg_AddPasswordGrant(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.PasswordGrant
	pwg.ClientID = cidPg
	pwg.UserID = "1234"
	res, id := odbPg.AddPasswordGrant(&pwg, &at, &rt)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_AddPasswordGrantFail1(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = false
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.PasswordGrant
	pwg.ClientID = cidPg
	pwg.UserID = "1234"
	res, id := odbPg.AddPasswordGrant(&pwg, &at, &rt)
	if res || id != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_AddPasswordGrantFail2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	var rt odb.RefreshToken
	rt.Token = ""

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.PasswordGrant
	pwg.ClientID = cidPg
	pwg.UserID = "1234"
	res, id := odbPg.AddPasswordGrant(&pwg, &at, &rt)
	if !res || id < 1 {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_AddPasswordGrantFail3(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = false
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = true
	mydb.MockInsertID3 = 1

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.PasswordGrant
	pwg.ClientID = cidPg
	pwg.UserID = "1234"
	res, id := odbPg.AddPasswordGrant(&pwg, &at, &rt)
	if res || id != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_AddPasswordGrantFail4(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	mydb.MockInsertSuccess2 = true
	mydb.MockInsertID2 = 1

	mydb.MockInsertSuccess3 = false
	mydb.MockInsertID3 = 1

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	var rt odb.RefreshToken
	rt.Token = "somereftoken2"

	var at odb.AccessToken
	at.Token = "someacctoken"
	at.Expires = time.Now()

	var pwg odb.PasswordGrant
	pwg.ClientID = cidPg
	pwg.UserID = "1234"
	res, id := odbPg.AddPasswordGrant(&pwg, &at, &rt)
	if res || id != 0 {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_DeletePasswordGrant(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true

	var rows [][]string
	row := []string{"1", "2", "user", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	res := odbPg.DeletePasswordGrant(cidPg, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_DeletePasswordGrant2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true

	var rows [][]string
	row := []string{}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows
	fmt.Println("rows in test", mydb.MockRows1)
	fmt.Println("rows in test len", len(mydb.MockRows1.Rows))

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	res := odbPg.DeletePasswordGrant(cidPg, "1234")
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_DeletePasswordGrantFail1(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = false

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = true

	var rows [][]string
	row := []string{"1", "2", "user", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	res := odbPg.DeletePasswordGrant(cidPg, "1234")
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_DeletePasswordGrantFail2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = false

	mydb.MockDeleteSuccess3 = true

	var rows [][]string
	row := []string{"1", "2", "user", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	res := odbPg.DeletePasswordGrant(cidPg, "1234")
	if res {
		t.Fail()
	}
}

func TestMySQLOauthDBPg_DeletePasswordGrantFail3(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbPg = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockDeleteSuccess1 = true

	mydb.MockDeleteSuccess2 = true

	mydb.MockDeleteSuccess3 = false

	var rows [][]string
	row := []string{"1", "2", "user", "4"}
	rows = append(rows, row)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var tt = time.Now()

	var getRow db.DbRow
	getRow.Row = []string{"1", "test1token", tt.Format("2006-01-02 15:04:05"), "1"}
	mydb.MockRow1 = &getRow

	var getRow2 db.DbRow
	getRow2.Row = []string{"1", "test1refreshtoken"}
	mydb.MockRow2 = &getRow2

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbPg

	odbPg = &moadb

	dbPg.Connect()

	res := odbPg.DeletePasswordGrant(cidPg, "1234")
	if res {
		t.Fail()
	}
}
