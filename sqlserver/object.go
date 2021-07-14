package sqlserver

import (
	"fmt"
	"strconv"

	"github.com/irishconstant/core/tech"
)

// GetObject возвращает Объект с подобъектами
func (s SQLServer) GetObject(id int) (*tech.Object, error) {

	rows, err := s.DB.Query(creatorSelect(s.DBName, "Object", "ID", "ID", strconv.Itoa(id)))
	if err != nil {
		fmt.Println("Ошибка c запросом в GetObject: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ID       int
			Name     string
			Address  string
			Region   string
			District string
			City     string
			Town     string
			Street   string
			House    string
			Building string
			FIAS     string
		)
		rows.Scan(&ID, &Name, &Address, &Region, &District, &City, &Town, &Street, &House, &Building, &FIAS)

		object := tech.Object{
			Key:          ID,
			Name:         Name,
			BuildAddress: Address}

		return &object, nil
	}

	return nil, err
}
