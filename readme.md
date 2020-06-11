**CanITravelTo**

Libraries used:
Gin
mysql-connector


Passport Info retrieval: (passportindex.com)

`https://github.com/ilyankou/passport-index-dataset`

Converted to mysql with: 

`https://www.convertcsv.com/csv-to-sql.htm`

If updated, and the list of countries has changed, it must also be changed in Countries.go list (or remove that list and just sanitise input to prevent sqlInjection)