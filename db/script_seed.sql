INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP1','Human Resources', 'Floor 1');
INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP2','IT (Cetec ERP)', 'Floor 1');
INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP3','Sales', 'Floor 1');
INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP4','Finance', 'Floor 2');
INSERT INTO Staff (FirstName, Lastname, Email, Profession, Sex, DOB, HireDate, DepartmentID) VALUES 
		('Mohamed Fares','Ben Ayed', 'feres.ba@qorevirtual.com', 'Software engineer', 'Male', '1997-03-06', '2023-01-20', 2);
-- Seed for Candidate table
INSERT INTO Candidate (FirstName, LasttName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB)
VALUES
    ('John', 'Doe', 'john.doe@example.com', 'Software Developer', 'Bachelor in Computer Science', '3 years at XYZ Company', 'Male', '1995-08-15'),
    ('Alice', 'Smith', 'alice.smith@example.com', 'Data Analyst', 'Master in Statistics', '2 years at ABC Corporation', 'Female', '1990-05-22'),
    ('Bob', 'Johnson', 'bob.johnson@example.com', 'Network Engineer', 'Bachelor in Electrical Engineering', '4 years at DEF Inc.', 'Male', '1988-11-10');

-- Seed for Staff table
INSERT INTO Staff (FirstName, Lastname, Email, Profession, Sex, DOB, HireDate, DepartmentID)
VALUES
    ('Sarah', 'Williams', 'sarah.williams@example.com', 'HR Specialist', 'Female', '1985-07-18', '2022-03-10', 1),
    ('Michael', 'Jones', 'michael.jones@example.com', 'System Administrator', 'Male', '1990-12-05', '2021-11-15', 2),
    ('Emily', 'Brown', 'emily.brown@example.com', 'Sales Representative', 'Female', '1993-04-30', '2023-02-05', 3),
    ('David', 'Miller', 'david.miller@example.com', 'Financial Analyst', 'Male', '1987-09-25', '2022-05-20', 4);