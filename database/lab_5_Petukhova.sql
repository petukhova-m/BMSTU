
use master;
go
if DB_ID (N'lab5') is not null
drop database lab5;
go
create database lab5
on (
NAME = lab5dat,
FILENAME = 'C:\data\lab5dat.mdf',
SIZE = 10,
MAXSIZE = UNLIMITED,
FILEGROWTH = 5
)
log on (
NAME = lab5log,
FILENAME = 'C:\data\lab5log.ldf',
SIZE = 5,
MAXSIZE = 20,
FILEGROWTH = 5
);
go 



use lab5;
go 
if OBJECT_ID(N'Users',N'U') is NOT NULL
	DROP TABLE Users;
go

CREATE TABLE Users (
	e_mail varchar(320) PRIMARY KEY NOT NULL,
	surname char(30) NOT NULL,
	name char(30) NOT NULL,
	phone char(11) NOT NULL);
go

INSERT INTO USERS(e_mail,surname,name,phone) 
VALUES ('run@gmail.com','Ivanov','Ivan', '89165678989')
go

SELECT * FROM USERS
go

use master;
go

alter database lab5
add filegroup lab5_fg
go

alter database lab5
add file
(
	NAME = lab5dat1,
	FILENAME = 'C:\data\lab5dat1.ndf',
	SIZE = 10MB,
	MAXSIZE = 100MB,
	FILEGROWTH = 5MB
)
to filegroup lab5_fg
go



alter database lab5
	modify filegroup lab5_fg default;
go



use lab5;
go 
if OBJECT_ID(N'Book',N'U') is NOT NULL
	DROP TABLE Book;
go

CREATE TABLE Book (
	name_book int PRIMARY KEY NOT NULL,
	date_of_publising numeric(4) NOT NULL,
	author char(50) NOT NULL);
	go


alter database lab5
	modify filegroup [primary] default;
go




use lab5;
go

drop table Book
go


--DBCC SHRINKFILE (lab5dat1, 0);
--
GO

alter database lab5
remove file lab5dat1
go

alter database lab5
remove filegroup lab5_fg;
go

use lab5;
go

CREATE SCHEMA library_schema
go

ALTER SCHEMA library_schema TRANSFER dbo.Users
go

DROP TABLE library_schema.Users
DROP SCHEMA library_schema
go
 




