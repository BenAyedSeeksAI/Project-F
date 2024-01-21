DROP TABLE IF EXISTS staff_info CASCADE;
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
CREATE TABLE staff_info (
    staff_status_id SERIAL PRIMARY KEY,
    contract_type TEXT,
    gross_salary NUMERIC,
    net_salary NUMERIC,
    bank_account_rib TEXT,
    social_account_code TEXT,
    title TEXT,
    advantages BOOLEAN,
    vacation_percentage NUMERIC,
    StaffID INT,
    CONSTRAINT fk_staff_staff_info FOREIGN KEY (StaffID) REFERENCES Staff(StaffID) ON DELETE CASCADE
);


-- Create Candidate tableexport GOPATH="$HOME/go"
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