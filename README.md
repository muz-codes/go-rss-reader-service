# go-rss-reader-service
<h1>Read Rss Feed</h1>
<ul>
  <li>Open service in your IDE</li>
  <li>Run <i>go mod tidy</i> to install the required packages</li>
  <li>Run the service. The service will run on port 8086</li>
  <li>The service contains POST endpoint which accepts array of urls in the body and in return gives rss feed posts</li>
</ul>
Import the following CURL to Postman for calling the Read Rss API. <br>
<code>
curl --location --request POST 'http://{{host}}:{{port}}/rss/reader' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "rss_urls": ["https://feeds.simplecast.com/54nAGcIl", "https://feeds.npr.org/1004/rss.xml"]
}'
</code>

<h1>Add Rss Url To Db</h1>
Sqlite is used to store and fetch URLs.<br>
<p>The url passed is first checked in the db. If the url exists in db, the endpoint will throw the error, and if url is not available in the db, the url will then be validated by calling the url. If response code is in 200 range, the url will be added to the db.</p>
Body Params:
<ul>
<li>url: string, required (in body)</li>
</ul>
<code>
curl --location --request POST 'http://127.0.0.1:8086/rss/url/add' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "url": "http://rss.cnn.com/rss/cnn_topstories.rss"
}'
</code>

<h1>Fetch Rss URLs List From Db</h1>
Fetch rss urls based on page no. and limit per page sorted by created_at either desc or asc.
Query Params:
<ul>
<li>limit: limit per page - int, required</li>
<li>page: page no - int, required</li>
<li>sort: sort by created_at asc OR created_at desc | Default: create_at desc</li>
</ul>

Response:
<ul>
<li>success: bool</li>
<li>message: string</li>
<li>rss_url_list: list of rss urls</li>
<li>total_count: total number of rss urls</li>
<li>total_pages: total number of pages</li>
</ul>
<code>
curl --location --request GET 'http://127.0.0.1:8086/rss/url/list?limit=2&page=1&sort=created_at desc'
</code>

<h1>Update Rss URLs</h1>
<p>Updates url against the given id. The url passed is first checked in the db. If the url exists in db, the endpoint will throw the error, and if url is not available in the db, the url will then be validated by calling the url. If response code is in 200 range, the url will be updated against the given ID.
Body Params:
<ul>
<li>id: int, required (in body)</li>
<li>url: string, required (in body)</li>
</ul>
Response:
<ul>
<li>success: bool</li>
<li>message: string</li>
<li>rss_url :{id, url}</li>
</ul>
<code>
curl --location --request PATCH 'http://127.0.0.1:8086/rss/url/update' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "id": 2,
    "url": "https://feeds.simplecast.com/54nAGcIl"
}'
</code>

<h1>Delete Rss URLs</h1>
<p>Deletes url against the given id if the url is available in the db else the API will throw error.
URI Params:
<ul>
<li>id: int, required (in uri)</li>
</ul>
Response:
<ul>
<li>success: bool</li>
<li>message: string</li>
<li>rss_url :{id, url}</li>
</ul>
<code>
curl --location --request DELETE 'http://127.0.0.1:8086/rss/url/delete/4'
</code>
