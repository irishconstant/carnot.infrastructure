CREATE TABLE Pipelines
(
	ID BIGINT IDENTITY(1,1)	
	, N_Value TINYINT 
	, F_Edge  BIGINT			-- Линия теплоснабжения
	, N_Year  INT				-- Год прокладки
	, D_Date_Begin DATETIME		-- Дата начала работы (обычно совпадает с годом прокладки)
	, D_Date_End DATETIME		-- Дата окончания работы
	, F_Layer_Type   TINYINT	-- Тип прокладки
	, N_Temp_Project INT		-- Температура проектирования
	, F_Network_Type TINYINT    -- Исполнение сети: однотрубная или двухтрубная
	, F_Diameter_Direct  INT	-- Диаметр подающего трубопровода
	, F_Diamtere_Reverse INT	-- Диаметр обратного трубопровода
	, N_Length_Direct    MONEY  -- Длина трубопровода подающего
	, N_Length_Reverse   MONEY  -- Длина трубопровода обратного
	, F_Calc_Type	     TINYINT -- Способ расчёта темп.коэффициента (ограничить в зависимости от исполнения сети!!!)
	, F_Temp_Graph_HP  INT -- Температурный график (ОП) (ограничить в зависимости от способа расчёта темп.коэффициента)
	, F_Temp_Graph_NHP INT -- Температурный график (МОП)
	, F_Isolation_Type TINYINT -- Теплоизоляционный материал
	, N_Thickness     INT -- Толщина теплоизоляции
	, F_Network_Purposes TINYINT -- Назначение сети
)

CREATE TABLE Layer_Types
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(100)
)

CREATE TABLE Isolation_Types
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(100)
)

CREATE TABLE Calc_Types
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(100)
)

CREATE TABLE Network_Types
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(100)
)

CREATE TABLE Network_Purporses
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(100)
)