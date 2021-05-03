package employees

import (
	config "northwindApi/config"
)

type EmployeeModel struct {
}

//GET
func (employeeModel EmployeeModel) FindAll() ([]Employee, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return nil, err
	} else {
		var employees []Employee
		db.Find(&employees)
		return employees, nil
	}

}

//GET{ID}
func (employeeModel EmployeeModel) FindByID(id int) (Employee, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return Employee{}, err
	} else {
		var employees Employee
		db.Where("EmployeeID = ?", id).Find(&employees) //Es literal el nombre de la columna
		return employees, nil
	}

}
