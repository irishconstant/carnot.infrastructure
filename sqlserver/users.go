package sqlserver

import (
	"fmt"
	"strconv"

	"github.com/irishconstant/core/auth"
)

//CreateUser создает нового Пользователя с ролью
func (s *SQLServer) CreateUser(u auth.User) error {
	// Создаём самого пользователя
	hashedPassword, _ := HashPassword(u.Password)

	query := fmt.Sprintf("INSERT INTO %s.dbo.Users (Login, Password, C_Family_Name, C_Name) SELECT '%s', '%s', '%s', '%s'",
		s.DBName, u.Key, hashedPassword, u.FamilyName, u.Name)
	_, err := s.DB.Query(query)
	if err != nil {
		fmt.Println("Ошибка c запросом в CreateUserWithRoles: ", query, "Ошибка", err)
		return err
	}

	// Создаём его роль (сделано через отдельную таблицу. Т.к. заложено на будущее. что ролей у одного Пользователя будет много)
	query = fmt.Sprintf("INSERT INTO %s.dbo.User_Roles (F_Users, F_Roles) SELECT '%s', %s",
		s.DBName, u.Key, strconv.FormatInt(int64(u.Role.Key), 10))
	_, err = s.DB.Query(query)
	if err != nil {
		fmt.Println("Ошибка c запросом в CreateUserWithRoles: ", query, "Ошибка", err)
		return err
	}

	return nil
}

// GetUser получает данные Пользователя из БД
func (s SQLServer) GetUser(login string) (*auth.User, error) {
	user := auth.User{Key: login}
	s.GetUserAttributes(&user)
	s.GetUserRoles(&user)

	return &user, nil
}

// GetUserRoles возвращает все роли Пользователя из БД
func (s *SQLServer) GetUserRoles(user *auth.User) error {

	rows, err := s.DB.Query(fmt.Sprintf("SELECT TOP 1 r.ID, r.C_Name FROM [%s].dbo.User_Roles AS ur INNER JOIN [%s].dbo.Roles AS r ON r.ID = ur.F_Roles  WHERE ur.F_Users = '%s'",
		s.DBName, s.DBName, user.Key))

	if err != nil {
		fmt.Println("Ошибка c запросом в GetUserRoles: ", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			a int
			b string
		)
		rows.Scan(&a, &b)
		role := auth.Role{Key: a, Name: b}
		user.Role = &role
	}

	// Получаем возможности для роли Пользователя
	s.GetRoleAbilities(user.Role)

	return nil
}

// GetUserAttributes выдает атрибуты пользователя из БД
func (s SQLServer) GetUserAttributes(user *auth.User) error {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT u.C_Name, u.C_Family_Name FROM [%s].dbo.Users AS u WHERE u.Login = '%s'",
		s.DBName, user.Key))

	if err != nil {
		fmt.Println("Ошибка c запросом в GetUserAttributes: ", err)
		return err
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Name, &user.FamilyName)
	}

	return nil
}

//GetAllUsers возвращает всех Пользователей
func (s *SQLServer) GetAllUsers() ([]auth.User, error) {
	rows, err := s.DB.Query(fmt.Sprintf("SELECT Login FROM %s.dbo.Users WHERE Login != ''", s.DBName))

	if err != nil {
		fmt.Println("Ошибка c запросом в GetAllUsers: ", err)
		return nil, err
	}
	var users []auth.User

	defer rows.Close()

	for rows.Next() {
		var (
			login string
		)
		rows.Scan(&login)

		user, err := s.GetUser(login)
		if err != nil {
			fmt.Println("Ошибка c запросом в GetAllUsers: ", err)
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil
}
