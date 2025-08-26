-- Table: Followers

-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | user_id     | int  |
-- | follower_id | int  |
-- +-------------+------+
-- (user_id, follower_id) is the primary key (combination of columns with unique values) for this table.
-- This table contains the IDs of a user and a follower in a social media app where the follower follows the user.
 

-- Write a solution that will, for each user, return the number of followers.

-- Return the result table ordered by user_id in ascending order.

-- The result format is in the following example.

-- Output: 
-- +---------+----------------+
-- | user_id | followers_count|
-- +---------+----------------+
-- | 0       | 1              |
-- | 1       | 1              |
-- | 2       | 2              |
-- +---------+----------------+

SELECT F.user_id, COUNT(F.user_id)
FROM Followers F
GROUP BY F.user_id
ORDER BY F.user_id