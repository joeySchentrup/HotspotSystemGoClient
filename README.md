# HotspotSystemGoClient
An API Client for HotspotSystem API

# API Docs
http://www.hotspotsystem.com/apidocs-v1/

# Functions

/customers
---

Get a list of the resource owner's customers.

**Code**

`Customers(feilds string, sort string, limit string, offset string)`

**Parameters definitions** 

http://www.hotspotsystem.com/apidocs/api/reference/#operation-getlocations

**Returns** 

"ArrayOfCustomers" Model



/locations
---

Get a list of the resource owner's locations

**Code**

`Locations(feilds string, sort string, limit string, offset string)`

**Parameters definitions:**

http://www.hotspotsystem.com/apidocs/api/reference/#operation-getlocations

**Returns:** 

"ArrayOfLocations" Model

/subscribers
---

Get a list of the resource owner's subscribers

**Code**

`Subscribers(feilds string, sort string, limit string, offset string)`

**Parameters definitions:**

http://www.hotspotsystem.com/apidocs/api/reference/#operation-getsubscribers

**Returns:** 

"ArrayOfSubscribers" Model

 /vouchers
---

Get a list of the resource owner's vouchers

**Code**

`Vouchers(feilds string, sort string, limit string, offset string)`

**Parameters definitions:**

http://www.hotspotsystem.com/apidocs/api/reference/#operation-getvouchers

**Returns:** 

"ArrayOfVouchers" Model
