CREATE TABLE Sources(
	ID 								INT 	IDENTITY(1,1) PRIMARY KEY
	, F_Object 						INT 
	, FOREIGN KEY (F_Object) 				REFERENCES Objects(ID)
	, F_Season_Mode 				TINYINT
	, FOREIGN KEY (F_Season_Mode) 			REFERENCES Season_Modes(ID)
	, F_Fuel_Type 					TINYINT
	, FOREIGN KEY (F_Fuel_Type) 			REFERENCES Fuel_Types(ID)
	, N_Norm_Supply_Value 			MONEY --  Нормативная подпитка тепловых сетей (м3)
	, F_Supplier_Electricity 		INT
	, FOREIGN KEY (F_Supplier_Electricity) REFERENCES Legal_Entities(ID)
	, F_Voltage_Nominal 			TINYINT
	, FOREIGN KEY (F_Voltage_Nominal) 		REFERENCES Voltage_Nominals(ID)
	, F_Transport_Gas        		INT -- Организация-транспортировщик природного газа на котельную
	, F_Supplier_Gas          		INT -- Организация-поставщик природного газа на котельную
	, F_Supplier_TechWater    		INT -- Организация-поставщик воды на технологические нужды котельной
	, F_Supplier_HotWater     		INT -- Организация-поставщик воды на ГВС
	, F_Supplier_Canalisation 		INT -- Организация, оказывающая услугу водоотведения на котельной
	, F_Supplier_Heat         		INT -- Организация-поставщик покупного тепла на котельную (ЦТП)
	, FOREIGN KEY (F_Transport_Gas) 		REFERENCES Legal_Entities(ID)-- Организация-транспортировщик природного газа на котельную
	, FOREIGN KEY (F_Supplier_Gas) 			REFERENCES Legal_Entities(ID) -- Организация-поставщик природного газа на котельную
	, FOREIGN KEY (F_Supplier_TechWater) 	REFERENCES Legal_Entities(ID) -- Организация-поставщик воды на технологические нужды котельной
	, FOREIGN KEY (F_Supplier_HotWater) 	REFERENCES Legal_Entities(ID) -- Организация-поставщик воды на ГВС
	, FOREIGN KEY (F_Supplier_Canalisation) REFERENCES Legal_Entities(ID) -- Организация, оказывающая услугу водоотведения на котельной
	, FOREIGN KEY (F_Supplier_Heat) 		REFERENCES Legal_Entities(ID) -- Организация-поставщик покупного тепла на котельную (ЦТП)
)

CREATE TABLE dbo.Source_Params(
	ID 								INT IDENTITY(1,1) PRIMARY KEY
	, F_Source 						INT
	, FOREIGN KEY (F_Source) 				REFERENCES Sources(ID)
	, N_Month 						TINYINT
	, N_Losses 						MONEY
	, N_Efficiency 					MONEY
)

CREATE TABLE dbo.Source_Facts(
	ID 								INT IDENTITY(1,1) PRIMARY KEY
	, F_Source						INT
	, FOREIGN KEY (F_Source) 				REFERENCES Sources(ID)
	, F_Calc_Period 				TINYINT
	, FOREIGN KEY (F_Calc_Period) 			REFERENCES Calc_Periods(ID)
	, N_Work_Duration           	INT   -- Продолжительность работы источника (в часах)
	, N_Temp_Water              	MONEY -- t°х.воды
	, N_Temp_Air                	MONEY -- t°возд
	, N_Heat_Duration           	INT   -- Отопление, час
	, N_Temp_Heat               	MONEY -- Отопление, град
	, N_Fuel_Consumption        	MONEY -- Расход натурального топлива, тыс.м3, тн
	, N_Electricity_Consumption 	MONEY -- Эл.энергия, тыс. кВт*час
	, N_TechWater_Constumption  	MONEY -- Вода на технологические нужды, тыс. м3
	, N_HotWater_Consumption    	MONEY -- Вода на ГВС, тыс. м3
	, N_Canalisation            	MONEY -- Канализирование, тыс. м3
	, N_PaidHeat                	MONEY -- Покупное тепло, Гкал
)

CREATE TABLE Source_Nodes
(
	ID 								INT IDENTITY(1,1) PRIMARY KEY
	, F_Source						INT
	, FOREIGN KEY (F_Source) 				REFERENCES Sources(ID)
	, F_Node BIGINT
	, FOREIGN KEY (F_Node) 					REFERENCES Nodes (ID)
	, N_Load      					MONEY -- Подключенная нагрузка, Гкал/час:
	, N_HeatLoad  					MONEY -- отопление
	, N_VentLoad  					MONEY -- вентиляция
	, N_HWSLoad   					MONEY -- ГВС
	, N_SteamLoad 					MONEY -- пар
	, F_Temp_Graph_HP  				INT -- Температурный график теплоснабжения в ОП
	, F_Temp_Graph_NHP 				INT -- Температурный график теплоснабжения в МОП
	, isDevice BIT -- Наличие коммерческого узла учета тепловой энергии
	, FOREIGN KEY (F_Temp_Graph_HP  ) 		REFERENCES Temp_Graphs (ID) -- Температурный график теплоснабжения в ОП
	, FOREIGN KEY (F_Temp_Graph_NHP ) 		REFERENCES Temp_Graphs (ID) -- Температурный график теплоснабжения в МОП
)

CREATE TABLE Calc_Periods
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name  NVARCHAR(50)
	, N_Year  INT
	, N_Month INT
)

CREATE TABLE Season_Modes
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(30)
)

CREATE TABLE Fuel_Types
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(30)
)

CREATE TABLE Voltage_Nominals
(
	ID TINYINT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(30)
)