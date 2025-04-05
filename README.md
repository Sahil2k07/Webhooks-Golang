# Webhooks-Golang

This is a simple and to the point implementation of Webhooks in golang.

## Tech

- `Golang`: Backend development
- `Echo`: Web framework for handling APIs

## Description

The Project has two Serviers

1. `Main Server`

- Exposes two key endpoints:

  - `POST /api/account/upgrade`: Simulates a user upgrading their account (e.g., after making a payment).

  - `PUT /api/account/upgrade/webhook`: Acts as the webhook receiver that gets called by the external payment processor and enables premium features for the user.

2. `Webhook Server`

- Mimics a third-party service (like a payment gateway).

- Accepts a payment request and makes a webhook call back to the Main Server’s /webhook endpoint after processing the payment.

This setup demonstrates how webhooks are typically used to communicate between services asynchronously, with one service triggering logic on another by making an HTTP request.
