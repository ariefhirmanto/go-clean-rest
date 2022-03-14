# Golang Clean Architecture for REST API
Sample Golang Clean Architecture for REST API

## **Tools used**
* Web Framework: Gin
* Database: MySQL
* Cache: Redis
* Config file: Viper
 
## **Project Description**
REST API for posting some content/message with CRUD (Create, Read, Update, Delete) capability added with caching feature.

### **Structure**
Based on repository pattern, this project use:
* Repository layer: For accessing db in the behalf of project to store/update/delete data
* Usecase layer: Contains set of logic/action needed to process data/orchestrate those data
* Models layer: Contains set of entity/actual data attribute
* Controller layer: Acts to mapping users input/request and presented it back to user as relevant responses

## **API Endpoints**
### **POST /api/v1/post**
Creates a new post
#### **Examples**:
```
{
    "title": "testing",
    "content": "contents",
    "image_url": "http://an.image",
    "category": "testing_post"
}
```

### **GET /api/v1/post**
Get all posts
#### **Examples responses**:
```
[
  {
    "id": 1,
    "title": "testing",
    "slug": "testing-1",
    "content": "contents",
    "image_url": "http://an.image",
    "category": "testing_post"
  }
]
```

### **GET /api/v1/post/:id**
Get post by id
#### **Examples request**:
> http://localhost:5000/api/v1/post/1
#### **Examples response**:
```
{
    "id": 1,
    "title": "testing",
    "slug": "testing-1",
    "content": "contents",
    "image_url": "http://an.image",
    "category": "testing_post"
}
```

### **GET /api/v1/post/slug/:slug**
Get post by slug
#### **Examples request**:
> http://localhost:5000/api/v1/post/slug/testing-1
#### **Examples response**:
```
{
    "id": 1,
    "title": "testing",
    "slug": "testing-1",
    "content": "contents",
    "image_url": "http://an.image",
    "category": "testing_post"
}
```

### **GET /api/v1/post/title/:title**
Get post by title
#### **Examples request**:
> http://localhost:5000/api/v1/post/title/testing
#### **Examples response**:
```
{
    "id": 1,
    "title": "testing",
    "slug": "testing-1",
    "content": "contents",
    "image_url": "http://an.image",
    "category": "testing_post"
}
```

### **DELETE /api/v1/post/:id**
Delete post by ID
#### **Examples request**:
> http://localhost:5000/api/v1/post/1

### **PUT /api/v1/post/:id**
Update post by ID
#### **Examples request**:
```
{
    "id": 1,
    "title": "testing",
    "slug": "testing-1",
    "content": "contents",
    "image_url": "http://an.image",
    "category": "testing_post"
}
```
#### **Examples response**:
```
{
    "message": "Post successfully updated!"
}
```