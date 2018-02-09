package mssql

//MssqlModel blablabla
type MssqlModel struct {
	isdebug  bool
	server   string
	user     string
	password string
	database string
	encrypt  string
	// port int32

}

//Model for mssql
func Model() MssqlModel {

	model := MssqlModel{isdebug: true}
	model.server = "12"
	model.user = "34"
	model.password = "56"
	model.database = "78"
	model.encrypt = "90"

	return model
}

//MS blabla
func MS() (string, int) {

	return "YamiOdymel", 123456
}

//Test ss
var Test string = "This is a test"
