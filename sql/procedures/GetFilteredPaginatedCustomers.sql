SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		ai
-- Create date: 2020.12.05
-- Alter date: 2020.12.07 Изменения связанные с переименованием таблицы и оптимизации для представлений
-- Description:	Возвращает отфильтрованных Потребителей в пределах одной страницы
-- =============================================
ALTER PROCEDURE [dbo].[GetFilteredPaginatedPersons] -- EXEC dbo.GetFilteredPaginatedPersons 'rode_orm', '', '', '', true, 0, 7, 1
	@Login NVARCHAR(100)
	, @FamilyName NVARCHAR(100) 
	, @Name NVARCHAR(100)
	, @PatrName NVARCHAR(100)
	, @Sex BIT
	, @OffSetRow INT 
	, @PageSize  INT
	, @Regime INT
WITH EXECUTE AS OWNER
AS
BEGIN
	SET NOCOUNT ON;

	IF @Regime = 1 
	BEGIN
		SELECT p.[ID], p.[C_Family_Name], p.[C_Name], p.[C_Patronymic_Name], p.[F_Users], c.[C_Name] AS [Cit_Name], p.[B_Sex], p.[D_Date_Birth], p.[D_Date_Death]
		FROM dbo.Persons  AS p
			INNER JOIN dbo.Citizenships AS c ON c.ID = p.[F_Citizenship]
		WHERE	p.F_Users = @Login
				AND p.C_Family_Name LIKE CONCAT('%',ISNULL(@FamilyName, p.C_Family_Name),'%')
				AND p.C_Name LIKE CONCAT('%',ISNULL(@Name, p.C_Name),'%') 
				AND p.C_Patronymic_Name LIKE CONCAT('%',ISNULL(@PatrName, p.C_Patronymic_Name),'%') 
				AND p.B_Sex = ISNULL(@Sex, B_Sex)
		ORDER BY 1,2
		OFFSET @OffSetRow ROWS 
		FETCH NEXT @PageSize ROWS ONLY
	END
	ELSE
	BEGIN
	SELECT p.[ID], p.[C_Family_Name], p.[C_Name], p.[C_Patronymic_Name], p.[F_Users], c.[C_Name] AS [Cit_Name], p.[B_Sex], p.[D_Date_Birth], p.[D_Date_Death]
		FROM dbo.Persons AS p
			INNER JOIN dbo.Citizenships AS c ON c.ID = p.[F_Citizenship]
		WHERE	p.F_Users = @Login
				AND p.C_Family_Name LIKE CONCAT('%',ISNULL(@FamilyName, p.C_Family_Name),'%')
				AND p.C_Name LIKE CONCAT('%',ISNULL(@Name, p.C_Name),'%') 
				AND p.C_Patronymic_Name LIKE CONCAT('%',ISNULL(@PatrName,p.C_Patronymic_Name),'%') 
				AND B_Sex = ISNULL(@Sex, B_Sex)

		ORDER BY 1,2
	END
END