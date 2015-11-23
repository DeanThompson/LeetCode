/* https://leetcode.com/problems/second-highest-salary/ */

/* 1. */
SELECT (SELECT DISTINCT Salary
        FROM Employee
        ORDER BY Salary DESC
        LIMIT 1 OFFSET 1) AS second;

/* 2. */
SELECT MAX(Salary)
FROM Employee
WHERE Salary NOT IN (SELECT MAX(Salary)
                     FROM Employee);