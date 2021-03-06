package mysqldb

import (
	"fmt"
	"testing"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var dbRt db.Database
var odbRt odb.Oauth2DB
var idRt int64

//var cidRti int64

func TestMySQLOauthDBReToken_Connect(t *testing.T) {
	//var db db.Database
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbRt = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 1

	var getRow db.DbRow
	getRow.Row = []string{"1", "somereftoken2"}
	mydb.MockRow1 = &getRow

	mydb.MockUpdateSuccess1 = true

	mydb.MockDeleteSuccess1 = true

	var moadb MySQLOauthDB
	var l lg.Logger
	moadb.Log = &l
	moadb.DB = dbRt

	odbRt = &moadb

	dbRt.Connect()
}

func TestMySQLOauthDBReToken_AddRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.Token = "somereftoken"
	res, id := odbRt.AddRefreshToken(nil, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idRt = id
	}
}

func TestMySQLOauthDBReToken_AddRefreshTokenTx(t *testing.T) {
	var tk odb.RefreshToken
	tk.Token = "somereftoken"

	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockInsertSuccess1 = true
	mdbx.MockInsertID1 = 1
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	var l lg.Logger
	moadbtx.Log = &l
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx

	res, id := odbbUri2TX.AddRefreshToken(&mtx, &tk)
	if !res || id <= 0 {
		t.Fail()
	} else {
		idRt = id
	}
}

func TestMySQLOauthDBReToken_UpdateRefreshToken(t *testing.T) {
	var tk odb.RefreshToken
	tk.ID = idRt
	tk.Token = "somereftoken2"
	res := odbRt.UpdateRefreshToken(&tk)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBReToken_GetRefreshToken(t *testing.T) {
	res := odbRt.GetRefreshToken(idRt)
	fmt.Println("ref token: ", res)
	if res == nil || (*res).Token != "somereftoken2" {
		t.Fail()
	}
}

func TestMySQLOauthDBReToken_DeleteRefreshToken(t *testing.T) {
	res := odbRt.DeleteRefreshToken(nil, idRt)
	fmt.Println("del ref token: ", res)
	if !res {
		t.Fail()
	}
}

func TestMySQLOauthDBReToken_DeleteRefreshTokenTx(t *testing.T) {
	var mtx mdb.MyDbTxMock
	var mdbx mdb.MyDBMock
	mdbx.MockDeleteSuccess1 = true
	mtx.MyDBMock = &mdbx
	var moadbtx MySQLOauthDB
	var l lg.Logger
	moadbtx.Log = &l
	//moadbtx.Tx = &mtx
	var odbbUri2TX = &moadbtx
	res := odbbUri2TX.DeleteRefreshToken(&mtx, idRt)
	fmt.Println("del ref token: ", res)
	if !res {
		t.Fail()
	}
}
