-- Cоздание вертикально фрагментированных таблиц средствами СУБД SQL Server 2012

use master;
go
if DB_ID (N'lab14_1') is not null
drop database lab14_1;
go
create database lab14_1
on (
NAME = lab141dat,
FILENAME = 'C:\data\lab141dat.mdf',
SIZE = 10,
MAXSIZE = 25,
FILEGROWTH = 5
)
log on (
NAME = lab141log,
FILENAME = 'C:\data\lab141log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

use master;
go
if DB_ID (N'lab14_2') is not null
drop database lab14_2;
go
create database lab14_2
on (
NAME = lab142dat,
FILENAME = 'C:\data\lab142dat.mdf',
SIZE = 10,
MAXSIZE = 25,
FILEGROWTH = 5
)
log on (
NAME = lab142log,
FILENAME = 'C:\data\lab142log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

-- 1.Создать в базах данных пункта 1 задания 13 таблицы, содержащие вертикально фрагментированные данные.

use lab14_1;
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
	--department nchar(20) NOT NULL CHECK (department IN (N'ИУ',N'РК',N'ИБМ',N'Э', N'СМ', N'БМТ')),
	--budget money NULL CHECK (budget >= 0.0),
)
go

use lab14_2;
go

if OBJECT_ID(N'Student',N'U') is NOT NULL
	DROP TABLE Student;
go

CREATE TABLE Student (
	student_id int NOT NULL PRIMARY KEY,
	--surname nchar(50) NOT NULL,
	--name nchar(50) NOT NULL,
	--patronymic nchar(50) NOT NULL,
	--year_birth numeric(4) NOT NULL CHECK (year_birth>1980 AND year_birth<2005),
	department nchar(20) NOT NULL CHECK (department IN (N'ИУ',N'РК',N'ИБМ',N'Э',
											N'СМ', N'БМТ')),
	budget money NULL CHECK (budget >= 0.0),
)
go



-- 2.Создать необходимые элементы базы данных (представления, триггеры), 
-- обеспечивающие работу с данными вертикально фрагментированных таблиц 
-- (выборку, вставку, изменение, удаление). 

if OBJECT_ID(N'StudentView',N'V') is NOT NULL
	DROP VIEW StudentView;
go

CREATE VIEW StudentView AS
	SELECT A.*, B.department,B.budget
	FROM lab14_1.dbo.Student A, lab14_2.dbo.Student B
	WHERE A.student_id = B.student_id
go



IF OBJECT_ID(N'InsertStudentView',N'TR') IS NOT NULL
	DROP TRIGGER InsertStudentView
go

CREATE TRIGGER InsertStudentView
	ON StudentView
	INSTEAD OF INSERT 
AS
	BEGIN
		IF EXISTS (SELECT A.student_id
					   FROM lab14_1.dbo.Student AS A,
							lab14_2.dbo.Student AS B,
							inserted AS I
					   WHERE A.name = I.name AND A.surname = I.surname AND A.year_birth = I.year_birth AND B.department = I.department)
			BEGIN
				EXEC sp_addmessage 50003, 15,N'Такой студент уже есть в базе!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END
		ELSE
			IF EXISTS (SELECT A.student_id 
						FROM lab14_1.dbo.Student AS A, 
							  inserted AS I
					   WHERE A.student_id = I.student_id)
				BEGIN
					EXEC sp_addmessage 50004, 15,N'ID занят! Попробуйте другой',@lang='us_english',@replace='REPLACE';
					RAISERROR(50004,15,-1)
				END
			ELSE
				BEGIN
					INSERT INTO lab14_1.dbo.Student(student_id,surname,name,patronymic, year_birth)
					SELECT student_id,surname,name,patronymic, year_birth FROM inserted

					INSERT INTO lab14_2.dbo.Student(student_id,department,budget)
					SELECT student_id,department,budget FROM inserted
				END
	END
go

INSERT INTO StudentView(student_id,surname,name,patronymic,year_birth, department, budget)
VALUES (1,N'Иванов', N'Иван', N'Иванович', 2000, N'СМ', 0),
	   (2,N'Карванова',N'Анна', N'Андреевна', 1999, N'ИБМ', 275000),
	   (3,N'Хачатурян',N'Инна', N'Максимовна', 2004, N'ИУ', 295000),
	   (4,N'Краснов',N'Артем', N'Викторович', 1998, N'СМ', 255000),
	   (5,N'Антонов',N'Алексей', N'Валентинович', 2000, N'БМТ', 0),
	   (6,N'Калинов',N'Антон', N'Андреевич', 1999, N'Э', 0),
	   (7,N'Ростова',N'Екатерина', N'Игоревна', 2003, N'ИБМ', 285000)
go

SELECT * FROM lab14_1.dbo.Student
go

SELECT * FROM lab14_2.dbo.Student
go
SELECT * FROM StudentView
go


IF OBJECT_ID(N'UpdateStudentView',N'TR') IS NOT NULL
	DROP TRIGGER UpdateStudentView
go

CREATE TRIGGER UpdateStudentView
	ON StudentView
	INSTEAD OF UPDATE
AS
	
	BEGIN
		IF UPDATE(student_id)
			BEGIN
				EXEC sp_addmessage 50001, 15,N'Запрещено изменение ID студента',@lang='us_english',@replace='REPLACE';
				RAISERROR(50001,15,-1)
			END
		IF UPDATE(surname) OR UPDATE(name) OR UPDATE(patronymic) OR UPDATE(year_birth)
			BEGIN
				EXEC sp_addmessage 50002, 15,N'Запрещено изменение информации о студенте.',@lang='us_english',@replace='REPLACE';
				RAISERROR(50002,15,-1)
			END

		DECLARE @temp_table TABLE (
					student_id int,
					add_budget money,
					delete_budget money,
					add_department nchar(20),
					delete_department nchar(20)
		);
		INSERT INTO @temp_table(student_id,add_budget,delete_budget, add_department, delete_department)
		SELECT A.student_id, A.budget,
							B.budget,
							A.department,
							B.department
		FROM inserted A
		INNER JOIN deleted B ON A.student_id = B.student_id

		UPDATE lab14_2.dbo.Student SET budget = inserted.budget FROM inserted WHERE inserted.student_id = lab14_2.dbo.Student.student_id
		UPDATE lab14_2.dbo.Student SET department = inserted.department FROM inserted WHERE inserted.student_id = lab14_2.dbo.Student.student_id
	END
go

UPDATE StudentView SET budget = 0 WHERE surname = N'Хачатурян'
go


SELECT * FROM StudentView
go

IF OBJECT_ID(N'DeleteStudentView',N'TR') IS NOT NULL
	DROP TRIGGER DeleteStudentView
go

CREATE TRIGGER DeleteStudentView
	ON StudentView
	INSTEAD OF DELETE
AS
	BEGIN
		DELETE C FROM lab14_1.dbo.Student AS C INNER JOIN deleted AS d ON C.student_id = d.student_id
		DELETE C FROM lab14_2.dbo.Student AS C INNER JOIN deleted AS d ON C.student_id = d.student_id
	END
go

DELETE FROM StudentView WHERE year_birth < 2000
go

SELECT * FROM StudentView
go

