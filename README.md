# go-rest-api-basic-structure

-- go mod init
-- go mod tidy
-- go get github.com/gin-gonic/gin
-- go run .


## will finish this later 
### structure 

---> main func calls app.StartApp in app package
    ---> app.StartApp should call another mapUrl in seperate file in app package
        ---> StartApp should call the controller
            ---> controller calls service