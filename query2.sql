select `shippingbills`.`TrackingNumber`,`shippingbills`.`Amount`,
`shippingbills`.`Customer`,`shippingbills`.`ReferenceLine1`,
`shippingbills`.`ReferenceLine2`,`invoices`.`InvoiceNumber`
from `shippingbills` 
join `shippinginvoicefiles` using (`ShippingInvoiceFileID`)
join `invoices` on `invoices`.`InvoiceNumber`=`shippingbills`.`PossibleInvoiceNumber`
and `invoices`.`__Active`=1
left join `invoicetrackingnumbers` `invoicetrackingnumbers` on `invoicetrackingnumbers`.`InvoiceID`=`invoices`.`InvoiceID` 
and `invoicetrackingnumbers`.`Active`=1
join `shippingaccounts` on `invoices`.`ShippingAccountID`=`shippingaccounts`.`ShippingAccountID`
where `invoicetrackingnumbers`.`InvoiceTrackingNumberID`is null 
and`shippingbills`.`Amount`>0 
and `shippingbills`.`_InvoiceTrackingNumberID`is null 
and `shippinginvoicefiles`.`AccountNumber`=`shippingaccounts`.`AccountNumber`
and abs(datediff(`invoices`.`DateTimeShipped`,`shippingbills`.`DateShipped`))<=@@MaxDays;