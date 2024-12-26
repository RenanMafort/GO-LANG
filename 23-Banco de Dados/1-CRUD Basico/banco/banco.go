package banco

import "database/sql"

func Connectar() (*sql.DB, error) {
	stringConexao := "root:cadastro@tcp(192.168.237.171:3306)/devbook?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
