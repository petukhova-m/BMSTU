use master;
go
if DB_ID (N'lab6') is not null
drop database lab6;
go
create database lab6
on (
NAME = lab6dat,
FILENAME = 'C:\data\lab6dat.mdf',
SIZE = 10,
MAXSIZE = UNLIMITED,
FILEGROWTH = 5
)
log on (
NAME = lab6log,
FILENAME = 'C:\data\lab6log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 

use lab6;
go 
if OBJECT_ID(N'Author',N'U') is NOT NULL
	DROP TABLE Author;
go



CREATE TABLE Author (
	author_id int IDENTITY(1,1) PRIMARY KEY,
	surname nchar(30) NOT NULL,
	name nchar(30) NOT NULL,
	date_of_birth numeric(4) NULL CHECK (date_of_birth>1500 AND date_of_birth<2000),
	date_of_death numeric(4) NULL CHECK (date_of_death>1500 AND date_of_death<2000), 
	biography nvarchar(1000) DEFAULT ('Unknown'),
	CONSTRAINT checkAuthor CHECK (date_of_birth<date_of_death)
	);
go

INSERT INTO Author(name,surname,date_of_birth,date_of_death)
VALUES (N'������',N'����', 1947, NULL),
	   (N'�������',N'�������',1962, NULL),
	   (N'������',N'�����',1969,NULL),
	   (N'�����',N'����� ����', 1859, 1930),
	   (N'������',N'�������', 1812, 1870),
	   (N'�����',N'������', 1854, 1900),
	   (N'������',N'�� ����-��������', 1900, 1944)
	  -- (N'���', N'��������', 1456, 1944)
go

SELECT * FROM Author
go


select IDENT_CURRENT('dbo.Author') as last_id
go


if OBJECT_ID(N'Book',N'U') is NOT NULL
	DROP TABLE Book;
go

CREATE TABLE Book (
	book_id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT (NEWID()),
	author nchar(50) NOT NULL,
	name nchar(50) NOT NULL,
	genre nchar(20) NOT NULL CHECK (genre IN (N'�����',N'������� ����������',N'�����',N'��������',
											N'�������', N'������',N'������', N'����������', N'�����')),
	publish_year numeric(4) NOT NULL CHECK (publish_year>1500 AND publish_year<2000),
	publish_house varchar(100) NULL DEFAULT ('Unknown'),
	cost_of smallmoney NULL CHECK (cost_of > 0),
	);
go

INSERT INTO Book(author,name,genre,publish_year)
VALUES (N'��������� ������',N'������� ������', N'�����', 1831),
	   (N'���� ����',N'20 000 ��� ��� �����',N'������� ����������', 1916),
	   (N'����� ������',N'�������� ������� �������',N'��������',1926),
	   (N'������ ����',N'1408', N'�������', 1926),
	   (N'������ ���������',N'1408', N'������', 1936),
	   (N'������ ����',N'������� �����', N'�����', 1910)
go
--

SELECT * FROM Book
go



IF EXISTS (SELECT * FROM sys.sequences WHERE NAME = N'TestSequence')
DROP SEQUENCE TestSequence
go

CREATE SEQUENCE TestSequence
	START WITH 0
	INCREMENT BY 1
	MAXVALUE 10;
go

if OBJECT_ID(N'Streets',N'U') is NOT NULL
	DROP TABLE Streets;
go

CREATE TABLE Streets (
	street_id int PRIMARY KEY NOT NULL,
	street_name nchar(50) DEFAULT (N'�����'),
	);
go

INSERT INTO Streets(street_id,street_name)
VALUES (NEXT VALUE FOR DBO.TestSequence,N'����������'),
	   (NEXT VALUE FOR DBO.TestSequence,N'�����������'),
	   (NEXT VALUE FOR DBO.TestSequence,N'������������'),
	   (NEXT VALUE FOR DBO.TestSequence,N'�����������'),
	   (NEXT VALUE FOR DBO.TestSequence,N'����������')
go

SELECT * From Streets
go


if OBJECT_ID(N'Client',N'U') is NOT NULL
	DROP TABLE Client;
go
	
CREATE TABLE Client (
	client_id int PRIMARY KEY NOT NULL,
	telephone bigint NOT NULL,
	name nchar(50) NOT NULL,
	email varchar(320) NULL,
	date_of_birth datetime NULL
	);
go

INSERT INTO Client(client_id, name, telephone)
VALUES (1,N'������', N'89674563454'),
	   (2,N'����',N'89674563454'),
	   (3,N'���������',N'89674563454'),
	   (4,N'��������', N'89674563454'),
	   (5,N'��������', N'89674563454'),
	   (6,N'������', N'89674563454'),
	   (7,N'��������', N'89674563454')
go

SELECT * FROM Client
go

if OBJECT_ID(N'Visit',N'U') is NOT NULL
	DROP TABLE Visit;
go

CREATE TABLE Visit (
	visit_id int IDENTITY(1,1) PRIMARY KEY,
	visit_date date DEFAULT (CONVERT(date,GETDATE())),
	visit_time time(0) DEFAULT (CONVERT(time,GETDATE())),
	visit_client int default 1,
	visit_procedure nchar(100) DEFAULT (N'����������'),
	CONSTRAINT FK_Client FOREIGN KEY (visit_client) REFERENCES Client (client_id)
	--ON DELETE Set default 
	--on delete cascade
	--on delete set null
	--on update cascade
	--on update set default
	--on update set null
	--on delete no action
	--on upate no action
	);
go

INSERT INTO Visit(visit_date,visit_time,visit_client)
VALUES (CONVERT(date,N'11-01-2014'),CONVERT(time,N'12:20:00'),3),
	   (CONVERT(date,N'19-03-2014'),CONVERT(time,N'13:30:00'),6),
	   (CONVERT(date,N'21-06-2014'),CONVERT(time,N'15:00:00'),2),
	   (CONVERT(date,N'26-08-2014'),CONVERT(time,N'16:10:00'),7),
	   (CONVERT(date,N'06-10-2014'),CONVERT(time,N'17:20:00'),3)  
go

SELECT * FROM Visit
go

delete from Client
where name = N'������'
go

--update Client
--set client_id=8
--where name = '����'


SELECT * FROM Client
go
SELECT * FROM Visit
go 