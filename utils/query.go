package utils

// task-05 Query
var Query = `
SELECT d.name AS Department, e1.name AS Employee, e1.salary AS Salary
FROM employee e1
JOIN department d ON e1.departmentId = d.ID
WHERE (
    SELECT COUNT(DISTINCT e2.salary)
    FROM employee e2
    WHERE e2.departmentId = e1.departmentId AND e2.salary >= e1.salary
) < 4
ORDER BY Department, Salary DESC;
`

var Statmentarray = []string{
	`
CREATE TABLE IF NOT EXISTS department (
    ID SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
)`,
	`
INSERT INTO department (ID, name) VALUES
    (1, 'IT'),
    (2, 'Sales');
`, `
CREATE TABLE IF NOT EXISTS employee (
    ID SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    salary INT NOT NULL,
    departmentId INT,
    FOREIGN KEY (departmentId) REFERENCES department(ID)
);
`,
	`
INSERT INTO employee (ID, name, salary, departmentId) VALUES
    (1, 'Joe', 85000, 1),
    (2, 'Henry', 80000, 2),
    (3, 'Sam', 60000, 2),
    (4, 'Max', 90000, 1),
    (5, 'Janet', 69000, 1),
    (6, 'Randy', 85000, 1),
    (7, 'Will', 70000, 1);
`,
}
