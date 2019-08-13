# wall-street-demo

wallStreetFirm

WEB_ADDRESS=;
RABBIT_MQ_HOST=;
RABBIT_MQ_PORT=;
RABBIT_MQ_USER=;
RABBIT_MQ_PASSWORD=;
RABBIT_MQ_TFC=


wallStreetCustomer

WEB_ADDRESS=;
RABBIT_MQ_HOST=;
RABBIT_MQ_PORT=;
RABBIT_MQ_USER=;
RABBIT_MQ_PASSWORD=;
RABBIT_MQ_TFC=

Identified problems
* Huge concurrency factor
* Latency in queries of stocks
* Periods of outages
* Delays in internal transactions
* Traders disatisfaction

Proposed Patterns
* Pattern reply correlation point-to-point for transactions and stocks
* Publish-Suscriber for trade forecasting

C4 Model
https://www.draw.io/#W87fbfff769e630f4%2F87FBFFF769E630F4!192

