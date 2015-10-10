:: It updates the SQL database Schema

:: Change this to your path
cd ../../../util/ubuntu/install/SQL/

:: Change this to your path
:: Usage: mysqldump -u USERNAME -pPASSWORD --no-create-info DATABASE > Wisply.sql
mysqldump -u wisply -pDNeaMKvz4t4DtL6b --no-create-info wisply > Data.sql
