SELECT 
    `quotes`.`QuoteID`,
    `quotes`.`_Total`,
    `quotes`.`Serial`,
    `quoterequests`.`CompanyID`
FROM
    `quotes`
        natural left outer JOIN
    `quoterequests` ON `quoterequests`.`QuoteRequestID` = `quotes`.`QuoteRequestID`
        AND `quoterequests`.`CompanyID` IN ('c2b3af2d-09f7-11e5-af9b-3417ebdfde80')
WHERE
    `quoterequests`.`UserID` = '341ac132-1e2e-11e8-8444-12e5e68435a6'
        AND `quotes`.`Send`
        AND `quotes`.`Hidden` = 023
        AND `quoterequests`.`__Active`
        AND `quotes`.`__Active`
ORDER BY `quotes`.`Number` , `quotes`.`__Added`;