-- 1. Создать две базы данных на одном экземпляре СУБД SQL Server 2012.
use master;
go
if DB_ID (N'lab13_1') is not null
drop database lab13_1;
go
create database lab13_1
on (
NAME = lab131dat,
FILENAME = 'C:\data\lab131dat.mdf',
SIZE = 10,
MAXSIZE = 25,
FILEGROWTH = 5
)
log on (
NAME = lab131log,
FILENAME = 'C:\data\lab131log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

use master;
go
if DB_ID (N'lab13_2') is not null
drop database lab13_2;
go
create database lab13_2
on (
NAME = lab132dat,
FILENAME = 'C:\data\lab132dat.mdf',
SIZE = 10,
MAXSIZE = 25,
FILEGROWTH = 5
)
log on (
NAME = lab132log,
FILENAME = 'C:\data\lab132log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

-- 2. Создать в базах данных п.1. горизонтально фрагментированные таблицы.

use lab13_1;
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
	department nchar(20) NOT NULL CHECK (department IN (N'ИУ',N'РК',N'ИБМ',N'Э',
											N'СМ', N'БМТ')),
	budget money NULL CHECK (budget >= 0.0),
	CONSTRAINT Seq_cinema_more CHECK (student_id <= 4)
)
go


use lab13_2;
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
	department nchar(20) NOT NULL CHECK (department IN (N'ИУ',N'РК',N'ИБМ',N'Э',
											N'СМ', N'БМТ')),
	budget money NULL CHECK (budget >= 0.0),
	CONSTRAINT Seq_cinema_more CHECK (student_id > 4)
)
go



-- 3. Создать секционированные представления, 
-- обеспечивающие работу с данными таблиц
-- (выборку, вставку, изменение, удаление).

use lab13_1;
go

if OBJECT_ID(N'StudentView',N'V') is NOT NULL
	DROP VIEW StudentView;
go

CREATE VIEW StudentView AS
	SELECT * FROM lab13_1.dbo.Student
	UNION ALL
	SELECT * FROM lab13_2.dbo.Student
go

use lab13_2;
go

if OBJECT_ID(N'StudentView',N'V') is NOT NULL
	DROP VIEW StudentView;
go

CREATE VIEW StudentView AS
	SELECT * FROM lab13_1.dbo.Student
	UNION ALL
	SELECT * FROM lab13_2.dbo.Student
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


SELECT * FROM StudentView;

SELECT * from lab13_1.dbo.Student;
SELECT * from lab13_2.dbo.Student;


DELETE FROM StudentView WHERE department = 'СМ'

SELECT * from lab13_1.dbo.Student;
SELECT * from lab13_2.dbo.Student;


UPDATE StudentView SET year_birth = 2000 WHERE name = 'Антон'
 
SELECT * from lab13_1.dbo.Student;
SELECT * from lab13_2.dbo.Student;