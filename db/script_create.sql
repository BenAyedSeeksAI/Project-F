DROP TABLE IF EXISTS Staff CASCADE;
DROP TABLE IF EXISTS Department CASCADE;
DROP TABLE IF EXISTS Candidate CASCADE;

-- Create Department table
CREATE TABLE Department (
    DepartmentID SERIAL PRIMARY KEY,
    DepartmentCode TEXT NOT NULL,
    DepartmentName TEXT NOT NULL,
    Location TEXT
);

-- Create Staff table with foreign key constraint
CREATE TABLE Staff (
    StaffID SERIAL PRIMARY KEY,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL,
    Email TEXT UNIQUE NOT NULL,
    Profession TEXT,
    Sex TEXT NOT NULL,
    DOB DATE,
    Username TEXT,
    Password TEXT,
    PhoneNumber TEXT,
    HireDate DATE,
    DepartmentID INT,
    CONSTRAINT FK_Staff_Department FOREIGN KEY (DepartmentID) REFERENCES Department(DepartmentID)
);

-- Create Candidate table
CREATE TABLE Candidate (
    CandidateID SERIAL PRIMARY KEY,
    FirstName TEXT NOT NULL,
    LastName TEXT NOT NULL, -- Typo: should be LastName
    PersonalEmail TEXT NOT NULL,
    JobOffer TEXT NOT NULL,
    Degree TEXT NOT NULL,
    RecentExperience TEXT NOT NULL,
    Sex TEXT NOT NULL,
    DOB DATE
);