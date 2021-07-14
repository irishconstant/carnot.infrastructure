
CREATE TABLE Temps (
	ID INT IDENTITY(1,1) PRIMARY KEY,
	F_Location	INT,
	F_Node		BIGINT,  
	N_Year		TINYINT
)

ALTER TABLE Temps
ADD FOREIGN KEY (F_Location) REFERENCES Locations(ID)

ALTER TABLE Temps
ADD FOREIGN KEY (F_Node) REFERENCES Nodes(ID)

CREATE TABLE SNIP_Temps (
	ID INT IDENTITY(1,1) PRIMARY KEY,
	F_Location	INT,
	N_Year		TINYINT
)

ALTER TABLE SNIP_Temps
ADD FOREIGN KEY (F_Location) REFERENCES Locations(ID)

CREATE TABLE Temp_Month_Values (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, N_Month       tinyint
	, F_Temp		int
	, N_Air_Temp    int
	, N_Ground_Temp int
	, N_Water_Temp  int
	, N_Heat_Days   int
)

ALTER TABLE Temp_Month_Values
ADD FOREIGN KEY (F_Temp) REFERENCES Temps(ID)

CREATE TABLE SNIP_Temp_Month_Values (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, N_Month       tinyint
	, F_Temp_SNIP	int
	, N_Air_Temp    int
	, N_Ground_Temp int
	, N_Water_Temp  int
	, N_Heat_Days   int
)


ALTER TABLE SNIP_Temp_Month_Values
ADD FOREIGN KEY (F_Temp_SNIP) REFERENCES SNIP_Temps(ID)

CREATE TABLE Temp_Daily_Values (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, F_Temp		int
	, N_Month		tinyint
	, N_Day			tinyint
	, N_Air_Temp    int
	, N_Ground_Temp int
	, N_Water_Temp  int
	, N_Heat_Days   int
)

ALTER TABLE Temp_Daily_Values
ADD FOREIGN KEY (F_Temp) REFERENCES Temps(ID)

CREATE TABLE SNIP_Temp_Daily_Values (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, F_Temp_SNIP	int
	, N_Month		tinyint
	, N_Day			tinyint
	, N_Air_Temp    int
	, N_Ground_Temp int
	, N_Water_Temp  int
	, N_Heat_Days   int
)

ALTER TABLE SNIP_Temp_Daily_Values
ADD FOREIGN KEY (F_Temp_SNIP) REFERENCES SNIP_Temps(ID)

CREATE TABLE Temp_Graphs (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, C_Name NVARCHAR(100)
)

CREATE TABLE Temp_Graph_Values (
	ID INT IDENTITY(1,1) PRIMARY KEY
	, F_Temp_Graph INT
	, N_Air	INT
	, N_Dir MONEY
	, N_Rev MONEY
	, N_Heat MONEY
	, B_Cut  BIT
)

ALTER TABLE Temp_Graph_Values
ADD FOREIGN KEY (F_Temp_Graph) REFERENCES Temp_Graphs(ID)

