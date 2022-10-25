# Dashboard Stocks

This is the backend for the stocks component for a custom dashboard project. The goal of this project is to have a fully customizeable dashboard that will load of system startup. Eventually this will be containerized and deployed to an AWS (EC2) instance. 

## Problem 

Most "Dashboard" style applications that you can download on Mobile/PC/MAC are awkward and clunky or are overpriced with not many customizeable features.

## Built with 

1. [GO v1.19.2](https://go.dev/) 
2. [Polygon Official GO Client](https://github.com/polygon-io/client-go)

## Resources

1. [RapidAPI RealStonks](https://rapidapi.com/amansharma2910/api/realstonks/) create an account, subscribe to RealStonks using basic tier. copy custom API key for data 

# Getting started

1. Make sure you have the latest version og GO installed.
2. 


# Problemms/Notes

Finding a free API to retrive current stock prices is more difficult then originally anticipated.

Polygon.io was the first choice but I quickly realized they do not offer prices, only historical market data with the free tier. 

Google, Yahoo and most major finance companies do not offer public APIs

## Options

1. Write a web scrapper instead
    Pros: 
        a. can query prices from almost any financial institution. (Google might be best bet as a custom query should be realively easy.)

    Cons: 
        a. Much more difficult to implement. 
        b. Webscrapers tend to be clunky and use a lot of resources. 
        c. increased compile time.
        
2. Use RapidAPI
    Pros:
        a. easily connect to 
         