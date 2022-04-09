USE master
GO

IF DB_ID(N'lab10') IS NOT NULL
	DROP DATABASE lab10
GO

CREATE DATABASE lab10
ON (
	NAME = lab10dat,
	FILENAME = 'C:\data\lab10dat.mdf',
	SIZE = 10,
	MAXSIZE = UNLIMITED,
	FILEGROWTH = 5
)
LOG ON (
	NAME = lab10log,
	FILENAME = 'C:\data\lab10log.ldf',
	SIZE = 5,
	MAXSIZE = 25, 
	FILEGROWTH = 5
)
GO

USE lab10
GO 

IF OBJECT_ID (N'Client') IS NOT NULL
	DROP TABLE Client
GO

CREATE TABLE Client (
	client_id int  PRIMARY KEY NOT NULL,
	telephone bigint NOT NULL,
	name nchar(50) NOT NULL,
	email varchar(320) NULL,
	date_of_birth datetime NULL
	);
go
INSERT INTO Client(client_id, name, telephone)
VALUES (1,N'������', N'89674563454'),
	   (2,N'����',N'8967478093'),
	   (3,N'���������',N'89990896578'),
	   (4,N'��������', N'84556789876'),
	   (5,N'��������', N'89078970989'),
	   (6,N'������', N'84956783456'),
	   (7,N'��������', N'89771238743')
go

if OBJECT_ID(N'Visit',N'U') is NOT NULL
	DROP TABLE Visit;
go

CREATE TABLE Visit (
	visit_id int IDENTITY(1,1) PRIMARY KEY,
	visit_date date DEFAULT (CONVERT(date,GETDATE())),
	visit_time time(0) DEFAULT (CONVERT(time,GETDATE())),
	visit_client int default 1,
	price money NOT NULL
	);
go

INSERT INTO Visit(visit_date,visit_time,visit_client, price)
VALUES (CONVERT(date,N'11-01-2021'),CONVERT(time,N'12:20:00'),3, 20),
	   (CONVERT(date,N'19-03-2021'),CONVERT(time,N'13:30:00'),6, 45),
	   (CONVERT(date,N'21-06-2021'),CONVERT(time,N'15:00:00'),2, 67),
	   (CONVERT(date,N'26-08-2021'),CONVERT(time,N'16:10:00'),7, 100),
	   (CONVERT(date,N'06-10-2021'),CONVERT(time,N'17:20:00'),3, 10)  
go





-- 1. ����������� � ����������������� �� ��������
--    ��������� ������ �������� ���������� MS SQL
--    Server, ��������������� � ��������������
--    ���������� SET TRANSACTION ISOLATION LEVEL.

-- 2. ������������� ���������� ����������� �
--    �������������� sys.dm_tran_locks


-- 1 ������� ��������

SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
    BEGIN TRANSACTION
		SELECT * FROM Visit
		UPDATE Visit SET Price = Price + $5 WHERE (Price < $25)
		SELECT * FROM Visit
		SELECT * FROM sys.dm_tran_locks
	COMMIT TRANSACTION
GO

-- 2 ������� �������� 

SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
    BEGIN TRANSACTION;
		SELECT * FROM Client
		WAITFOR DELAY '00:00:05'
		SELECT * FROM Client
		SELECT * FROM sys.dm_tran_locks
    COMMIT TRANSACTION
GO

BEGIN TRANSACTION
	SELECT * FROM Visit
	UPDATE Visit SET price = price + 5 WHERE (visit_id = 2)
	SELECT * FROM Visit
	SELECT * FROM sys.dm_tran_locks
COMMIT TRANSACTION
GO

-- 3 ������� ��������

SET TRANSACTION ISOLATION LEVEL REPEATABLE READ
    BEGIN TRANSACTION
		INSERT INTO Client(client_id, name, telephone)
		VALUES (8, N'������', N'89674563454')
		SELECT * FROM Client
		WAITFOR DELAY '00:00:10'
		SELECT * FROM Client
		SELECT * FROM sys.dm_tran_locks
    COMMIT TRANSACTION
GO

-- 4 ������� ��������

SET TRANSACTION ISOLATION LEVEL SERIALIZABLE
    BEGIN TRANSACTION;
		INSERT INTO Visit(visit_date,visit_time,visit_client, price)
		VALUES (CONVERT(date,N'12-05-2021'),CONVERT(time,N'12:20:00'),3, 40)
		SELECT * FROM Visit
		WAITFOR DELAY '00:00:05'
		SELECT * FROM Visit
		SELECT * FROM sys.dm_tran_locks
    COMMIT TRANSACTION
GO