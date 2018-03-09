insert into`commissiontotals`(`CommissionTotalID`,`InvoiceID`,
    `ManagingCompanyID`,`Percentage`,`DateTimeAdded`)
select 0x11E823CE831B78238D060e30096EF3AA`BuuID`,
'200fc7d0-1e6c-11e8-a382-12237f224cef' `ID`,
`networks`.`ManagingCompanyID`,
`salesteamnetworks`.`Percentage`,
now()
from`companies`
left join`salesteams`using(`SalesTeamID`)
left join`salesteams` `currentsalesteams`on`currentsalesteams`.`ParentSalesTeamID`in(
    `salesteams`.`SalesTeamID`,`salesteams`.`ParentSalesTeamID`)
and`currentsalesteams`.`Deleted`=0 
left join`salesteamnetworks`on`salesteamnetworks`.`SalesTeamID`=ifnull(
    `currentsalesteams`.`SalesTeamID`,`salesteams`.`SalesTeamID`)
join`networks`on`networks`.`NetworkID`=`salesteamnetworks`.`NetworkID`
where`CompanyID`='5a0958d4-04a4-11e5-a03b-3417ebdfde80';