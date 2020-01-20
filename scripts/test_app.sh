#!/bin/bash

curl http://localhost:8081/webhook/Monzo2Ynab -X POST --data @monzo_example.json -vvv
