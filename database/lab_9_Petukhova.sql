use master;
go
if DB_ID (N'lab9_1') is not null
drop database lab9_1;
go
create database lab9_1
on (
NAME = lab9_1dat,
FILENAME = 'C:\data\lab9_1dat.mdf',
SIZE = 10,
MAXSIZE = UNLIMITED,
FILEGROWTH = 5
)
log on (
NAME = lab9_1log,
FILENAME = 'C:\data\lab9_1log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

use lab9_1;
go 
if OBJECT_ID(N'Student',N'U') is NOT NULL
	DROP TABLE Student;
go

if OBJECT_ID(N'Uniq_Student',N'UQ') IS NOT NULL
	ALTER TABLE Student DROP CONSTRAINT Uniq_Student
go

CREATE TABLE Student (
	Student_id int  PRIMARY KEY NOT NULL,
	telephone bigint NOT NULL,
	lastname nchar(50) NOT NULL,
	name nchar(50) NOT NULL,
	email varchar(320) NULL,
	date_of_birth date NULL,
	IsGrad nchar(50) Not NULL Check (IsGrad in ('Graduate', 'UnderGraduate')),
	CONSTRAINT Uniq_Student UNIQUE (telephone)
	);
go


if OBJECT_ID(N'Graduate',N'U') is NOT NULL
	DROP TABLE Graduate;
go

if OBJECT_ID(N'FK_Student',N'F') IS NOT NULL
	ALTER TABLE Graduate DROP CONSTRAINT FK_Student
go



CREATE TABLE Graduate (
	Graduate_id int IDENTITY(1,1) PRIMARY KEY,
	number_diploma int NOT NULL,
	year_grad numeric(4) NOT NULL,
	Student_id int default 1,
	CONSTRAINT FK_Student FOREIGN KEY (Student_id) REFERENCES Student (Student_id)
	ON DELETE CASCADE
	ON UPDATE CASCADE
	);
go

if OBJECT_ID(N'UnderGraduate',N'U') is NOT NULL
	DROP TABLE UnderGraduate;
go

if OBJECT_ID(N'FK_Student1',N'F') IS NOT NULL
	ALTER TABLE UnderGraduate DROP CONSTRAINT FK_Student1
go

CREATE TABLE UnderGraduate (
	UnderGraduate_id int IDENTITY(1,1) PRIMARY KEY,
	student_card int NOT NULL,
	year_admission numeric(4) NOT NULL CHECK (year_admission > 2010),
	Student_id int default 1,
	CONSTRAINT FK_Student1 FOREIGN KEY (Student_id) REFERENCES Student (Student_id)
	ON DELETE CASCADE
	ON UPDATE CASCADE
	);
go
SET IDENTITY_INSERT GRADUATE OFF
SET IDENTITY_INSERT underGRADUATE ON



-- ��� ����� �� ������ ������ 2 ������� 7 ������� �������� �� �������, �������� � ����������,
-- ��� ���������� �������� ������� ���� �� ��������� ������ ������������ ������������� ������
-- (RAISERROR / THROW)

-- �������� �� ��������
use lab9_1;
go


IF OBJECT_ID(N'Delete_Student',N'TR') IS NOT NULL
	DROP TRIGGER Delete_Student
go

CREATE TRIGGER Delete_Student
	ON Student
	INSTEAD OF DELETE 
AS
	BEGIN
				DELETE FROM Student WHERE Student_id IN (SELECT Student_id FROM deleted)
				IF (SELECT DISTINCT COUNT(*) FROM deleted) > 1
					PRINT '������ � ��������� ������� �� �������!'
				ELSE
					PRINT '������ � �������� ������� �� �������!'
	END
go




-- ������� �� ���������� --

IF OBJECT_ID(N'Update_info_Student',N'TR') IS NOT NULL
	DROP TRIGGER Update_info_Student
go

CREATE TRIGGER Update_info_Student
	ON Student
	AFTER UPDATE
AS
	BEGIN
		IF ((UPDATE(Student_id)))
			BEGIN;
				EXEC sp_addmessage 50001, 15,N'��������� id �������� ����������!',  @lang='us_english', @replace='REPLACE'
				RAISERROR(50001,15,-1)
			END;
		ELSE
		IF ((UPDATE(name)) OR (UPDATE(lastname))
		OR (UPDATE(date_of_birth) AND EXISTS (SELECT TOP 1 date_of_birth FROM deleted WHERE date_of_birth is not null)))
			BEGIN;
				EXEC sp_addmessage 50002, 15,N'��������� ������ ������ �������� ����������!',  @lang='us_english', @replace='REPLACE'
				RAISERROR(50002,15,-1)
			END;
		
		ELSE
			BEGIN;
				IF UPDATE(date_of_birth) AND EXISTS (SELECT TOP 1 date_of_birth FROM deleted WHERE date_of_birth IS NULL)
					PRINT N'���� ��������� ���� ��������'
				ELSE
				IF UPDATE(telephone) 
					PRINT N'��� ������� �������'
				IF UPDATE(email) AND EXISTS (SELECT TOP 1 email FROM deleted WHERE email IS NULL)
					PRINT N'��� �������� email'
					else
				IF UPDATE(email)
					PRINT N'��� ������� email'

				DECLARE @number int;
				SET @number = (SELECT COUNT(*) FROM inserted);
				IF @number > 1
					PRINT N'� ' + CAST(@number AS VARCHAR(1)) + '���������'
				ELSE
					PRINT N'� 1 ��������'
		END;
	END
go 




-- ������� �� ������� --
		
IF OBJECT_ID(N'Add_Student',N'TR') IS NOT NULL
	DROP TRIGGER Add_Student
go

CREATE TRIGGER Add_Student
	ON Student
	INSTEAD OF INSERT 
AS
	BEGIN
		IF EXISTS (SELECT S.Student_id
					   FROM Student AS S,
							inserted AS I
					   WHERE S.lastname = I.lastname And S.name = I.name)
			BEGIN
				EXEC sp_addmessage 50003, 15,N'����� ������� ��� ���� � ����!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END
		ELSE
		BEGIN
			INSERT INTO Student(Student_id, telephone, lastname, name, email, date_of_birth, IsGrad)
			SELECT Student_id, telephone, lastname, name, email, date_of_birth, IsGrad FROM inserted
			IF (SELECT COUNT(*) FROM inserted) > 1
				PRINT '��������� ����� �������� � �������'
			ELSE
				PRINT '�������� ����� ������� � �������'
		END
	END
go



-- ��� ������������� ������ 2 ������� 7 ������� �������� �� �������, �������� � ����������,
-- �������������� ����������� ���������� �������� � ������� ��������������� ����� �������������

-- ������������� (View)


if OBJECT_ID(N'JoinStudentView',N'V') is NOT NULL
	DROP VIEW JoinStudentView;
go

CREATE VIEW JoinStudentView AS
	SELECT c.Student_id as Student_id, c.lastname as lastname, c.name as name,c.telephone as telephone, c.email as email,c.date_of_birth as date_of_birth, c.IsGrad as IsGrad,
		   v.number_diploma as number_diploma, v.year_grad as year_grad
	FROM Student as c INNER JOIN Graduate as v ON v.Student_id = c.Student_id
go


IF OBJECT_ID(N'Add_View_Student',N'TR') IS NOT NULL
	DROP TRIGGER Add_View_Student
go

CREATE TRIGGER Add_View_Student
	ON JoinStudentView
	INSTEAD OF INSERT 
AS
	BEGIN
		IF EXISTS (SELECT S.Student_id
					   FROM Student AS S,
							inserted AS I
					   WHERE S.lastname = I.lastname And S.name = I.name)
			BEGIN
				EXEC sp_addmessage 50003, 15,N'����� ������� ��� ���� � ����!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END
		ELSE
		BEGIN
			INSERT INTO Student(Student_id, telephone, lastname, name, email, date_of_birth, IsGrad)
			SELECT Student_id, telephone, lastname, name, email, date_of_birth, IsGrad
			FROM inserted
				INSERT INTO Graduate(Student_id, number_diploma, year_grad)
				SELECT Student_id, number_diploma, year_grad
				FROM inserted
		END

	END
go




IF OBJECT_ID(N'Delete_View_Student',N'TR') IS NOT NULL
	DROP TRIGGER Delete_View_Student
go

CREATE TRIGGER Delete_View_Student
	ON JoinStudentView
	INSTEAD OF DELETE
AS
	BEGIN
		DELETE FROM Student WHERE Student_id IN (SELECT Student_id FROM deleted)
		DECLARE @count int
		SET @count = (SELECT  COUNT(*) FROM deleted)
		IF @count > 1
			PRINT '������� ������ � ���������'
		IF @count = 1
			PRINT '������� ������ � ��������'
	END
go--*/



IF OBJECT_ID(N'Update_View_Student',N'TR') IS NOT NULL
	DROP TRIGGER Update_View_Student
go

CREATE TRIGGER Update_View_Student
	ON JoinStudentView
	INSTEAD OF UPDATE
AS
	BEGIN
		

		IF (UPDATE(number_diploma) OR UPDATE(year_grad))
			BEGIN
				EXEC sp_addmessage 50004, 15,N'��������� ��������� ������ ������ �������',@lang='us_english',@replace='REPLACE';
				RAISERROR(50004,15,-1)
			END
		ELSE
		/*IF UPDATE(client_id)
			BEGIN
				EXEC sp_addmessage 50005, 15,N'��������� ��������� id �������',@lang='us_english',@replace='REPLACE';
				RAISERROR(50005,15,-1)
			END
		ELSE*/
		BEGIN
			

			IF UPDATE(telephone)
				BEGIN
					UPDATE Student SET telephone = I.telephone FROM inserted as I, Student as S where I.Student_id = S.Student_id
					PRINT N'������� ��� �������'
				END
			IF UPDATE(email)
				BEGIN
					UPDATE Student SET email = I.email FROM inserted as I, Student as S where I.Student_id = S.Student_id
					PRINT N'Email ��� �������'
				END
		END
	END
go


INSERT INTO Student(Student_id, lastname, name, telephone, IsGrad)
VALUES (1, N'���������', N'������', N'89674563454', 'Graduate'),
	   (2,N'��������',N'����',N'8967478093', 'Graduate'),
	   (3,N'��������',N'���������',N'89990896578', 'UnderGraduate'),
	   (4,N'������',N'��������', N'84556789876', 'Graduate'),
	   (5,N'������',N'��������', N'89078970989', 'UnderGraduate'),
	   (6,N'�������',N'������', N'84956783456', 'UnderGraduate'),
	   (7,N'�����',N'��������', N'89771238743', 'Graduate')
go

INSERT INTO Graduate(Student_id, number_diploma, year_grad)
VALUES  (2, 34567, 2015),
		(1, 90786, 2017),
		(4, 1234, 2021),
		(7, 56789, 2019)
go

INSERT INTO UnderGraduate(UnderGraduate_id, Student_id, student_card, year_admission)
VALUES  (1, 3, 8900, 2020),
		(2, 5, 12345, 2019),
		(3, 6, 7800, 2021)
go


SELECT * FROM Student
SELECT * FROM Graduate
SELECT * FROM UnderGraduate
SELECT * FROM JoinStudentView
go-- */
SELECT *  FROM JoinStudentView
UPDATE JoinStudentView SET email=telephone 
SELECT *  FROM JoinStudentView

/* DELETE FROM Student WHERE Student_id in (2,3)
SELECT * FROM Student
SELECT * FROM Graduate
SELECT * FROM UnderGraduate
go --*/

/*UPDATE Student SET email = 'rty@yandex.ru' WHERE Student_id = 1
 UPDATE Student SET date_of_birth = CONVERT(date,N'11-01-2000') WHERE Student_id = 3
 --UPDATE Student SET Student_id = 0 WHERE Student_id = 1
 -- UPDATE Student SET name = 'Maria' WHERE Student_id = 2
SELECT * FROM Student
go --*/

/* INSERT INTO Student(Student_id,telephone, lastname, name, IsGrad)
VALUES (8,N'89778675645', N'������', N'���������', 'UnderGraduate'),
	   (9,N'89167563423', N'���������', N'�����', 'UnderGraduate')
SELECT * FROM Student
go --*/

/*

 INSERT INTO JoinStudentView(Student_id,telephone, lastname, name, IsGrad, number_diploma, year_grad)
VALUES (8, N'89776785698', N'��������', N'�����', 'Graduate', 6789, 2013)
SELECT * FROM JoinStudentView
SELECT * FROM Student
SELECT * FROM Graduate
go --*/

 /* UPDATE JoinStudentView SET email=N'opi@gmail.ru' WHERE name=N'��������'
  UPDATE JoinStudentView SET telephone = '89093089023' WHERE Student_id = 2
 -- UPDATE JoinStudentView SET year_grad = 2016 WHERE Student_id = 7
SELECT * FROM Student
SELECT * FROM Graduate
SELECT * FROM JoinStudentView
go --*/

/*DELETE FROM JoinStudentView WHERE (name=N'��������')
SELECT * FROM Student
SELECT * FROM Graduate
SELECT * FROM JoinStudentView
go--*/