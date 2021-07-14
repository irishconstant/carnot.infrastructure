CREATE TABLE Node_Types (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(300)
	, C_Short_Name NVARCHAR(10)
	, B_Source BIT
	, B_Energy BIT
)

CREATE TABLE Node_Types_Hierarchies(
	F_Parent_Type INT
	, F_Child_Type INT 
	, PRIMARY KEY (F_Parent_Type, F_Child_Type)
)

ALTER TABLE Node_Types_Hierarchies
ADD FOREIGN KEY (F_Parent_Type) REFERENCES Node_Types(ID);
ALTER TABLE Node_Types_Hierarchies
ADD FOREIGN KEY (F_Child_Type) REFERENCES Node_Types(ID);

INSERT INTO Node_Types
(
	C_Name 
	, C_Short_Name 
	, B_Source
	, B_Energy 
)
SELECT 'Коллектор источника тепловой энергии'					, 'ТИ', 1, 1
UNION ALL
SELECT 'Тепловая камера'										, 'ТК', 0, 1
UNION ALL
SELECT 'Врезка в тепловую сеть (без тепловой камеры)'			, 'ВР', 0, 1
UNION ALL
SELECT 'Центральный тепловой пункт'								, 'ЦТП',0, 1
UNION ALL
SELECT 'Индивидуальный тепловой пункт'							, 'ИТП',0, 0
UNION ALL
SELECT 'Ввод в объект потребления'								, 'Ввод',0,0
UNION ALL
SELECT 'Неизвестный участок тепловой сети'						, 'ТС', 0, 0
UNION ALL
SELECT 'Распределительный коллектор'							, 'РК', 0, 1

INSERT INTO Node_Types_Hierarchies(
	F_Parent_Type 
	, F_Child_Type 
)
SELECT 
 3
, 2
UNION ALL
SELECT 
 3
, 3
UNION ALL
SELECT 
 3
, 4
UNION ALL
SELECT 
 3
, 5
UNION ALL
SELECT 
 3
, 6
UNION ALL
SELECT 
 3
, 7
UNION ALL
SELECT 
 3
, 8

/*
SELECT 
 3
, 2
UNION ALL
SELECT 
 3
, 3
UNION ALL
SELECT 
 3
, 4
UNION ALL
SELECT 
 3
, 5
UNION ALL
SELECT 
 3
, 6
UNION ALL
SELECT 
 3
, 7
UNION ALL
SELECT 
 3
, 8
UNION ALL
SELECT 
 2
, 2
UNION ALL
SELECT 
 2
, 3
UNION ALL
SELECT 
 2
, 4
UNION ALL
SELECT 
 2
, 5
UNION ALL
SELECT 
 2
, 6
UNION ALL
SELECT 
 2
, 7
UNION ALL
SELECT 
 2
, 8
UNION ALL
SELECT 
 1
, 2
UNION ALL
SELECT 
 1
, 3
UNION ALL
SELECT 
 1
, 4
UNION ALL
SELECT 
 1
, 5
UNION ALL
SELECT 
 1
, 6
UNION ALL
SELECT 
 1
, 7
UNION ALL
SELECT 
 1
, 8
*/