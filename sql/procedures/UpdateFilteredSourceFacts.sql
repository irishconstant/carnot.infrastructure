-- =============================================
-- Author:		ai
-- Create date: 2021.01.17
-- Alter date:	
-- Description:	Обновляет/вставляет фактические данные по источникам с учётом фильтров
-- =============================================
ALTER PROCEDURE [dbo].[UpdateFilteredSourceFacts]	
	@C_Name NVARCHAR(1000)			   NULL
	, @C_Address NVARCHAR(1000)		   NULL
	, @F_Season_Mode TINYINT 		   NULL
	, @F_Fuel_Type TINYINT 			   NULL
	, @F_Period		TINYINT			   
	, @N_Norm_Supply_Value_Min MONEY   NULL
	, @N_Norm_Supply_Value_Max MONEY   NULL
	, @N_Time			  MONEY		   NULL
	, @N_Temp_Cold_Water  MONEY		   NULL
	, @N_Temp_Air		  MONEY		   NULL
	, @N_Time_Heat		  MONEY		   NULL
	, @N_Temp_Heat		  MONEY		   NULL
	, @N_Heat_Bought	  MONEY		   NULL
WITH EXECUTE AS OWNER
AS
BEGIN
	SET NOCOUNT ON;

	CREATE TABLE #Sources 
	(
		ID INT
	)

	INSERT INTO #Sources
	(
		ID
	)
	EXEC [dbo].[GetFilteredPaginatedSources] @C_Name
											, @C_Address 
											, @F_Season_Mode
											, @F_Fuel_Type
											, @N_Norm_Supply_Value_Min 
											, @N_Norm_Supply_Value_Max 
											/*
											, @F_Supplier_Electricity int 	   
											, @F_Voltage_Nominal tinyint 	   
											, @F_Transport_Gas int 			   
											, @F_Supplier_Gas int 			   
											, @F_Supplier_TechWater int 	   
											, @F_Supplier_HotWater int 		   
											, @F_Supplier_Canalisation int 	   
											, @F_Supplier_Heat int 			   
											*/
											, 0		-- @OffSetRow INT 				   
											, 0		-- @PageSize  INT				  
											, 0		-- @Regime
	BEGIN TRANSACTION
		UPDATE sf WITH (UPDLOCK, SERIALIZABLE)
		SET 
					sf.[N_Work_Duration] =	ISNULL(@N_Time, sf.[N_Work_Duration])			  
				,	sf.[N_Temp_Water] =		ISNULL(@N_Temp_Cold_Water, sf.[N_Temp_Water])  
				,	sf.[N_Temp_Air] =		ISNULL(@N_Temp_Air, sf.[N_Temp_Air])		  
				,	sf.[N_Heat_Duration] =	ISNULL(@N_Time_Heat, sf.[N_Heat_Duration])		  
				, 	sf.[N_Temp_Heat] =		ISNULL(@N_Temp_Heat, sf.[N_Temp_Heat])		  
				,	sf.[N_PaidHeat] =		ISNULL(@N_Heat_Bought, sf.[N_PaidHeat])	  
		FROM #Sources AS s
			INNER JOIN [dbo].[Source_Facts] AS sf ON s.ID = sf.F_Source
		WHERE sf.F_Calc_Period = @F_Period

		IF @@ROWCOUNT = 0
		BEGIN
			INSERT [dbo].[Source_Facts] 
			(
				[F_Source]
				, [F_Calc_Period]
				, [N_Work_Duration]		  
				, [N_Temp_Water]		  
				, [N_Temp_Air]		
				, [N_Heat_Duration]  
				, [N_Temp_Heat] 	
				, [N_PaidHeat] 
			)
			SELECT 
				s.ID
				, @F_Period
				, @N_Time
				, @N_Temp_Cold_Water
				, @N_Temp_Air
				, @N_Time_Heat
				, @N_Temp_Heat
				, @N_Heat_Bought
			FROM #Sources AS s  
		END
	COMMIT TRANSACTION;
END

