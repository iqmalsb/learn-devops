-- Dump file for table: employees
-- Created for: Enterprise Suite DevOps Testing
-- Data: 10 Dummy Rows

INSERT INTO employees (
    id, first_name, last_name, address_one, address_two, postcode, city, state, 
    date_of_birth, sex, ethnicity, position, department, joined_at, 
    created_at, updated_at, deleted_at
) VALUES 
(1, 'Ahmad', 'Zaki', '12, Jalan Ampang', 'Taman Dato Ahmad', '50450', 'Kuala Lumpur', 'W.P. Kuala Lumpur', '1990-05-15', 'Male', 'Malay', 'Senior Developer', 'Technology', '2020-01-10 09:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(2, 'Sarah', 'Tan', 'No 5, Lorong 2/4A', 'Bandar Puteri', '47100', 'Puchong', 'Selangor', '1992-11-20', 'Female', 'Chinese', 'HR Manager', 'Human Resources', '2021-03-15 08:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(3, 'Ravi', 'Kumar', 'Lot 102, Section 5', 'PJ Old Town', '46000', 'Petaling Jaya', 'Selangor', '1988-02-28', 'Male', 'Indian', 'Payroll Specialist', 'Finance', '2019-06-01 09:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(4, 'Siti', 'Aminah', 'Unit B-10-3', 'Sri Putamas Condominium', '51200', 'Kuala Lumpur', 'W.P. Kuala Lumpur', '1995-07-12', 'Female', 'Malay', 'Finance Executive', 'Finance', '2022-08-20 08:45:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(5, 'John', 'Doe', '88, Jalan SS15/4', 'Subang Jaya', '47500', 'Subang Jaya', 'Selangor', '1991-03-03', 'Male', 'Others', 'Sales Lead', 'Sales', '2023-01-05 10:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(6, 'Nur', 'Aisya', 'No 22, Jalan Kristal', 'Shah Alam', '40000', 'Shah Alam', 'Selangor', '1994-12-05', 'Female', 'Malay', 'Customer Success', 'Support', '2021-11-11 09:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(7, 'Marcus', 'Lim', 'A-3-1, Menara Hartamas', 'Sri Hartamas', '50480', 'Kuala Lumpur', 'W.P. Kuala Lumpur', '1985-09-30', 'Male', 'Chinese', 'Project Manager', 'Events', '2018-05-20 09:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(8, 'Vanesa', 'Anand', '7, Jalan Bangsar', 'Bangsar Baru', '59100', 'Kuala Lumpur', 'W.P. Kuala Lumpur', '1993-01-14', 'Female', 'Indian', 'HR Assistant', 'Human Resources', '2022-02-01 08:30:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(9, 'Lee', 'Wei Kiat', 'No 45, Jalan Dedap', 'Taman Johor Jaya', '81100', 'Johor Bahru', 'Johor', '1990-06-25', 'Male', 'Chinese', 'Backend Engineer', 'Technology', '2020-07-15 09:15:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(10, 'Fatimah', 'Zahra', 'Lot 15, Kg Baru', 'Jalan Raja Alang', '50300', 'Kuala Lumpur', 'W.P. Kuala Lumpur', '1996-04-18', 'Female', 'Malay', 'Admin Clerk', 'Administrative', '2023-05-10 08:45:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);