package sqlserver

import (
	"fmt"

	"github.com/irishconstant/core/contract"
)

// GetDivision возващает Подразделение по его идентификатору
func (s SQLServer) GetDivision(id int) (*contract.Division, error) {
	if id == 0 {
		id = 1
	}

	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name, F_Entity, F_Current_Period, F_Last_Period FROM %s.dbo.Divisions WHERE ID = %d", s.DBName, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Подразделения")
		return nil, err
	}
	defer rows.Close()

	var (
		ID              int
		name            string
		entityID        int
		currentPeriodID int
		lastPeriodID    int
	)
	rows.Scan(
		&ID,
		&name,
		&entityID,
		&currentPeriodID,
		&lastPeriodID)

	if err != nil {
		fmt.Printf("Ошибка с получением Расчётного периода подразделений")
		return nil, err
	}
	division := contract.Division{
		Key:  ID,
		Name: name,
	}

	return &division, nil

}
