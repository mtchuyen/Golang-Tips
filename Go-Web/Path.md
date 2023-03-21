Một URL có 3 phần
- Path: Path là chung, chỉ định tuyến (route) để phân biệt cho URL.
- Parameter: Là tham số hỗ trợ cho phần Server biết cách hành xử với mỗi PATH
- Fragment: Là tham sỗ bổ trợ cho client để Client (web-browser/JS) biết cách xử lý phần cần hiển thị.

## fragment
A ***fragment*** is an internal page reference, sometimes called a named ***anchor***. It usually appears ***at the end*** of a URL and begins with a ***hash (#)*** character followed by an identifier. It refers to a ***section within a web page***.

In HTML documents, the browser looks for an anchor tag with a name attribute matching the fragment.


There are a few things about the fragments, ***the most important*** may be that they ***aren't sent*** in HTTP request messages but you can find some more info about them on [this page](https://blog.httpwatch.com/2011/03/01/6-things-you-should-know-about-fragment-urls/).

## Some Note

***URL fragments are not sent to the server over HTTP.*** They're just a browser concept.
If for some reason you need it, you could probably string together some JavaScript to include the fragment in the request as a `GET` parameter or something.
- Có nghĩa là, nếu muốn phía server biết được fragments content thì Client (JS) phải làm cách nào đó (mã hóa,...) để gửi tới server.

Trong [link này](https://www.webdevsplanet.com/post/url-parameters-vs-fragments) có phần so sánh `Differences between URL parameters and Fragments`

<table class="table table-striped table-bordered" width="100%"><thead><tr><th>URL Parameters</th><th>URL Fragments</th></tr></thead><tbody><tr><td>A question mark (?) is used to mark the start of a query string.</td><td>A hash (#) is used to mark the beginning of a fragment.</td></tr><tr><td>URL parameters are used to pass information to the page.</td><td>URL fragments are used from pointing and scrolling to elements on the HTML document.</td></tr><tr><td>A URL can have multiple parameters.</td><td>A URL cannot have more than one fragment.</td></tr><tr><td>URL parameters are passed in key-value pairs.</td><td>URL fragments comprise just a string of text after the hash (#).</td></tr><tr><td>Changing the URL parameters or their values and hitting enter reloads the page.</td><td>Changing the URL fragment identifier doesn’t reload the page. It just makes the browser scroll the page to the new location and records a new entry in the browser history.</td></tr><tr><td>The URL parameters are sent in the HTTP request and can be collected in the server-side code.</td><td>The fragment identifier is only used by the browser and not sent in the HTTP requests to the server. You cannot collect or see them in your server-side code.</td></tr></tbody></table>
