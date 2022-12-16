# go-rss-reader-service
<ul>
  <li>Open service in your IDE</li>
  <li>Run <i>go mod tidy</i> to install the required packages</li>
  <li>Run the service. The service will run on port 8086</li>
  <li>The service contains POST endpoint which accepts array of urls in the body and in return gives rss feed posts</li>
</ul>
Import the following CURL to call the API. <br>
<code>
curl --location --request POST 'http://{{host}}:{{port}}/rss/reader' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "rss_urls": ["https://feeds.simplecast.com/54nAGcIl", "https://feeds.npr.org/1004/rss.xml"]
}'
</code>
