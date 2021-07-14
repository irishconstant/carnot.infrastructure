package abstract

import (
	"github.com/irishconstant/core/auth"
	"github.com/irishconstant/core/contract"
	"github.com/irishconstant/core/ref"
	"github.com/irishconstant/core/tech"
)

//DatabaseConnection обеспечивает интерфейс для соединения с СУБД
type DatabaseConnection interface {
	// Работа с соединением БД
	ConnectToDatabase() error
	CloseConnect() error
	GetConnectionParams(filePath string) error

	// Работа со справочниками
	GetContactType(int) (*contract.ContactType, error)
	GetAllContactTypes() ([]*contract.ContactType, error)
	GetCitizenship(int) (*ref.Citizenship, error)
	GetAllCitizenship() ([]*ref.Citizenship, error)
	GetSeasonMode(int) (*ref.SeasonMode, error)
	GetAllSeasonModes() ([]*ref.SeasonMode, error)
	GetFuelType(int) (*ref.FuelType, error)
	GetAllFuelTypes() ([]*ref.FuelType, error)
	GetDivision(id int) (*contract.Division, error)
	GetCalcPeriod(id int) (*ref.CalcPeriod, error)
	GetAllCalcPeriods() ([]*ref.CalcPeriod, error)
	GetCurrentPeriod() (*ref.CalcPeriod, error)

	// Работа с Техническими данными
	GetSource(id int, period *ref.CalcPeriod) (*tech.Source, error)
	GetAllSources(int, int, int, string, string, int, int, *ref.CalcPeriod) (map[int]*tech.Source, error)
	GetSourceQuantityFiltered(u auth.User, name string, address string, seasonMode int, fuelType int, period *ref.CalcPeriod) (int, error)

	// Работа с Юридическими лицами
	GetEntity(id int) (*contract.LegalEntity, error)

	// Работа с Физическими лицами
	GetPersonsFiltered(auth.User, int, int, int, string, string, string, string) (map[int]*contract.Person, error)
	GetPersonQuantityFiltered(auth.User, string, string, string, string) (int, error)
	CreatePerson(*contract.Person) error
	GetPerson(int) (*contract.Person, error)
	UpdatePerson(*contract.Person) error
	DeletePerson(*contract.Person) error

	// Подсистема авторизации и аутентификации
	CreateUser(auth.User) error
	GetUser(string) (*auth.User, error)
	GetAllUsers() ([]auth.User, error)
	CheckPassword(string, b string) bool
	GetUserRoles(*auth.User) error
	GetAllRoles() (map[int]*auth.Role, error)
	GetRoleAbilities(*auth.Role) error
	GetRoleByID(int) (*auth.Role, error)
	GetUserAttributes(*auth.User) error
}
