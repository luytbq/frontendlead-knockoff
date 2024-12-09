# FrontendLead Knockoff

## Introduction

This app is a self-hosted clone of [FrontendLead.com](https://frontendlead.com), designed to provide challenging JavaScript coding questions for free.

The inspiration for this project came from my experience with the original app, which restricts free trial runs and frequently prompts users to purchase a membership—something I found frustrating.

To overcome these limitations, I crawled their data and built my own version. By all means, this is a clumsy clone and lacks many features compared to the original app, but it gets the job done for personal use.

---

## Features

- Solve JavaScript coding challenges.
- Host the app locally without any restrictions.
- Extendable: Modify the app to suit your needs.

---

## Prerequisites

Before running the app, ensure you have the following installed:

- Docker
- Go (optional, for crawling questions yourself)

---

## Run the App Locally

1. Clone the repository:
``` bash
git clone https://github.com/luytbq/frontendlead-knockoff.git 
```
2. Change dir to the downloaded folder:
``` bash
cd frontendlead-knockoff
```
3. Build the image:
``` bash
sudo docker build -t frontendlead-knockoff .
```
4. Start docker container:
``` bash
sudo docker run --name frontendlead-knockoff -p 5130:5130 --rm frontendlead-knockoff
```

5. Open your browser and navigate to http://localhost:5130 to start solving challenges!


## Crawl the questions yourself
Crawl the questions yourself by running:
``` bash
go run crawl/main.go
```
Note: you might need to modify the `questionDetailEndpoint` variable. Inspect the network tab on the original app to find correct end point.

## Disclaimer
This project is intended for educational and personal use only. It is not affiliated with or endorsed by FrontendLead.com. Crawling data or replicating functionality from the original app may violate its terms of service. Use this app responsibly.