SELECT p.firstName, p.lastName, a.city, a.state
FROM Person As p
         LEFT JOIN
     Address As a
     ON a.personId = p.personId;