﻿# SSFilm_Venue_Reservation
### Overview
This is a project for Soongsil Film students to reserve venues and manage reservations.

## Stacks
### Infra
- AWS - API gateway, Lambda, SQS
- IaC - Terraform
### Backend
- Go
- Python
- DynamoDB
### Frontend
- Svelte

## Architecture
insert image here

## Main features
🗒️ Reserve venue at specific time period
- limit 6 hours of use per day for each user

📧 Send e-mail notification to users
- for each successful reservation
- for each modified reservation
- for each rejected reservation

👤 Admin functions
- accept reservations
- reject reservations
- set unavailable time periods
