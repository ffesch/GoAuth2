package managers

import (
	"testing"

	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"

	msdb "github.com/Ulbora/GoAuth2/mysqldb"
	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

func TestOauthManagerClientRoles_AddClientRole(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	mydb.MockInsertID1 = 2
	mydb.MockInsertSuccess1 = true

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	man.Db = odbAu
	var m Manager
	m = &man

	var cr ClientRole
	cr.Role = "tester"
	cr.ClientID = 3
	suc, id := m.AddClientRole(&cr)
	if !suc || id != 2 {
		t.Fail()
	}
}

func TestOauthManagerClientRoles_GetClientRole(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	// mydb.MockInsertID1 = 2
	// mydb.MockInsertSuccess1 = true

	var rows [][]string
	row1 := []string{"1", "role1", "2"}
	rows = append(rows, row1)
	row2 := []string{"2", "role2", "2"}
	rows = append(rows, row2)
	var dbrows db.DbRows
	dbrows.Rows = rows
	mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	man.Db = odbAu
	var m Manager
	m = &man

	rl := m.GetClientRoleList(2)
	if len(*rl) != 2 || (*rl)[1].ID != 2 {
		t.Fail()
	}
}

func TestOauthManagerClientRoles_DeleteClientRole(t *testing.T) {

	var dbAu db.Database
	var odbAu odb.Oauth2DB
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "ulbora_oauth2_server"
	dbAu = &mydb

	var mTestRow db.DbRow
	mTestRow.Row = []string{}
	mydb.MockTestRow = &mTestRow

	//mydb.MockInsertID1 = 2
	mydb.MockDeleteSuccess1 = true

	// var rows [][]string
	// row1 := []string{"1", "code", "2"}
	// rows = append(rows, row1)
	// var dbrows db.DbRows
	// dbrows.Rows = rows
	// mydb.MockRows1 = &dbrows

	var moadb msdb.MySQLOauthDB
	moadb.DB = dbAu

	odbAu = &moadb

	var man OauthManager
	man.Db = odbAu
	var m Manager
	m = &man

	suc := m.DeleteClientRole(2)
	if !suc {
		t.Fail()
	}
}
