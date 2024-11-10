

# Credit Card Validator & Generator

## Features

### 1. Validate Credit Card Number
Check if a given credit card number is valid using the Luhn algorithm.
Support for both specific card numbers and reading input from stdin.
### 2. Generate Credit Card Number
Generate valid credit card numbers based on a given card format.
Option to generate random numbers or pick a valid number from a provided pattern.
### 3. Issue Credit Card
Generate a valid credit card number based on a brand and issuer.
Requires external files (brands.txt and issuers.txt).
### 4. Card Information Retrieval
Retrieve and display information about a card (brand, issuer) from external files.


# *Example Workflow*

## *Example 1: Validate a Card*
`./creditcard validate 411111111111111111`
## *Example 2: Generate a Card*
`./creditcard generate 41111111111111****`

## *Example 3: Get Card Information*
`./creditcard information --brands=brands.txt --issuers=issuers.txt 4111111111111111`

## *Example 4: Issue a Card*
`./creditcard issue --brands=brands.txt --issuers=issuers.txt Visa BankName`
