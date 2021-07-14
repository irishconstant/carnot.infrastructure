package sqlserver

import (
	"fmt"
	"strconv"
	"time"

	"github.com/irishconstant/core/contract"
)

// GetEntity возвращает Юридическое лицо
func (s SQLServer) GetEntity(id int) (*contract.LegalEntity, error) {

	rows, err := s.DB.Query(creatorSelect(s.DBName, "Entity", "ID", "ID", strconv.Itoa(id)))

	if err != nil {
		fmt.Println("Ошибка c запросом GetEntity: ", err)
		return nil, err
	}
	defer rows.Close()
	var (
		ID         int
		User       string
		Name       string
		ShortName  string
		INN        string
		KPP        string
		OGRN       string
		EntityType int
		DateReg    string
	)
	rows.Scan(&ID, &User, &Name, &ShortName, &INN, &KPP, &OGRN, &EntityType, &DateReg)
	user, err := s.GetUser(User)
	if err != nil {
		fmt.Println("Ошибка c получением Пользователя: ", err)
		return nil, err
	}

	DateRegG, _ := time.Parse(time.RFC3339, DateReg)

	entity := contract.LegalEntity{
		Key:       ID,
		Name:      Name,
		User:      *user,
		DateReg:   DateRegG,
		ShortName: ShortName,
		INN:       INN,
		KPP:       KPP,
	}
	return &entity, nil
}
