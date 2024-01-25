# GO-AdminPro

<br />

## 在 GO-AdminPro中、角色和權限之間的關係:

每位管理員可以擁有多個角色，同時每個角色可以擁有多個權限。
<br />
這種架構使得管理員只需要擁有特定的角色，就能夠繼承該角色下的所有權限，實現了高度的模組化管理。
<br />
透過添加或移除角色，可以輕鬆地調整管理員的權限，使權限管理變得直觀而靈活。
<br />
支援對管理員進行額外的權限授予，提供了更細緻的權限管理選項，以滿足不同應用場景的需求。


<br />
<br />

## 後端技術概覽:
-  Bcrypt : 一種用於安全存儲和驗證密碼的哈希函數，特別適用於登錄密碼的加密和驗證過程<br />
-  Sonyflake : Golang 中的全域性唯一 ID 生成器，用於生成唯一的分散式 ID，包含時間戳和機器 ID<br />
-  Gorm : Golang 中的資料庫 ORM 操作庫，提供了簡化的 API 來執行資料庫操作，包括事務支援<br />
-  Viper : 一個配置管理庫，支援讀取不同資料來源和格式的設定文件，提供了方便的方法來處理應用程式配置<br />
-  jwt-go : 一個用於 Golang 的 JSON Web Token (JWT) 實現庫，用於生成和驗證 JWT，可用於實現使用者驗證和生成安全令牌 <br />

<br />
<br />

## 實現功能(API):
- 添加管理員<br />
- 角色添加權限<br />
- 為管理員添加腳色<br />
- 為管理員添加權限<br />
- 查詢所有管理員<br />
- 查詢指定管理員的角色<br />
- 查詢指定管理員額外的權限<br />
- 添加角色<br />
- 查詢指定管理員所有的權限(包含角色)<br />
- 移除角色所屬的權限<br />
- 移除管理員額外的權限<br />
- 移除管理員的角色<br />
- 取得角色權限<br />
- 查詢所有角色<br />
- 查詢所有權限<br />
<br />

## 運行項目:
* 需要先安裝 docker-compose<br />

<br />

### Docker Compose 運行前需要加上虛擬內部網域
docker network create --subnet=192.168.200.0/24 redis-cluster-net<br />

<br />

### Mysql5.7

進到目錄: 
cd GO-AdminPro/docker/goAdminMysql57/docker-compose.yml

執行:
dokcer compose up -d 

完成後會自動添加表和預設data<br />

<br />

### 測試帳號
-最高管理員 admin : 12345678<br />
-超級管理員 manager : 12345678<br />

<br />

### 預設腳色:
- ADMIN : 可以使用全部權限<br />
- SUPER_MANAGER : 除了創建管理員 其他權限全部都可使用<br />
- NORMAL_MANAGER : 只能使用查詢相關權限<br />
- USER : 佔無任何權限<br />

<br />
<br />

## 資源：
PostMan引入檔 : GO-AdminPro/postman/GoAdminPro.postman_collection.json

Mysql 初始 : /Users/sai/GolandProjects/GO-AdminPro/db/initV1.sql

<br />
<br />

## 權限關係總覽圖：
- 管理員可以自由添加或刪除對應角色<br />
- 角色可以自由添加或刪除對應權限<br />
- 可以單獨為某個管理員添加指定權限<br />
![image](https://github.com/lzz0826/GO-AdminPro/blob/main/img/005.png)

<br />
<br />

## 管理員對應角色關係圖：
### 一位管理員可以對應多個角色
![image](https://github.com/lzz0826/GO-AdminPro/blob/main/img/002.png)

<br />
<br />

## 角色對應權限關係圖：
### 一個角色可以對應多個權限
![image](https://github.com/lzz0826/GO-AdminPro/blob/main/img/003.png)

<br />
<br />

## 管理員對應權限關係圖：
### 一位管理員可以對應多個權限(相當於可以開額外給予權限)
![image](https://github.com/lzz0826/GO-AdminPro/blob/main/img/004.png)
