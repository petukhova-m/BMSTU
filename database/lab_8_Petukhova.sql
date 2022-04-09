use master;
go
if DB_ID (N'lab8_1') is not null
drop database lab8_1;
go
create database lab8_1
on (
NAME = lab8_1dat,
FILENAME = 'C:\data\lab8_1dat.mdf',
SIZE = 10,
MAXSIZE = UNLIMITED,
FILEGROWTH = 5
)
log on (
NAME = lab8_1log,
FILENAME = 'C:\data\lab8_1log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

use lab8_1;
go 
if OBJECT_ID(N'Student',N'U') is NOT NULL
	DROP TABLE Student;
go

CREATE TABLE Student (
	student_id int NOT NULL PRIMARY KEY,
	surname nchar(50) NOT NULL,
	name nchar(50) NOT NULL,
	patronymic nchar(50) NOT NULL,
	year_birth numeric(4) NOT NULL CHECK (year_birth>1980 AND year_birth<2005),
	department nchar(20) NOT NULL CHECK (department IN (N'ИУ',N'РК',N'ИБМ',N'Э', N'СМ', N'БМТ')),
	budget money NULL CHECK (budget >= 0.0),
)
go

INSERT INTO Student(student_id,surname,name,patronymic,year_birth, department, budget)
VALUES (1,N'Иванов', N'Иван', N'Иванович', 2000, N'СМ', 0),
	   (2,N'Карванова',N'Анна', N'Андреевна', 1999, N'ИБМ', 275000),
	   (3,N'Хачатурян',N'Инна', N'Максимовна', 2004, N'ИУ', 295000),
	   (4,N'Краснов',N'Артем', N'Викторович', 1998, N'СМ', 255000),
	   (5,N'Антонов',N'Алексей', N'Валентинович', 2000, N'БМТ', 0),
	   (6,N'Калинов',N'Антон', N'Андреевич', 1999, N'Э', 0),
	   (7,N'Ростова',N'Екатерина', N'Игоревна', 2003, N'ИБМ', 285000)
go

--SELECT * FROM Student
--go

--Создать хранимую процедуру, производящую выборку из некоторой таблицы
-- и возвращающую результат выборки в виде курсора.

IF OBJECT_ID(N'dbo.select_proc', N'P') IS NOT NULL
	DROP PROCEDURE dbo.select_proc
GO

CREATE PROCEDURE dbo.select_proc
	@cursor CURSOR VARYING OUTPUT
AS
	SET @cursor = CURSOR 
	FORWARD_ONLY STATIC FOR
	SELECT surname,name,patronymic,year_birth
	FROM Student

	OPEN @cursor;
GO

DECLARE @student_cursor CURSOR;
EXECUTE dbo.select_proc @cursor = @student_cursor OUTPUT;


FETCH NEXT FROM @student_cursor;
WHILE (@@FETCH_STATUS = 0)
BEGIN
	FETCH NEXT FROM @student_cursor;
END

CLOSE @student_cursor;
DEALLOCATE @student_cursor;
GO

-- Модифицировать хранимую процедуру п.1. таким образом, чтобы выборка   --
-- осуществлялась с формированием столбца, значение которого формируется --
-- пользовательской функцией. --

IF OBJECT_ID(N'random_number',N'FN') IS NOT NULL
	DROP FUNCTION random_number
go

IF OBJECT_ID(N'view_number',N'V') IS NOT NULL
	DROP VIEW view_number
go

CREATE VIEW view_number AS
	SELECT CAST(CAST(NEWID() AS binary(3)) AS INT) AS NextID
go

-- Возвращает случайное число в интервале [a;b]
CREATE FUNCTION random_number(@a int,@b int)
	RETURNS int
	AS
		BEGIN
			DECLARE @number int
			SELECT TOP 1 @number=NextID from view_number
			SET @number = @number % @b + @a
			RETURN (@number)
		END;
go

IF OBJECT_ID(N'dbo.select_proc_with_add', N'P') IS NOT NULL
	DROP PROCEDURE dbo.select_proc_with_add
GO

CREATE PROCEDURE dbo.select_proc_with_add
	@cursor CURSOR VARYING OUTPUT
AS
	SET @cursor = CURSOR 
	FORWARD_ONLY STATIC FOR
	SELECT surname,name,patronymic,dbo.random_number(1,100) as rating
	FROM Student

	OPEN @cursor;

GO

DECLARE @student_rating_cursor CURSOR;
EXECUTE dbo.select_proc_with_add @cursor = @student_rating_cursor OUTPUT;

FETCH NEXT FROM @student_rating_cursor;
WHILE (@@FETCH_STATUS = 0)
	BEGIN
		FETCH NEXT FROM @student_rating_cursor;
	END

CLOSE @student_rating_cursor;
DEALLOCATE @student_rating_cursor;
GO

-- Создать хранимую процедуру, вызывающую процедуру п.1., осуществляющую прокрутку возвращаемого  --
-- курсора и выводящую сообщения, сформированные из записей при выполнении условия, заданного     --
-- еще одной пользовательской функцией.													          --

IF OBJECT_ID(N'century_birth',N'FN') IS NOT NULL
	DROP FUNCTION century_birth
go

CREATE FUNCTION century_birth(@a numeric(4))
	RETURNS int
	AS
		BEGIN
			DECLARE @result int
			IF (@a>2000)
				SET @result = 21
			ELSE
				SET @result = 20
			RETURN (@result)
		END;
go

IF OBJECT_ID(N'dbo.external_proc',N'P') IS NOT NULL
	DROP PROCEDURE dbo.external_proc
GO

CREATE PROCEDURE dbo.external_proc 
AS
	DECLARE @external_cursor CURSOR;
	DECLARE	@table_surname nchar(50);
	DECLARE @table_name nchar(50);
	DECLARE @table_patronymic nchar(50);
	DECLARE @table_year_birth numeric(4);
	
	EXECUTE dbo.select_proc @cursor = @external_cursor OUTPUT;

	FETCH NEXT FROM @external_cursor INTO @table_surname,@table_name,@table_patronymic,@table_year_birth
	
	WHILE (@@FETCH_STATUS = 0)
	BEGIN
		IF (dbo.century_birth(@table_year_birth)=21)
			PRINT @table_surname + ' ' + @table_name + ' ' + @table_patronymic + N' (родился в 21 веке)'
		ELSE
			PRINT @table_surname + ' ' + @table_name + ' ' + @table_patronymic + N' (родился в 20 веке)'
		FETCH NEXT FROM @external_cursor INTO @table_surname,@table_name,@table_patronymic,@table_year_birth;
	END

	CLOSE @external_cursor;
	DEALLOCATE @external_cursor;

GO

EXECUTE dbo.external_proc
GO
		
--- Модифицировать хранимую процедуру п.2. таким образом, чтобы выборка
--- формировалась с помощью табличной функции.
IF OBJECT_ID(N'table_function_inline',N'TF') IS NOT NULL
	DROP FUNCTION table_function_inline
go	


/*CREATE FUNCTION table_function_inline()
	RETURNS TABLE
AS
	RETURN (
		SELECT surname,name,patronymic,year_birth,dbo.random_number(1,100) as rating
		FROM Student
		WHERE (dbo.century_birth(year_birth) = 21)
	)
GO --*/

if object_id(N'dbo.table_function_line',N'FN') is not null
drop procedure dbo.table_function_line;
go

create function dbo.table_function_line ()
returns @return_table table
(
	surname nchar(50) NOT NULL,
	name nchar(50) NOT NULL,
	patronymic nchar(20) NOT NULL ,
	year_birth numeric(4) NOT NULL,
	 rating numeric(2) NOT NULL
)
as
begin
insert @return_table select surname,name,patronymic,year_birth,dbo.random_number(1,100) as rating
from Student
where (dbo.century_birth(year_birth) = 21)
return;
end
go  --*/

ALTER PROCEDURE dbo.select_proc_with_add
	@cursor CURSOR VARYING OUTPUT
AS
	SET @cursor = CURSOR 
	FORWARD_ONLY STATIC FOR 
	SELECT * FROM dbo.table_function_line()
	OPEN @cursor;
GO	

DECLARE @student_table_cursor CURSOR;
EXECUTE dbo.select_proc_with_add @cursor = @student_table_cursor OUTPUT;

FETCH NEXT FROM @student_table_cursor;
WHILE (@@FETCH_STATUS = 0)
	BEGIN
		FETCH NEXT FROM @student_table_cursor;
	END

CLOSE @student_table_cursor;
DEALLOCATE @student_table_cursor;
GO