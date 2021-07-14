CREATE TABLE Nodes
(
	ID BIGINT IDENTITY(1,1) PRIMARY KEY
	, C_Name     NVARCHAR(100)
	, F_Type     INT NOT NULL
	, C_Location NVARCHAR(100) -- В случаях, когда расположение нельзя указать через Объект
	, F_Object   INT
)

ALTER TABLE Nodes
ADD FOREIGN KEY (F_Type) REFERENCES Node_Types(ID);

ALTER TABLE Nodes
ADD FOREIGN KEY (F_Object) REFERENCES Objects(ID);
