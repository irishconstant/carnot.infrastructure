package sqlserver

import (
	"fmt"
	"strconv"

	"github.com/irishconstant/core/auth"

	"golang.org/x/crypto/bcrypt"
)

//CheckPassword проверяет пароль
func (s *SQLServer) CheckPassword(login string, password string) bool {

	rows, err := s.DB.Query(fmt.Sprintf("SELECT TOP 1 Password FROM %s.dbo.Users WHERE Login = '%s'", s.DBName, login))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false
	}
	defer rows.Close()
	var passwordDB string
	for rows.Next() {
		rows.Scan(&passwordDB)
	}

	if CheckPasswordHash(password, passwordDB) {
		fmt.Printf("Успешная аутентификация пользователя %s", login)
		return true
	}
	//	fmt.Println(fmt.Sprintf("Провалена аутентификация пользователя %s", login))
	return false
}

//HashPassword хэширует пароль
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash проверят соответствие пароля и хэша
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetPossibleRoles возвращает роли, которые может присваивать Пользователь с определенной ролью (неавторизованный = Гость)
func (s *SQLServer) GetPossibleRoles(auth.Role) map[int]auth.Role {
	var roles = make(map[int]auth.Role) ///!!!
	return roles                        ///!!!
}

//GetAllRoles возвращает все роли, возможные в Системе
func (s *SQLServer) GetAllRoles() (map[int]*auth.Role, error) {
	var roles = make(map[int]*auth.Role)

	rows, err := s.DB.Query(fmt.Sprintf("SELECT ID, C_Name, B_Admin_Only FROM %s.dbo.Roles", s.DBName))
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Key       int
			Name      string
			AdminOnly bool
		)

		rows.Scan(&Key, &Name, &AdminOnly)
		role := auth.Role{
			Key:       Key,
			Name:      Name,
			AdminOnly: AdminOnly}
		if Key != 0 {
			roles[Key] = &role
		}
	}
	return roles, nil
}

//GetRoleByID возвращает роль и её возможности по идентификатору
func (s *SQLServer) GetRoleByID(id int) (*auth.Role, error) {
	// Получаем роль из БД
	rows, err := s.DB.Query(fmt.Sprintf("SELECT r.ID, r.C_Name FROM [%s].dbo.Roles AS r WHERE r.ID = %s",
		s.DBName, strconv.FormatInt(int64(id), 10)))

	var role auth.Role

	if err != nil {
		fmt.Println("Ошибка c запросом в GetRoleByID: ", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			a int
			b string
		)
		rows.Scan(&a, &b)
		role = auth.Role{Key: a, Name: b}
	}

	s.GetRoleAbilities(&role)

	return &role, nil
}

// GetRoleAbilities получает данные о возможностях роли
func (s SQLServer) GetRoleAbilities(role *auth.Role) error {

	rows, err := s.DB.Query(fmt.Sprintf("SELECT a.ID, a.C_Name, ar.B_Create, ar.B_Read, ar.B_Update, ar.B_Delete FROM [%s].dbo.Area_Roles AS ar INNER JOIN [%s].dbo.Areas AS a ON a.ID = ar.F_Areas WHERE ar.F_Roles = %s",
		s.DBName, s.DBName, strconv.FormatInt(int64(role.Key), 10)))

	if err != nil {
		fmt.Println("Ошибка c получение возможностей роли: ", err)
		return err
	}
	defer rows.Close()

	createMap := make(map[int]*auth.Area)
	readMap := make(map[int]*auth.Area)
	updateMap := make(map[int]*auth.Area)
	deleteMap := make(map[int]*auth.Area)

	for rows.Next() {
		var (
			ID     int
			name   string
			create bool
			read   bool
			update bool
			delete bool
		)
		rows.Scan(&ID, &name, &create, &read, &update, &delete)

		area := auth.Area{
			Key:  ID,
			Name: name,
		}

		// Код ниже выглядит странно, но там приходится делать, т.к. при изначальной инициализации структуры создаётся nil мапа
		if create {
			createMap[area.Key] = &area
		}
		if read {
			readMap[area.Key] = &area
		}
		if update {
			updateMap[area.Key] = &area
		}
		if delete {
			deleteMap[area.Key] = &area
		}

	}
	role.CreateAbility = createMap
	role.ReadAbility = readMap
	role.UpdateAbility = updateMap
	role.DeleteAbility = deleteMap

	return nil
}
