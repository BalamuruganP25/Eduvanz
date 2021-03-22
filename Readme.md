 Prerequisites 

 OS - Ubuntu  
1.Golang version go1.15 required

2.Before Run program please follow the steps
2.1.install golang  (if your already install golang please skip the step)
    Note  - User - root
   
    Download  Golang package 
		https://golang.org/doc/install
	Unzip the download folder 
   	   tar -xzf golangpackage name 
	Copy go folder 
		cp -rf  go /user/local
	Set golang path 
 		vim ~/.bashrc
	Add the below line in the bashrc 
		export GOROOT=/usr/local/go
		export GOPATH=$HOME/Goprojects
		export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
	Save the file
		source  ~/.bashrc
	Check go version in comment line
		
		go version 
	 Out put 
		go version go1.15 linux/amd64

 2.2 Add environment variable in bashrc file 
    vim ~/.bashrc

	 export  Eduvanz_DATABASENAME=eduvanz
	 export  Eduvanz_DATABASEPASSWORD=abcd1234
	 export  Eduvanz_DATABASEUSER=postgres
	 export  Eduvanz_DATABASEPORT=5432
	 export  Eduvanz_DATABASEHOST=localhost
	 export  Eduvanz_API_IP=192.168.43.100
	 export  Eduvanz_API_PORT=8081
	 export  Eduvanz_LOG_FOLDER_NAME=eduvanz_log
	 export  Eduvanz_LOG_FILEPATH=/home/
	 export  GIN_MODE=release

	
	source  ~/.bashrc

3.Unit Test
   Go to server folder 
     path - Eduvanz/src/server 
     please run below comments 
      go test
     OutPut should run with out error  
       
	
5..Install postgres database 
 	3.1 Install postgres database 
 	3.2 Import eduvanz database 
 		 sql file present in the pkg folder

4 .Run the program  
    
    Go to project src folder 
    go run main.go
 
 

