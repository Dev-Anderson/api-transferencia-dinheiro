package usecase

import (
	"api-transferencia/database"
	"database/sql"
	"fmt"
)

func verifyTable(db *sql.DB, tableName string) (bool, error) {
	query := "select tablename from pg_catalog.pg_tables where tablename = $1"

	var existingTableName string
	err := db.QueryRow(query, tableName).Scan(&existingTableName)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func isFKExists(db *sql.DB, tableName, fk string) (bool, error) {
	query := `
		select count(*)
		from information_schema.table_constraints
		where table_name = $1
		and constraint_name = $2
		and constraint_type = 'FOREIGN KEY'
	`

	var count int
	err := db.QueryRow(query, tableName, fk).Scan(&count)
	if err != nil {
		return false, nil
	}
	return count > 0, nil
}

func TableUser() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tableExists, err := verifyTable(db, "user")
	if tableExists == false {
		createTableUser := `
		CREATE TABLE "user" (
			"id" serial PRIMARY KEY,
			"name" text,
			"document" text,
			"dateCreate" timestamp,
			"balance" numeric, 
			"dateAtualization" timestamp
		  );`

		_, err := db.Exec(createTableUser)
		if err != nil {
			fmt.Println("Erro created table user ", err.Error())
		}
		fmt.Println("Table user created successfully")
	} else {
		fmt.Println("Table user already exists")
	}
	if err != nil {
		fmt.Println("Error in verify table user ", err.Error())
	}
}

func TableBalance() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tableExists, err := verifyTable(db, "balance")
	if tableExists == false {
		createBalance := `
			CREATE TABLE "balance" (
				"id" serial PRIMARY KEY,
				"id_user_origin" int,
				"id_user_destiny" int,
				"value" numeric,
				"dateBalance" timestamp
		  	);`
		_, err := db.Exec(createBalance)
		if err != nil {
			fmt.Println("Error created table balance ", err.Error())
		}
		fmt.Println("Table balance created successfully")
	} else {
		fmt.Println("Table balance already exists")
	}
	if err != nil {
		fmt.Println("Error in verify table balance", err.Error())
	}
}

func FKBalanceIDUserOrigin() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `ALTER TABLE "balance" ADD FOREIGN KEY ("id_user_origin") REFERENCES "user" ("id");`

	fkExists, err := isFKExists(db, "balance", "balance_id_user_origin_fkey")
	if fkExists == false {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println("Error created FK name balance_id_user_origin in table balance ", err.Error())
		}
		fmt.Println("FK name balance_id_user_origin created sucessfully")
	} else if err != nil {
		fmt.Println("Failed to create the FK in table balance and field id_user_origin ", err)
	} else {
		fmt.Println("FK name balance_id_user_origin in table balance already exists")
	}
}

func FKBalanceIDUserDestiny() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `ALTER TABLE "balance" ADD FOREIGN KEY ("id_user_destiny") REFERENCES "user" ("id");`

	fkExists, err := isFKExists(db, "balance", "balance_id_user_destiny")
	if fkExists == false {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println("Error created FK name balance_id_user_destiny in table balance ", err.Error())
		}
		fmt.Println("FK name balance_id_user_destiny created sucessfully")
	} else if err != nil {
		fmt.Println("Failed to creat the FK in table balance id_user_destiny", err)
	} else {
		fmt.Println("FK name balance_id_user_destiny in table balance already exists")
	}
}
