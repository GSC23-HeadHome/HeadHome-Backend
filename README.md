<div align="center">
    <div >
        <img width="200px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fheadhome_square.png?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/>
    </div>
    <div >
            <p style="font-size:50px;"><b>HeadHome</b></p>
            <p style="font-size:18px"><i>Your companion, every step of the way</i></p>
    </div>      
</div>
<br>

<h1 align="center">HeadHome Backend</h1>
The <b>HeadHome backend</b> is responsible for real-time interactions between the dementia patients, caregivers and volunteers. This ensures that dementia patients can receive timely assistance from caregivers or nearby volunteers and help them to head home safely!
<br>
<h2>👨🏻‍💻 Technology Stack</h2>
<br />
<div align="center">
  <a href="https://go.dev/">
      <kbd>
      <img src="./assets/icons/Go.png" height="60" />
      </kbd>
    </a>
     <a href="https://firebase.google.com/">
      <kbd>
      <img src="./assets/icons/Gin.png" height="60" />
      </kbd>
    </a>
    <a href="https://firebase.google.com/">
      <kbd>
      <img src="./assets/icons/Firebase.png" height="60" />
      </kbd>
    </a>
    <a href="https://mapsplatform.google.com/">
      <kbd>
      <img src="./assets/icons/Maps.png" height="60" />
      </kbd>
    </a>
    <a href="https://cloud.google.com/">
      <kbd>
      <img src="./assets/icons/GCP.png" height="60" />
      </kbd>
    </a>
    <br>
    <h4>Go | Gin | Firebase | Google Maps Platform | Google Cloud Platform</h4>
</div>

<br>

# Getting Started

This repo contains the backend application for HeadHome. Click [here](https://github.com/GSC23-HeadHome/HeadHome) to view the full solution.
<br><br>
[Go `(Version 1.19+)`](https://go.dev/doc/install) must be installed to run this application.

## ⚙️ &nbsp;Steps to Setup

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
   Note: Place the entire Firebase Admin SDK private key json object on a single line and escape all `\`, `\n` and quotation(") characters with `\`. Lastly, surround the json object with double quotations `""`.
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

## 🔑 &nbsp; Files and Directories

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
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── replit.nix
```

<br>

## 👥 &nbsp;Contributors

| <a href="https://github.com/chayhuixiang"><img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fhuixiang.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/></a> | <a href="https://github.com/changdaozheng"><img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fdaozheng.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/></a> | <a href="https://github.com/Trigon25"><img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fmarc.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/></a> | <a href="https://github.com/ongjx16"><img width="180px" src="https://firebasestorage.googleapis.com/v0/b/gsc23-12e94.appspot.com/o/members%2Fjingxuan.jpeg?alt=media&token=96a55b42-7c9f-4e68-b41f-d986efe79c01" alt=""/></a> |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <div align="center"><h3><b><a href="https://github.com/chayhuixiang">Chay Hui Xiang</a></b></h3><p><i>Nanyang Technological University</i></p></div>                                                                               | <div align="center"><h3><b><a href="https://github.com/changdaozheng">Chang Dao Zheng</a></b></h3></a><p><i>Nanyang Technological University</i></p></div>                                                                          | <div align="center"><h3><b><a href="https://github.com/Trigon25">Marc Chern Di Yong</a></b></h3></a><p><i>Nanyang Technological University</i></p></div></a>                                                               | <div align="center"><h3><b><a href="https://github.com/ongjx16">Ong Jing Xuan</a></b></h3></a><p><i>Nanyang Technological University</i></p></div>                                                                            |
