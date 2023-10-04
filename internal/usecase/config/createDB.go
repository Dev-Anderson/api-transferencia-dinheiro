package config

import (
	"api-transferencia/pkg/database"
	"database/sql"
	"fmt"
)

func verifyTable(db *sql.DB, tableName string) (bool, error) {
	query := "SELECT tablename FROM pg_catalog.pg_tables WHERE tablename = $1"

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

func isFKExists(db *sql.DB, tableName, FK string) (bool, error) {
	query := `
	SELECT COUNT(*)
		FROM information_schema.table_constraints
		WHERE table_name = $1
		  AND constraint_name = $2
		  AND constraint_type = 'FOREIGN KEY'
	`

	var count int
	err := db.QueryRow(query, tableName, FK).Scan(&count)
	if err != nil {
		return false, nil
	}
	return count > 0, nil

}

func TableUser() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Error connect database", err.Error())
	}
	defer db.Close()

	tableExists, err := verifyTable(db, "user_balance")
	if tableExists == false {
		createTable := `
		CREATE TABLE "user_balance" (
			"id" serial PRIMARY KEY,
			"name" text,
			"document" text,
			"dateCreate" timestamp,
			"dateDelete" timestamp
		  );`

		_, err := db.Exec(createTable)
		if err != nil {
			fmt.Println("Error created table user_balance", err.Error())
		}
		fmt.Println("Table user_balance created successfully")
	} else {
		fmt.Println("Table user_balance already exists")
	}
}

func TableExtractUser() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Error connect database", err.Error())
	}
	defer db.Close()

	tableExists, err := verifyTable(db, "extract_user")
	if tableExists == false {
		createTable := `
		CREATE TABLE "extract_user" (
			"id" serial PRIMARY KEY,
			"typeOperation" text,
			"idUserOrigem" int,
			"valueOrigem" numeric(15,2),
			"idUserDestiny" int,
			"valueDestiny" numeric(15,2),
			"dateInsert" timestamp
		  );`

		_, err := db.Exec(createTable)
		if err != nil {
			fmt.Println("Error created table extract_user", err.Error())
		}
		fmt.Println("table extract_user created successfully")
	} else {
		fmt.Println("Table extract_user already exists")
	}
}

func TableBalance() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Error connect database", err.Error())
	}
	defer db.Close()

	tableExists, err := verifyTable(db, "balance")
	if tableExists == false {
		createTable := `
			CREATE TABLE "balance" (
				"id" serial,
				"idUser" int,
				"value" numeric(15,2)
			);`

		_, err := db.Exec(createTable)
		if err != nil {
			fmt.Println("Error created table balance", err.Error())
		}
		fmt.Println("Table balance created successfully")
	} else {
		fmt.Println("Table balance already exists")
	}
}

func FKExtractUserIDOrigem() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Erro connection database", err.Error())
	}
	defer db.Close()

	sqlStatement := `ALTER TABLE "extract_user" ADD FOREIGN KEY ("idUserOrigem") REFERENCES "user_balance" ("id");`

	fkExists, err := isFKExists(db, "extract_user", "extract_user_idUserOrigem_fkey")
	if fkExists == false {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println("Error created FK in table extract_user", err.Error())
		}
		fmt.Println("FK created Successfully")
	} else {
		fmt.Println("FK in table extract_user already exists")
	}
}

func FKExtractUserIDDestiny() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Error connection database", err.Error())
	}
	defer db.Close()

	sqlStatement := `ALTER TABLE "extract_user" ADD FOREIGN KEY ("idUserDestiny") REFERENCES "user_balance" ("id");`

	fkExists, err := isFKExists(db, "extract_user", "extract_user_idUserDestiny_fkey")
	if fkExists == false {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println("Error created FK in table extract_user", err.Error())
		}
		fmt.Println("FK created sucessfully")
	} else {
		fmt.Println("FK in table extract_user already exists")
	}
}

func FKBalanceIDUser() {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Error connection database", err.Error())
	}
	defer db.Close()

	sqlStatement := `ALTER TABLE "balance" ADD FOREIGN KEY ("idUser") REFERENCES "user_balance" ("id");`

	fkExists, err := isFKExists(db, "balance", "balance_idUser_fkey")
	if fkExists == false {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println("Erorr created FK in table balance", err.Error())
		}
		fmt.Println("FK created sucessfully")
	} else {
		fmt.Println("FK in table balance already exists")
	}
}
