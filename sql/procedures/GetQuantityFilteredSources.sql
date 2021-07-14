USE [Administratum]
GO
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		ai
-- Create date: 2021.01.01
-- Alter date:	2021.01.03; 2021.01.02
-- Description:	Возвращает отфильтрованные источники тепло-электро-водоснабжения в пределах одной страницы
-- =============================================
ALTER PROCEDURE [dbo].[GetQuantityFilteredSources]	-- EXEC  [dbo].[GetQuantityFilteredSources] NULL, 'Казачья', NULL, NULL, NULL, NULL, NULL, NULL,  0
	@C_Name nvarchar(1000)			   NULL
	, @C_Address nvarchar(1000)		   NULL
	, @F_Season_Mode tinyint 		   NULL
	, @F_Fuel_Type tinyint 			   NULL
	, @N_Norm_Supply_Value_Min money   NULL
	, @N_Norm_Supply_Value_Max money   NULL
	/*
	, @F_Supplier_Electricity int 	   NULL
	, @F_Voltage_Nominal tinyint 	   NULL
	, @F_Transport_Gas int 			   NULL
	, @F_Supplier_Gas int 			   NULL
	, @F_Supplier_TechWater int 	   NULL
	, @F_Supplier_HotWater int 		   NULL
	, @F_Supplier_Canalisation int 	   NULL
	, @F_Supplier_Heat int 			   NULL
	*/
	, @OffSetRow INT 				   NULL
	, @PageSize  INT				   NULL
	, @Regime INT					   NULL
WITH EXECUTE AS OWNER
AS
BEGIN
	SET NOCOUNT ON;
	SELECT COUNT(*)
	FROM [dbo].[Sources] AS s
			LEFT JOIN [dbo].[Objects] AS o ON s.F_Object = o.ID
		WHERE	(s.C_Name LIKE CONCAT('%', @C_Name, '%') OR @C_Name IS NULL)
				AND (o.C_Address LIKE CONCAT('%', @C_Address, '%') OR @C_Address IS NULL)
				AND (F_Season_Mode = @F_Season_Mode OR ISNULL(@F_Season_Mode, NULL) = 0) 
				AND (F_Fuel_Type = @F_Fuel_Type OR ISNULL(@F_Fuel_Type, NULL) = 0) 
				AND (N_Norm_Supply_Value >= @N_Norm_Supply_Value_Min OR @N_Norm_Supply_Value_Min IS NULL) 
				AND (N_Norm_Supply_Value <= @N_Norm_Supply_Value_Max OR @N_Norm_Supply_Value_Max IS NULL) 
	
END
