# HeadHome Backend
<div align="center">
    <div >
        <img width="200px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fheadhome_square.png?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/>
    </div>
    <div >
            <b style="font-size:25px">HeadHome</b>
            <br> 
            <p style="font-size:18px">Dementia Aid Solution</p>
    </div>    
        
</div>

<br>

# Getting Started
This repo contains the backend application for HeadHome. Click [here](https://github.com/GSC23-HeadHome/HeadHome) to view the full solution.

<br>

## Prerequisites
[Go `(Version 1.19+)`](https://go.dev/doc/install) must be installed to run this application.

<br>

## Tech Stack

<div align="center">
  <p>
    <a href="https://go.dev/">
      <img width="100px" src="https://skillicons.dev/icons?i=go" />
    </a>
    <a href="https://firebase.google.com/">
      <img width="100px" src="https://skillicons.dev/icons?i=firebase" />
    </a>
  </p>
</div>

## Quick Start:
1. Clone Repo
```
$ git clone https://github.com/GSC23-HeadHome/HeadHome-Backend.git
```
2. Run the following code in bash to install the required dependencies
```
$ go get all
```
3. Create a `.env` file and insert your Firebase Admin SDK private key and Maps API api key. 
<br>
<font color="#888888">
    Note: Place the entire Firebase Admin SDK private key json object on a single line and escape all `\` and `\n` characters with `\`. Lastly, surround the json object with double quotations `""`.
</font>

```css
/*.env file*/
FIREBASE_ADMIN_PRIVATE_KEY=<your inline firebase admin private key>
MAPS_API_KEY=<your maps api key>
```
4. Launch the server at `0.0.0.0:8080`
```
$ go run ./cmd
```

<br>

## File structure:

```tree
├── cmd
│   └──main.go
├── controllers
│   ├── care_giver_controller.go
│   ├── care_receiver_controller.go
│   ├── map_controller.go
│   ├── sos_controller.go
│   ├── travel_log_controller.go
│   └── volunteers_controller.go
├── database
│   ├── care_giver_collection.go
│   ├── care_receiver_collection.go
│   ├── database.go
│   ├── sos_log_collection.go
│   ├── travel_log_collection.go
│   └── volunteers_collection.go
├── fcm
│   └── fcm.go
├── firebase_app
│   └── firebase_app.go
├── logic
│   ├── direction.go
│   └── util.go
├── models
│   ├── care_giver.go
│   ├── care_receiver.go
│   ├── sos_log.go
│   ├── travel_log.go
│   └── volunteers.go
├── routes
│   └── routes.go
├── websocket.go
│   ├── client.go
│   ├── msg_pump.go
│   ├── websocket.go
│   └── ws_hub.go
├── .env (not included in github repo)
├── .gitignore
├── .replit
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── replit.nix
```

<br>

# Developers
              
|<img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fhuixiang.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/>|<img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fdaozheng.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/>|<img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fmarc.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/>| <img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fjingxuan.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/>
|--------------------------|--------------------------|--------------------------|--------------------------|
|<div align="center"> <h3><b>Chay Hui Xiang</b></h3><p><i>Nanyang Technological University</i></p></div>|<div align="center"><h3><b>Chang Dao Zheng</b></h3><p><i>Nanyang Technological University</i></p></div>|<div align="center"><h3><b>Marc Chern Di Yong</b></h3><p><i>Nanyang Technological University</i></p></div>|<div align="center"><h3><b>Ong Jing Xuan</b></h3><p><i>Nanyang Technological University</i></p></div>|