update`invoices`
set`Name`='Nikki Morton - (#6)',`CompanyID`='88ee3dbf-0a0b-11e5-af9b-3417ebdfde80',`FactoryID`='8429fe44-053b-11e5-af8f-3417ebdfde80',
`ShippingAccountID`='4dac2555-2b0b-11e5-a116-3417ebdfde80',
`DateTimeShipped`='2018-03-19 00:00:00',`DateTimeDelivered`='2018-03-21 00:00:00',
`DateTimeAdded`='2018-03-09 17:52:00',`Notes`='Use bright yellow gold.',
`DateTimeUpdated`=now(),`FrontOfficeNotes`='',`ChargeTeam`='0',`NoInvoice`='0',
`SendProductDescriptions`='1',`Pending`='0',`SalesTeamID`=if(trim(
        `SalesTeamID`)=''
    or`SalesTeamID`is null,'e5557127-1478-11e5-9556-3417ebdfde80',
    `SalesTeamID`),`AdminUserID`=if(trim(`AdminUserID`)=''
    or`AdminUserID`is null,'829e2346-1379-11e5-a091-3417ebdfde80',
    `AdminUserID`),`_Admin`=if(trim(`_Admin`)=''
    or`_Admin`is null,'Tammy Reger',`_Admin`),`SampleRequest`='0',
`_Total`=`_Total`+'0',`_TotalSansTax`=`_TotalSansTax`+'0',
`_ItemProducts`=`_ItemProducts`+'0',`_NonItemProductTotal`=`_NonItemProductTotal`+
'0',`UpdateHash`=unhex(
    'eac0f20e2cd1b4e0082e8cdebfea390ecc639c97ac52359c4dfb80879e12c07c')
where`InvoiceID`='9b54bede-23ec-11e8-a2aa-12237f224cef'
and(`UpdateHash`is null 
    or`UpdateHash`<>unhex(
        'eac0f20e2cd1b4e0082e8cdebfea390ecc639c97ac52359c4dfb80879e12c07c'));