/* https://leetcode.com/problems/consecutive-numbers/ */

SELECT DISTINCT a.Num
FROM Logs a, Logs b, Logs c
WHERE b.Id = a.Id + 1 AND c.Id = a.Id + 2 AND a.Num = b.Num AND a.Num = c.Num;