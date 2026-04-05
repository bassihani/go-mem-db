package main

import (
	"fmt"
	"storageEngine/internal/db"
)

// workflow:- main -> manager -> engine
func main() {
	engineMap := db.NewEngineMap()
	dbmanagerMap := db.NewDBManager(engineMap)
	// Saving Employee Salary Data just string double
	err := dbmanagerMap.Set("Emp1", 196000)
	if err != nil {
		fmt.Errorf("not able to save value to database %s", err)
	}
	EngineSlice := db.NewEngineMap()
	dbmanagerSlice := db.NewDBManager(EngineSlice)
	// Saving Movie Name Data just string double
	err = dbmanagerSlice.Set("Movie1", "Sully")
	if err != nil {
		fmt.Errorf("not able to save value to database %s", err)
	}
	//Fetch both values and print them
	salary, err := dbmanagerMap.Get("Emp1")
	if err != nil {
		fmt.Errorf("failing to get value from database %s", err)
	}
	fmt.Println("map database stored Value ", salary)
	movie, err := dbmanagerSlice.Get("Movie1")
	if err != nil {
		fmt.Errorf("failing to get value from database %s", err)
	}
	fmt.Println("slice database stored Value ", movie)
}
