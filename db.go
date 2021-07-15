package db

import (
	"github.com/golobby/container"
	"github.com/irishconstant/db/abstract"
	"github.com/irishconstant/db/sqlserver"
)

func InitDBConnection() abstract.DatabaseConnection {
	/*
		Всё началось тогда, когда были выкованы мега-кольца
		Три первых кольца задарили бесссмертным эльфам - чисто для проверки, не передохнут ли
		Семь - коротышкам из подземных канализаций
		Ну а девять колец задарили расе людей. И (как показала практика) напрасно...
	*/
	dbc := GetDependency()
	dbc.GetConnectionParams("config.ini")
	dbc.ConnectToDatabase()

	return dbc
}

//getDependency создаёт привязку между интерфейсом и реализацией (IoC)
func GetDependency() abstract.DatabaseConnection {

	// Если надо изменить реализацию на другую БД, достаточно реализовать её в repo и сослаться на новую реализацию здесь
	container.Singleton(func() abstract.DatabaseConnection {
		return &sqlserver.SQLServer{}
	})
	var dbc abstract.DatabaseConnection
	container.Make(&dbc)
	return dbc
}
