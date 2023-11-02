package models

import (
	"api-transferencia/database"
	"fmt"
)

func TransferBalance(idOrigin, idDestiny int, value float64) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceOrigin, err := getBalance(idOrigin)
	if err != nil {
		return err
	}

	if balanceOrigin < value {
		return fmt.Errorf("Saldo insuficiente para a transacao")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	updateBalanceOrigin := `update "user" set balance = balance - $1 where id = $2`
	_, err = tx.Exec(updateBalanceOrigin, value, idOrigin)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateBalanceDestiny := `update "user" set balance = balance + $1 where id = $2`
	_, err = tx.Exec(updateBalanceDestiny, value, idDestiny)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getBalance(id int) (float64, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var balance float64
	query := `select balance from "user" where id = $1`
	err = db.QueryRow(query, id).Scan(&balance)
	if err != nil {
		return 0, nil
	}
	return balance, nil
}
