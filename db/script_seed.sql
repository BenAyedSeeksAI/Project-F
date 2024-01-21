INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP1','Human Resources', 'Floor 1');
INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP2','IT (Cetec ERP)', 'Floor 1');
INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP3','Sales', 'Floor 1');
INSERT INTO Department (DepartmentCode, DepartmentName, Location) VALUES ('DP4','Finance', 'Floor 2');
INSERT INTO Staff (FirstName, Lastname, Email, Profession, Sex, DOB, HireDate, DepartmentID) VALUES 
		('Mohamed Fares','Ben Ayed', 'feres.ba@qorevirtual.com', 'Software engineer', 'Male', '1997-03-06', '2023-01-20', 2);
-- Seed for Candidate table
INSERT INTO Candidate (FirstName, LastName, PersonalEmail, JobOffer, Degree, RecentExperience, Sex, DOB)
VALUES
    ('John', 'Doe', 'john.doe@example.com', 'Software Developer', 'Bachelor in Computer Science', '3 years at XYZ Company', 'Male', '1995-08-15'),
    ('Alice', 'Smith', 'alice.smith@example.com', 'Data Analyst', 'Master in Statistics', '2 years at ABC Corporation', 'Female', '1990-05-22'),
    ('Bob', 'Johnson', 'bob.johnson@example.com', 'Network Engineer', 'Bachelor in Electrical Engineering', '4 years at DEF Inc.', 'Male', '1988-11-10');
INSERT INTO Candidate (
    FirstName,
    LastName,
    PersonalEmail,
    JobOffer,
    Degree,
    RecentExperience,
    Sex,
    DOB
) VALUES
    ('Ikey', 'Hellwig', 'ihellwig0@irs.gov', 'VP Marketing', 'Doctorate Degree', 'Recruiting Manager', 'Male', '1996-11-07'),
    ('Abigale', 'Gumb', 'agumb1@artisteer.com', 'Editor', 'Doctorate Degree', 'Accounting Assistant III', 'Female', '1982-08-09'),
    ('Dorree', 'Tynemouth', 'dtynemouth2@vimeo.com', 'Environmental Specialist', 'Bachelor Degree', 'Environmental Tech', 'Female', '1997-11-13'),
    ('Osbourn', 'Caen', 'ocaen3@ning.com', 'Statistician I', 'Master Degree', 'Environmental Specialist', 'Male', '1999-05-28'),
    ('Addison', 'Noddle', 'anoddle4@a8.net', 'Marketing Assistant', 'Master Degree', 'Software Test Engineer II', 'Male', '1990-08-27'),
    ('Jasmin', 'Targett', 'jtargett5@sun.com', 'Nurse Practitioner', 'Bachelor Degree', 'Web Designer I', 'Female', '1984-09-21'),
    ('Thaine', 'Osburn', 'tosburn6@deviantart.com', 'Quality Engineer', 'Doctorate Degree', 'Research Associate', 'Male', '1991-08-07'),
    ('Blake', 'Threadkell', 'bthreadkell7@livejournal.com', 'Research Associate', 'Doctorate Degree', 'Social Worker', 'Male', '1998-08-02'),
    ('Carleen', 'Marfe', 'cmarfe8@cargocollective.com', 'Data Coordinator', 'Master Degree', 'Marketing Assistant', 'Female', '1996-03-29'),
    ('Jehu', 'Quittonden', 'jquittonden9@eventbrite.com', 'Structural Engineer', 'Bachelor Degree', 'Compensation Analyst', 'Male', '1981-11-23'),
    ('Dino', 'Malpass', 'dmalpassa@narod.ru', 'Project Manager', 'Bachelor Degree', 'Quality Control Specialist', 'Bigender', '1994-01-30'),
    ('Ram', 'Bilbey', 'rbilbeyb@who.int', 'Product Engineer', 'Master Degree', 'Programmer III', 'Male', '1994-06-22'),
    ('Dulci', 'Langran', 'dlangranc@vinaora.com', 'Office Assistant II', 'Master Degree', 'Nurse', 'Female', '1995-07-17'),
    ('Ollie', 'McTurlough', 'omcturloughd@abc.net.au', 'Data Coordinator', 'Doctorate Degree', 'Design Engineer', 'Male', '1995-10-02'),
    ('Ingar', 'Malloch', 'imalloche@umich.edu', 'Recruiter', 'Doctorate Degree', 'Help Desk Technician', 'Male', '1999-06-04'),
    ('Alessandro', 'Gooders', 'agoodersf@omniture.com', 'Desktop Support Technician', 'Master Degree', 'VP Quality Control', 'Male', '1987-10-08'),
    ('Gerrard', 'Garrit', 'ggarritg@yelp.com', 'Occupational Therapist', 'Bachelor Degree', 'Data Coordinator', 'Male', '1994-05-23'),
    ('Chicky', 'McCambridge', 'cmccambridgeh@blog.com', 'Graphic Designer', 'Bachelor Degree', 'Accounting Assistant III', 'Female', '1998-02-26'),
    ('Jedediah', 'Rowcliffe', 'jrowcliffei@globo.com', 'Environmental Tech', 'Master Degree', 'Web Developer II', 'Male', '1998-09-19'),
    ('Lou', 'Kyffin', 'lkyffinj@time.com', 'Nuclear Power Engineer', 'Master Degree', 'Web Designer III', 'Male', '1990-05-18'),
    ('Solomon', 'Fido', 'sfidok@yahoo.co.jp', 'Associate Professor', 'Doctorate Degree', 'Social Worker', 'Male', '1983-03-11'),
    ('Cecilius', 'Wildes', 'cwildesl@yandex.ru', 'Compensation Analyst', 'Doctorate Degree', 'Environmental Specialist', 'Male', '1987-10-07'),
    ('Eric', 'Petschel', 'epetschelm@liveinternet.ru', 'Chemical Engineer', 'Doctorate Degree', 'Paralegal', 'Male', '1982-08-04'),
    ('Sylvester', 'Eyree', 'seyreen@berkeley.edu', 'Project Manager', 'Master Degree', 'Clinical Specialist', 'Male', '1989-06-06'),
    ('Wylie', 'Dumphreys', 'wdumphreyso@blogtalkradio.com', 'Editor', 'Master Degree', 'Librarian', 'Male', '1999-02-02');
