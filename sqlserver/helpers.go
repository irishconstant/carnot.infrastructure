package sqlserver

import (
	"fmt"
	"time"
)

//creatorSelect возвращает базовый запрос для получения данных с минимальной фильтрацией
func creatorSelect(databaseName string, sourceName string, orderParam string, whereParam string, whereValue string) string {
	switch sourceName {
	//  Person
	case "Person":
		return fmt.Sprintf("SELECT [ID], [C_Family_Name], [C_Name], [C_Patronymic_Name], [F_Users], ISNULL([F_Citizenship],0), CAST(B_Sex AS TINYINT), [D_Date_Birth], [D_Date_Death] FROM [%s].dbo.Persons WHERE %s = %s",
			databaseName, whereParam, whereValue)
	// Contact
	case "Contact":
		return fmt.Sprintf("SELECT [ID], [F_Contact_Type], [C_Value], [B_Primary] FROM [%s].dbo.Contacts WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "ContactList":
		return fmt.Sprintf("SELECT [ID] FROM %s.dbo.Contacts WHERE  %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	// ContactType
	case "ContactType":
		return fmt.Sprintf("SELECT [ID], [C_Name] FROM [%s].dbo.Contact_Types WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "ContactTypeList":
		return fmt.Sprintf("SELECT [ID] FROM [%s].dbo.Contact_Types",
			databaseName)
	// Documents
	case "Document":
		return fmt.Sprintf("SELECT [ID], [C_Serial_Number], [C_Number], [C_From_Name], [C_From_Code], [D_Date_Begin], [D_Date_End] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	case "DocumentList":
		return fmt.Sprintf("SELECT [ID] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	// DocumentType
	case "DocumentTypeList":
		return fmt.Sprintf("SELECT [ID], [C_Serial_Number], [C_Number], [C_From_Name], [C_From_Code], [D_Date_Begin], [D_Date_End] FROM [%s].dbo.Documents WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	// Citizenship
	case "CitizenshipList":
		return fmt.Sprintf("SELECT [ID], [C_Name] FROM [%s].[dbo].[Citizenships] WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	// Source
	case "Source":
		return fmt.Sprintf("SELECT [ID], [C_Name], [F_Object], [F_Season_Mode], [F_Fuel_Type], [N_Norm_Supply_Value], [F_Supplier_Electricity], [F_Voltage_Nominal], [F_Transport_Gas]"+
			", [F_Supplier_Gas], [F_Supplier_TechWater], [F_Supplier_HotWater], [F_Supplier_Canalisation], [F_Supplier_Heat], [F_Division] FROM [%s].dbo.Sources WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "SourceList":
		return fmt.Sprintf("SELECT [ID]"+
			" FROM [%s].dbo.Sources WHERE %s = %s",
			databaseName, whereParam, whereValue)
	// CalcPeriod
	case "CalcPeriod":
		return fmt.Sprintf("SELECT  [ID], [C_Name], [N_Year], [N_Month], [B_Current], [D_Date_Close] FROM [%s].[dbo].[Calc_Periods] WHERE %s = %s ORDER BY %s",
			databaseName, whereParam, whereValue, orderParam)
	// Entity
	case "Entity":
		return fmt.Sprintf("SELECT [ID], [F_User], [C_Name], [C_Short_Name], [INN], [KPP], [OGRN], [F_Entity_Type], [D_Date_Reg] "+
			"FROM [%s].dbo.Legal_Entities WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "EntityList":
		return fmt.Sprintf("SELECT [ID]"+
			" FROM [%s].dbo.Legal_Entities WHERE %s = %s",
			databaseName, whereParam, whereValue)
	// EntityType
	case "EntityType":
		return fmt.Sprintf("SELECT [ID], [F_User], [C_Name], [C_Short_Name], [INN], [KPP], [OGRN], [F_Entity_Type], [D_Date_Reg]"+
			" FROM [%s].dbo.Entity_Types WHERE %s = %s",
			databaseName, whereParam, whereValue)
	case "EntitTypeList":
		return fmt.Sprintf("SELECT [ID]"+
			" FROM [%s].dbo.Entity_Types WHERE %s = %s",
			databaseName, whereParam, whereValue)
	// Object
	case "Object":
		return fmt.Sprintf("SELECT [ID], [C_Name], [C_Address], [F_Region], [F_District], [F_City], [F_Town], [F_Street], [C_House], [C_Building], [U_FIAS]"+
			" FROM [%s].dbo.Objects WHERE %s = %s",
			databaseName, whereParam, whereValue)
	}
	return ""
}

// GetBoolValue конвертирует дату из формата db (возвращается как int) в формат golang bool
func GetBoolValue(i int) bool {

	return i == 1
}

// ConvertDate конвертирует дату из формата golang time.Time в формат для записи в БД
func ConvertDate(d time.Time) string {
	return d.Format("2006.01.02 15:04:05")
}

//ConvertSDBToTime конвертирует дату из формата sqlserver в формат golang time.time
func ConvertSDBToTime(s string) time.Time {
	date, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t := new(time.Time)
		return *t
	}
	return date
}
