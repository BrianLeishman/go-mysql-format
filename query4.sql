update`invoiceitems`
set`ProductID`=
case`InvoiceItemID`
when'9b55a650-23ec-11e8-a6f3-12237f224cef'then'6d12be0d-1e60-11e5-9da7-
3417ebdfde80'
when'9b55a9fc-23ec-11e8-957b-12237f224cef'then'bf36ecc2-2a8a-11e5-a116-
3417ebdfde80'
when'9b55ab32-23ec-11e8-939a-12237f224cef'then'bf370ab6-2a8a-11e5-a116-
3417ebdfde80'
end,`Price`=
case`InvoiceItemID`
when'9b55a650-23ec-11e8-a6f3-12237f224cef'then'2.82'
when'9b55a9fc-23ec-11e8-957b-12237f224cef'then'0'
when'9b55ab32-23ec-11e8-939a-12237f224cef'then'0'
end,`Quantity`=
case`InvoiceItemID`
when'9b55a650-23ec-11e8-a6f3-12237f224cef'then'100'
when'9b55a9fc-23ec-11e8-957b-12237f224cef'then'1'
when'9b55ab32-23ec-11e8-939a-12237f224cef'then'1'
end,`Size`=
case`InvoiceItemID`
when'9b55a650-23ec-11e8-a6f3-
12237f224cef'then'1.25\"'
when'9b55a9fc-23ec-11e8-957b-12237f224cef'then''
when'9b55ab32-23ec-11e8-939a-12237f224cef'then''
end,`Description`=
case`InvoiceItemID`
when'9b55a650-23ec-11e8-a6f3-12237f224cef'then'Custom Lapel Pins'
when'9b55a9fc-23ec-11e8-957b-12237f224cef'then'Custom Mold Fee'
when'9b55ab32-23ec-11e8-939a-12237f224cef'then'UPS International Shipping'
end,`InvoiceShippingOptionID`=
case`InvoiceItemID`
when'9b55a650-23ec-11e8-a6f3-12237f224cef'then 0x11E823EC9B57FED2943612237F224CEF 
when'9b55a9fc-23ec-11e8-957b-12237f224cef'then 0x11E823EC9B57FED2943612237F224CEF 
when'9b55ab32-23ec-11e8-939a-12237f224cef'then 0x11E823EC9B57FED2943612237F224CEF 
end 
where`InvoiceItemID`in('9b55a650-23ec-11e8-a6f3-12237f224cef','9b55a9fc-23ec-11e8-957b-12237f224cef','9b55ab32-23ec-11e8-939a-12237f224cef');