-- Seed for Staff table
INSERT INTO Staff (FirstName, Lastname, Email, Profession, Sex, DOB, HireDate, DepartmentID)
VALUES
    ('Sarah', 'Williams', 'sarah.williams@example.com', 'HR Specialist', 'Female', '1985-07-18', '2022-03-10', 1),
    ('Michael', 'Jones', 'michael.jones@example.com', 'System Administrator', 'Male', '1990-12-05', '2021-11-15', 2),
    ('Emily', 'Brown', 'emily.brown@example.com', 'Sales Representative', 'Female', '1993-04-30', '2023-02-05', 3),
    ('David', 'Miller', 'david.miller@example.com', 'Financial Analyst', 'Male', '1987-09-25', '2022-05-20', 4);

INSERT INTO staff_info (
    contract_type,
    gross_salary,
    net_salary,
    bank_account_rib,
    social_account_code,
    title,
    advantages,
    vacation_percentage,
    StaffID
) VALUES
    ('CDI', 60000.00, 48000.00, '12345678901234567890', 'SSN123456789', 'Senior', true, 15, 1),
    ('SIVP', 45000.00, 36000.00, '09876543210987654321', 'SSN987654321', 'Junior', true, 12, 2),
    ('CDI', 75000.00, 60000.00, '56789012345678901234', 'SSN345678901', 'Senior', true, 18, 3),
    ('CDD', 50000.00, 40000.00, '54321098765432109876', 'SSN678901234', 'Senior', false, 25, 4),
    ('CDI', 90000.00, 72000.00, '98765432109876543210', 'SSN901234567', 'Junior', true, 20, 5);
