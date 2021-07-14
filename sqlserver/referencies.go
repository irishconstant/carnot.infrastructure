package sqlserver

import (
	"fmt"

	"github.com/irishconstant/core/contract"
	"github.com/irishconstant/core/ref"
)

// GetContactType возвращает Тип контакта по его ИД
func (s SQLServer) GetContactType(id int) (*contract.ContactType, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name, C_Validation_Mask, B_Address  FROM %s.dbo.Contact_Types WHERE ID = %d", s.DBName, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Типа контакта")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID             int
			name           string
			validationMask string
			isAddress      int
		)
		rows.Scan(
			&ID,
			&name,
			&validationMask,
			&isAddress)

		contactType := contract.ContactType{
			Key:        ID,
			Name:       name,
			Validation: validationMask,
			IsAddress:  GetBoolValue(isAddress)}
		return &contactType, nil
	}
	return nil, err
}

// GetAllContactTypes возвращает все возможные типы контактов
func (s SQLServer) GetAllContactTypes() ([]*contract.ContactType, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Contact_Types", s.DBName))
	if err != nil {
		fmt.Printf("Ошибка с получением Типов контакта")
		return nil, err
	}
	defer rows.Close()
	var contactTypes []*contract.ContactType
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID,
		)
		newContactType, _ := s.GetContactType(ID)
		contactTypes = append(contactTypes, newContactType)
	}
	return contactTypes, err
}

// GetCitizenship возвращает гражданство по его ид
func (s SQLServer) GetCitizenship(id int) (*ref.Citizenship, error) {
	if id == 0 {
		id = 1
	}

	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name FROM %s.dbo.Citizenships WHERE ID = %d", s.DBName, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Гражданства")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID   int
			name string
		)
		rows.Scan(
			&ID,
			&name)

		citizenship := ref.Citizenship{
			Key:  ID,
			Name: name}
		return &citizenship, nil
	}
	return nil, err
}

// GetAllCitizenship возвращает все возможные варианты гражданства
func (s SQLServer) GetAllCitizenship() ([]*ref.Citizenship, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Citizenships", s.DBName))
	if err != nil {
		fmt.Printf("Ошибка с получением Гражданств")
		return nil, err
	}
	defer rows.Close()

	var citizenships []*ref.Citizenship
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID)

		newCitizenship, _ := s.GetCitizenship(ID)
		citizenships = append(citizenships, newCitizenship)
	}
	return citizenships, err
}

// GetFuelType возвращает Тип топлива
func (s SQLServer) GetFuelType(id int) (*ref.FuelType, error) {
	if id == 0 {
		id = 1
	}

	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name FROM %s.dbo.Fuel_Types WHERE ID = %d", s.DBName, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Типа топлива")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID   int
			name string
		)
		rows.Scan(
			&ID,
			&name)

		fuelType := ref.FuelType{
			Key:  ID,
			Name: name}
		return &fuelType, nil
	}
	return nil, err
}

// GetAllFuelTypes возвращает все возможные типы топлива
func (s SQLServer) GetAllFuelTypes() ([]*ref.FuelType, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Fuel_Types", s.DBName))
	if err != nil {
		fmt.Printf("Ошибка с получением типов топлива")
		return nil, err
	}
	defer rows.Close()

	var fuelTypes []*ref.FuelType
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID)

		newFuelType, _ := s.GetFuelType(ID)
		fuelTypes = append(fuelTypes, newFuelType)
	}
	return fuelTypes, err
}

// GetSeasonMode возвращает Тип топлива
func (s SQLServer) GetSeasonMode(id int) (*ref.SeasonMode, error) {
	if id == 0 {
		id = 1
	}

	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name FROM %s.dbo.Season_Modes WHERE ID = %d", s.DBName, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Типа топлива")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID   int
			name string
		)
		rows.Scan(
			&ID,
			&name)

		seasonMode := ref.SeasonMode{
			Key:  ID,
			Name: name}
		return &seasonMode, nil
	}
	return nil, err
}

// GetAllSeasonModes возвращает все возможные категории сезонности
func (s SQLServer) GetAllSeasonModes() ([]*ref.SeasonMode, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Season_Modes", s.DBName))
	if err != nil {
		fmt.Printf("Ошибка с получением категорий сезонности")
		return nil, err
	}
	defer rows.Close()

	var seasonModes []*ref.SeasonMode
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID)

		newSeasonMode, _ := s.GetSeasonMode(ID)
		seasonModes = append(seasonModes, newSeasonMode)
	}
	return seasonModes, err
}

//GetCalcPeriod возвращает расчётный период по идентификатору
func (s SQLServer) GetCalcPeriod(id int) (*ref.CalcPeriod, error) {
	if id == 0 {
		id = 1
	}

	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name, N_Year, N_Month, ISNULL(D_Date_Close,'20790606'),  CAST(B_Current AS TINYINT) FROM %s.dbo.Calc_Periods WHERE ID = %d", s.DBName, id))
	if err != nil {
		fmt.Printf("Ошибка с получением Расчётного периода")
		return nil, err
	}
	defer rows.Close()
	var (
		ID        int
		name      string
		year      int
		month     int
		iscurrent int
		dateclose string
	)
	rows.Scan(
		&ID,
		&name,
		&year,
		&month,
		&dateclose,
		&iscurrent,
	)

	calcPeriod := ref.CalcPeriod{
		Key:       ID,
		Name:      name,
		Year:      year,
		Month:     month,
		IsCurrent: GetBoolValue(iscurrent),
	}

	if dateclose != "2079-06-06T00:00:00Z" {
		calcPeriod.DateClose = ConvertSDBToTime(dateclose)
	}
	return &calcPeriod, nil
}

// GetAllCalcPeriods возвращает все возможные расчётные периоды
func (s SQLServer) GetAllCalcPeriods() ([]*ref.CalcPeriod, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID FROM %s.dbo.Calc_Periods  ORDER BY ID DESC", s.DBName))
	if err != nil {
		fmt.Printf("Ошибка с получением расчётных периодов")
		return nil, err
	}
	defer rows.Close()

	var calcPeriods []*ref.CalcPeriod
	for rows.Next() {
		var (
			ID int
		)
		rows.Scan(
			&ID)

		newPeriod, _ := s.GetCalcPeriod(ID)
		calcPeriods = append(calcPeriods, newPeriod)
	}
	return calcPeriods, err
}

//GetCurrentPeriod возвращает расчётный период по идентификатору
func (s SQLServer) GetCurrentPeriod() (*ref.CalcPeriod, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name, N_Year, N_Month FROM %s.dbo.Calc_Periods WHERE B_Current = 1", s.DBName))
	if err != nil {
		fmt.Printf("Ошибка с получением текущего Расчётного периода")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ID    int
			name  string
			year  int
			month int
		)
		rows.Scan(
			&ID,
			&name,
			&year,
			&month,
		)

		calcPeriod := ref.CalcPeriod{
			Key:       ID,
			Name:      name,
			Year:      year,
			Month:     month,
			IsCurrent: true,
		}
		return &calcPeriod, nil
	}
	return nil, err
}
