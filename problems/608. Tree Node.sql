SELECT
    id,
    CASE
        WHEN p_id IS NULL THEN 'Root'
        WHEN id IN (Select p_id FROM Tree) THEN 'Inner'
        ELSE 'Leaf'
        END As type
FROM Tree;