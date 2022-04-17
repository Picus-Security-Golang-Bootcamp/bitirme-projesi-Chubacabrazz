# E-Commerce App in GOLANG
( Development of App incomplete due to illness =( , I will complete it when i have time )

### Build with Go. Used Gorm , PostgreSQL , Godotenv , Gin Gonic , JWT , Zap , Viper . Also used Postman.

This is a basket app for basic usage. It has products and product categories which people register then do shopping. It has authentication tools such as "jwt/utils.go , helper/token.go". Only admin user has permission to create categories and products. Customers can search items by name,sku than add them to their basket, every basket is unique for that user.
Customers can also see order history and they can cancel their order if 14 days have not passed.


For CRUD operations from db I used ; Gorm,PostgreSQL. For endpoint handling I used ; Gin Gonic. For user authentication I used; JWT. For config implementation I used ; Viper. For testing API I used ; Postman. For logging I used ; Zap.

### Before use

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv
go get -u go.uber.org/zap
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/viper



For checking APIs, I used Postman. For use your url should look like <yourHost>/api/v1/picus-storeApp/


### App has multiple endpoints to communicate, /user /cart /category etc.
### For creating bulk categories as admin I used concurrent csv reading logic to fasten process.

#### I will be grateful for all the contributions. Please feel free to contact me.


