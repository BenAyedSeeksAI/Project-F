DROP TABLE IF EXISTS Staff;
DROP TABLE IF EXISTS Department;
DROP TABLE IF EXISTS Candidate;
CREATE TABLE Department (
			DepartmentID INTEGER PRIMARY KEY AUTOINCREMENT,
			DepartmentCode TEXT NOT NULL,
			DepartmentName TEXT NOT NULL,
			Location TEXT
			);
CREATE TABLE Staff (
			StaffID INTEGER PRIMARY KEY AUTOINCREMENT,
			FirstName TEXT NOT NULL,
			LastName TEXT NOT NULL,
			Email TEXT UNIQUE NOT NULL,
			Profession TEXT,
			Sex TEXT NOT NULL,
			DOB TEXT,
			Username TEXT,
			Password TEXT,
			PhoneNumber TEXT,
			HireDate TEXT,
			DepartmentID INT,
			CONSTRAINT FK_Staff_Department FOREIGN KEY (DepartmentID) REFERENCES Department(DepartmentID)
			);
		CREATE TABLE Candidate (
            CandidateID INTEGER PRIMARY KEY AUTOINCREMENT,
			FirstName TEXT NOT NULL,
			LasttName TEXT NOT NULL,
			PersonalEmail TEXT NOT NULL,
			JobOffer TEXT NOT NULL,
			Degree TEXT NOT NULL,
			RecentExperience TEXT NOT NULL,
			Sex TEXT NOT NULL,
			DOB TEXT
		);
