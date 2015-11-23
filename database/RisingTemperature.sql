/* https://leetcode.com/problems/rising-temperature/ */

SELECT a.Id
FROM Weather a, Weather b
WHERE
  a.Date = DATE_ADD(b.Date, INTERVAL 1 DAY) AND a.Temperature > b.Temperature;