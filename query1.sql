insert ignore into`quoterequests`(`QuoteRequestID`,`CompanyID`,`UserID`,
    `Company`,`FirstName`,`LastName`,`AddressLine1`,`AddressLine2`,`City`,`CountryID`,
    `StateID`,`Zip`,`Phone`,`Email`,`__Added`)
values
(0x11E823D9D3B9195484331206BB36A2DB,'3e55d1bb-d8b6-11e4-b38f-b8ca3a83b4c8',
    'd3b2cd24-23d9-11e8-b3e7-1206bb36a2db',null,'Jeremiah','Smith',null,null,null,
    'a26f5174-047e-11e5-8309-3417ebdfde80','4a7f2709-0480-11e5-8309-3417ebdfde80',
    null,'+1 719-433-8314','propaintball826@gmail.com','2018-03-09 15:38:00')
on duplicate key update
`Company`=values(`Company`),
`FirstName`=values(`FirstName`),
`LastName`=values(`LastName`),
`AddressLine1`=values(`AddressLine1`),
`AddressLine2`=values(`AddressLine2`),
`City`=values(`City`),
`CountryID`=values(`CountryID`),
`StateID`=values(`StateID`),
`Zip`=values(`Zip`),
`Phone`=values(`Phone`),
`Email`=values(`Email`)