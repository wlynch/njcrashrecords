NJDOT Crash Records
==============

The NJDOT exports a lot of data on [their site](http://www.state.nj.us/transportation/refdata/accident/), but it isn't in a
format that can be easily managed. This project parses this data out into json, so that it can then be exported to other tools such as BigQuery.

You can play around with an example BigQuery table [here](https://bigquery.cloud.google.com/dataset/njcrashrecords:njcrashrecords). I currently only export Middlesex 2013 data because I'm still making changes, and it's a pain to parse and upload ~8GB of data for the entire set. When the schema becomes more stable I'll roll out the complete set (v1 should be soon).
