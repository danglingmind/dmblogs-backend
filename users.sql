CREATE DATABASE USERSDB;

USE USERSDB;

-- users table 
CREATE TABLE USERS (
    ID int NOT NULL AUTO_INCREMENT UNIQUE,
    Firstname CHAR(100) NOT NULL,
    Middlename CHAR(100) NOT NULL,
    Lastname CHAR(100) NOT NULL,
    Mobile CHAR(30) UNIQUE,
    Email CHAR(30) UNIQUE,
    Password CHAR(100),
    Created TIMESTAMP, 
    Modified TIMESTAMP, 
    Active BOOL,
    PRIMARY KEY (ID)
);