SELECT
   p1."ID",
   p1."UserName",
   (
      SELECT
         p2."UserName"
      FROM
         users p2
      WHERE
         p2."ID" = p1."Parent") AS "ParentUserName"
FROM
   users p1
