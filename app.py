#!/usr/bin/env python
# -*- coding: utf-8 -*-
#

import requests
import json
import redis
import time

# get data from conibase
def getRates(): 
  results = []
  urls = [ 
           "https://api.coinbase.com/v2/prices/BTC-CZK/buy",
           "https://api.coinbase.com/v2/prices/BCH-CZK/buy",
           "https://api.coinbase.com/v2/prices/LTC-CZK/buy",
           "https://api.coinbase.com/v2/prices/ETH-CZK/buy"
         ]
  headers = { "CB-VERSION": "2018-03-04" }

  for url in urls:
    res = requests.get(url, headers=headers)
    results.append(res.json())

  return results

def parseRates(rates):
  datas = ""
  for rate in rates:
    ckey = '"' + rate['data']['base'] + '"'
    cval = '"' + rate['data']['amount'] + '"'
    data = ckey + ":" + cval

    if datas:
      datas = datas + "," + data
    else:
      datas = data

  datas = "{" + datas + "}"
      
  return datas

def putRatesToCache(redis_host, redis_key, rates):
  try:
    cache = redis.StrictRedis(host=redis_host, port=6379, db=0)
    cache.set(redis_key, rates)
  finally:
    del cache

  return 0

while True:
  rates = getRates()
  rates = parseRates(rates)
  putRatesToCache('127.0.0.1', 'exchange_rates', rates)
  time.sleep(300)

exit(2)
