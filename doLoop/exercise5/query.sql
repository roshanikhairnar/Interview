-- create database student
create database student;

--use database
use student;

--create table
create table studentData(id int,firstName varchar,lastName varchar,department varchar);

-- insert values create student
Insert into studentData values(1,"abc","def","tyu");
Insert into studentData values(2,"asd","ghj","lkj");
Insert into studentData values(3,"zxc","bnm","rfv");

-- get student using id
SELECT * from studentData WHERE id=1;

-- get all students from table
SELECT * from studentData;

-- update student from table
UPDATE studentData SET firstName="roshani" WHERE id=1;

-- delete student from table
DELETE FROM studentData WHERE id=3;
