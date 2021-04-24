<h1 align="center">Simple RESTful API Hiring Channel App</h1>
<p align="center">
  <img height="200" src="https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png"/>
  <img height ="200" src="https://download.logo.wine/logo/Redis/Redis-Logo.wine.png"/>
</p>

## Table of Contents

- [Built With](#built-with)
- [Introduction](#introduction)
- [Requirements](#requirements)
- [Usage](#usage-for-development)
- [Create Environment Variable](#create-environment-variable)
- [Contact](#contact)
- [Contributors](#contributors)

## Built With

[![Redis](https://img.shields.io/badge/Redis-v%206.15.9-%23D6062)](https://redis.io/)
[![golang](https://img.shields.io/badge/go-v%201.16-skyblue)](https://golang.org/)

## Introduction

<b>Hiring Channel App</b> is a feature that allows Engineers and Company/Employers to do communicate for looking/hiring jobs. This project for handle backend only. Built using Gorm Framework.
Gorm is a web application framework for Golang.

## Requirements

1. <a href="https://www.npmjs.com/">go</a>
2. <a href="https://www.getpostman.com/">Postman</a>
3. Web Server (ex. localhost)
4. MySql
5. [Redis](https://redis.io/)

## Usage for development

1. Open app's directory in CMD or Terminal
2. Type `npm install`
3. Make new file a called **.env**, set up first [here](#create-environment-variable)
4. Turn on Web Server and MySQL can using Third-party tool like xampp, etc.
5. Create a database with the name db_hiring_channel_app, and Import file [db_hiring_channel_app.sql](db_hiring_channel_app.sql) to **phpmyadmin**
6. Open Postman desktop application or Chrome web app extension that has installed before
7. Choose HTTP Method and enter request url.(ex. localhost:3000/engineers)
8. You can see all the end point in routes folder

## Create Environment Variable

```
$ cp .env.example .env
$ nano .env
```

```
# Set port
PORT= YOUR-PORT
# Set database connection information
DB_HOST='YOUR DB HOST'
DB_USER='YOUR DB USER'
DB_PASS='YOUR DB PASSWORD'
DB_NAME='YOUR DB NAME'
# Set JWT Secret
JWT_SECRET= YOUR-SECRET-KEY
```

## Contact

If you want to contact me you can reach me at <shoifuddin.arkademy@gmail.com>.

Copyright © 2021 by Iip Shoifuddin

## Contributors

<center>
  <table>
    <tr>
      <td align="center">
        <a href="https://github.com/iipshoifuddin">
          <img width="100" src="https://github.githubassets.com/images/modules/logos_page/Octocat.png" alt="Iip Shoifuddin"><br/>
          <sub><b>Iip Shoifuddin</b></sub>
        </a>
      </td>
    </tr>
  </table>
</center>
