-- Создание распределенных баз данных со связанными таблицами средствами СУБД SQL Server 2012

use master;
go
if DB_ID (N'lab15_1') is not null
drop database lab15_1;
go
create database lab15_1
on (
NAME = lab151dat,
FILENAME = 'C:\data\lab151dat.mdf',
SIZE = 10,
MAXSIZE = 25,
FILEGROWTH = 5
)
log on (
NAME = lab151log,
FILENAME = 'C:\data\lab151log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

use master;
go
if DB_ID (N'lab15_2') is not null
drop database lab15_2;
go
create database lab15_2
on (
NAME = lab152dat,
FILENAME = 'C:\data\lab152dat.mdf',
SIZE = 10,
MAXSIZE = 25,
FILEGROWTH = 5
)
log on (
NAME = lab152log,
FILENAME = 'C:\data\lab152log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

-- 1.Создать в базах данных пункта 1 задания 13 связанные таблицы.

use lab15_1;
go


if OBJECT_ID(N'Client',N'U') is NOT NULL
	DROP TABLE Client;
go

CREATE TABLE Client (
	client_id int  PRIMARY KEY NOT NULL,
	name nchar(50) NOT NULL,
	telephone bigint NOT NULL,
	
	email varchar(320) NULL,
	date_of_birth date NULL CHECK (date_of_birth>1930 AND date_of_birth<2005),
	);
go





use lab15_2;
go


if OBJECT_ID(N'Visit',N'U') is NOT NULL
	DROP TABLE Visit;
go

if OBJECT_ID(N'FK_Client',N'F') IS NOT NULL
	ALTER TABLE Visit DROP CONSTRAINT FK_Client
go



CREATE TABLE Visit (
	visit_id int PRIMARY KEY,
	visit_date date DEFAULT (CONVERT(date,GETDATE())),
	visit_time time(0) DEFAULT (CONVERT(time,GETDATE())),
	visit_client int default 1,
	visit_procedure nchar(100) DEFAULT (N'Неизвестно'),
	--CONSTRAINT FK_Client FOREIGN KEY (visit_client) REFERENCES Client (client_id)
	--ON DELETE CASCADE
	--ON UPDATE CASCADE
	);
go


-- 2.Создать необходимые элементы базы данных (представления, триггеры), 
-- обеспечивающие работу с данными связанных таблиц 
-- (выборку, вставку, изменение, удаление). 

if OBJECT_ID(N'Client_VisitView',N'V') is NOT NULL
	DROP VIEW Client_VisitView;
go

CREATE VIEW Client_VisitView AS
	SELECT C.client_id as client_id, C.name as name,C.email as email, C.date_of_birth as date_of_birth,
			V.visit_date as date, V.visit_time as time
	FROM lab15_1.dbo.Client C, lab15_2.dbo.Visit V
	WHERE C.client_id = V.visit_client
go

use lab15_1
IF OBJECT_ID(N'InsertClient',N'TR') IS NOT NULL
	DROP TRIGGER InsertClient
go

CREATE TRIGGER InsertClient
	ON Client
	INSTEAD OF INSERT 
AS
	BEGIN
		IF EXISTS (SELECT C.client_id
					   FROM lab15_1.dbo.Client AS C,
							inserted AS I
					   WHERE C.client_id = I.client_id)
			BEGIN
				EXEC sp_addmessage 50001, 15,N'ID занят! Попробуйте другой!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50001,15,-1)
			END
		ELSE
			IF EXISTS (SELECT C.client_id
					   FROM Client AS C,
							inserted AS I
					   WHERE C.telephone = I.telephone)
			BEGIN
				EXEC sp_addmessage 50003, 15,N'Такой телефон уже есть в базе!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END
			ELSE
				INSERT INTO lab15_1.dbo.Client(client_id, name, telephone, email, date_of_birth)
				SELECT client_id, name, telephone, email, date_of_birth FROM inserted
		IF (SELECT  COUNT(*) FROM inserted) > 1
			PRINT 'Добавлены новые клиенты в таблицу'
		ELSE
			PRINT 'Добавлен новый клиент в таблицу'
	END
go

IF OBJECT_ID(N'DeleteClient',N'TR') IS NOT NULL
	DROP TRIGGER DeleteClient
go

CREATE TRIGGER DeleteClient
	ON Client
	INSTEAD OF DELETE
AS
	BEGIN
		DELETE B FROM lab15_2.dbo.Visit AS B INNER JOIN deleted AS d ON B.visit_client = d.client_id
		DELETE A FROM lab15_1.dbo.Client AS A INNER JOIN deleted AS d ON A.Client_id = d.Client_id
	END
go

IF OBJECT_ID(N'UpdateClient',N'TR') IS NOT NULL
	DROP TRIGGER UpdateClient
go

CREATE TRIGGER UpdateClient
	ON Client
	AFTER UPDATE
AS
	BEGIN
		IF (UPDATE(name)
			OR (UPDATE(date_of_birth) OR UPDATE(client_id)))
			BEGIN;
				EXEC sp_addmessage 50003, 15,N'Запрещено изменение личных данных клиента!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END;
		ELSE
			BEGIN;
				DECLARE @temp_table TABLE (
					client_id int PRIMARY KEY,
					add_telephone bigint, add_name nchar(50), add_email varchar(320), add_date_of_birth datetime,
					delete_telephone bigint, delete_name nchar(50), delete_email varchar(320), delete_date_of_birth datetime
				);

				INSERT INTO @temp_table(client_id,add_telephone,add_name, add_email, add_date_of_birth,
										delete_telephone, delete_name, delete_email, delete_date_of_birth)
				SELECT A.client_id, A.telephone,A.name,A.email,A.date_of_birth,
					                B.telephone,B.name,B.email,B.date_of_birth
				FROM inserted A
				INNER JOIN deleted B ON A.client_id = B.client_id

				IF UPDATE(telephone) 
					PRINT N'Был изменен телефон'
				IF UPDATE(email) AND EXISTS (SELECT TOP 1 delete_email FROM @temp_table WHERE delete_email IS NULL)
					PRINT N'Был добавлен email'
					else
				IF UPDATE(email)
					PRINT N'Был изменен email'

				DECLARE @number int;
				SET @number = (SELECT DISTINCT COUNT(*) FROM @temp_table);
				IF @number > 1
					PRINT N'у ' + CAST(@number AS VARCHAR(1)) + ' клиентов'
				ELSE
					PRINT N'у 1 клиента'
		END;
	END
go 


use lab15_2;
go

IF OBJECT_ID(N'InsertVisit',N'TR') IS NOT NULL
	DROP TRIGGER InsertVisit
go

CREATE TRIGGER InsertVisit
	ON Visit
	INSTEAD OF INSERT 
AS
	BEGIN
		IF EXISTS (SELECT B.Visit_id
					   FROM	Visit AS B,
							inserted AS I
					   WHERE B.visit_date = I.visit_date AND B.visit_time = I.visit_time AND B.visit_procedure = I.visit_procedure)
			BEGIN
				EXEC sp_addmessage 50004, 15,N'Данная дата и время уже заняты для этой процедуры!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50004,15,-1)
			END
		ELSE
			IF EXISTS (SELECT B.Visit_id
						FROM lab15_2.dbo.Visit AS B, 
							  inserted AS I
					   WHERE B.Visit_id = I.Visit_id)
				BEGIN
					EXEC sp_addmessage 50002, 15,N'ID занят! Попробуйте другой',@lang='us_english',@replace='REPLACE';
					RAISERROR(50002,15,-1)
				END
			ELSE
				IF EXISTS (SELECT Visit_id 
						   FROM inserted 
						   WHERE Visit_client NOT IN 
						   (SELECT client_id FROM lab15_1.dbo.Client))
				BEGIN
					EXEC sp_addmessage 50005, 15,N'Добавление посещения для несуществующего клиента!',@lang='us_english',@replace='REPLACE';
					RAISERROR(50005,15,-1)
				END
					
			ELSE
				INSERT INTO Visit(Visit_id, visit_date, Visit_time, visit_client, visit_procedure)
				SELECT Visit_id, visit_date, Visit_time, visit_client, visit_procedure FROM inserted
		IF (SELECT DISTINCT COUNT(*) FROM inserted) > 1
			PRINT 'Добавлены новые посещения в таблицу'
		ELSE
			PRINT 'Добавлен новое посещение в таблицу'
	END
go

IF OBJECT_ID(N'DeleteVisit',N'TR') IS NOT NULL
	DROP TRIGGER DeleteVisit
go

CREATE TRIGGER DeleteVisit
	ON Visit
	INSTEAD OF DELETE
AS
	BEGIN
		DELETE B FROM lab15_2.dbo.Visit AS B INNER JOIN deleted AS d ON B.Visit_id = d.Visit_id
	END
go

IF OBJECT_ID(N'UpdateVisit',N'TR') IS NOT NULL
	DROP TRIGGER UpdateVisit
go

CREATE TRIGGER UpdateVisit
	ON Visit
	AFTER UPDATE
AS
	BEGIN

		IF (UPDATE(Visit_id))
			BEGIN;
				EXEC sp_addmessage 50003, 15,N'Запрещено изменение id!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END;
		ELSE
		IF (UPDATE(Visit_client))
			BEGIN;
				EXEC sp_addmessage 50003, 15,N'Запрещено привязывать посещение к другому клиенту!',@lang='us_english',@replace='REPLACE';
				RAISERROR(50003,15,-1)
			END;
		ELSE
			BEGIN;

				IF UPDATE(visit_date)
					PRINT N'Была изменена дата посещения'
				IF UPDATE(visit_time)
					PRINT N'Было изменено время посещения'
				IF UPDATE(visit_procedure)
					PRINT N'Была изменена процедура'

				DECLARE @number int;
				SET @number = (SELECT DISTINCT COUNT(*) FROM inserted);
				IF @number > 1
					PRINT N'у ' + CAST(@number AS VARCHAR(1)) + ' посещений'
				ELSE
					PRINT N'у 1 посещения'
		END;
	END
go 

INSERT INTO lab15_1.dbo.Client(client_id, name, telephone)
VALUES (1,N'Полина', N'89674563454'),
	   (2,N'Анна',N'8967478093'),
	   (3,N'Станислав',N'89990896578'),
	   (4,N'Виктория', N'84556789876'),
	   (5,N'Вероника', N'89078970989'),
	   (6,N'Кирилл', N'84956783456'),
	   (7,N'Вячеслав', N'89771238743')
go
SELECT * FROM lab15_1.dbo.Client;

INSERT INTO lab15_2.dbo.Visit(visit_id, visit_date,visit_time,visit_client)
VALUES (1, CONVERT(date,N'11-01-2021'),CONVERT(time,N'12:20:00'),3),
	   (2, CONVERT(date,N'19-03-2021'),CONVERT(time,N'13:30:00'),6),
	   (3, CONVERT(date,N'21-06-2021'),CONVERT(time,N'15:00:00'),2),
	   (4, CONVERT(date,N'26-08-2021'),CONVERT(time,N'16:10:00'),7),
	   (5, CONVERT(date,N'06-10-2021'),CONVERT(time,N'17:20:00'),3)  
go

SELECT * FROM lab15_2.dbo.Visit;

/*UPDATE lab15_1.dbo.Client SET email = 'opi@mail.ru' WHERE client_id = 2
go
SELECT * FROM lab15_1.dbo.Client;--*/


/*UPDATE lab15_2.dbo.Visit SET visit_date =  CONVERT(date,N'07-10-2021') WHERE visit_client = 7
go

SELECT * FROM lab15_2.dbo.Visit;
go--*/

/*DELETE FROM lab15_2.dbo.Visit WHERE visit_id = 2
SELECT * FROM lab15_2.dbo.Visit;
go--*/


DELETE FROM lab15_1.dbo.Client WHERE name = N'Станислав'
SELECT * FROM lab15_1.dbo.Client;
SELECT * FROM lab15_2.dbo.Visit;
go--*/