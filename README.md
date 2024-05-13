<!DOCTYPE html>
<html lang="en" data-theme="auto"><head>
<meta http-equiv="content-type" content="text/html; charset=UTF-8">

<link rel="preconnect" href="https://www.googletagmanager.com/">
<script async="" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/gtm.js"></script><script async="" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/gtm.js"></script><script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00add8">
<link rel="stylesheet" href="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/css.css">
<link rel="stylesheet" href="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/styles.css">
<link rel="icon" href="https://go.dev/images/favicon-gopher.png" sizes="any">
<link rel="apple-touch-icon" href="https://go.dev/images/favicon-gopher-plain.png">
<link rel="icon" href="https://go.dev/images/favicon-gopher.svg" type="image/svg+xml">
<link rel="me" href="https://hachyderm.io/@golang">

  
  <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/site.js"></script>
<meta name="og:url" content="https://go.dev/doc/tutorial/web-service-gin">
<meta name="og:title" content="Tutorial: Developing a RESTful API with Go and Gin - The Go Programming Language">
<title>Tutorial: Developing a RESTful API with Go and Gin - The Go Programming Language</title>

<meta name="og:image" content="https://go.dev/doc/gopher/gopher5logo.jpg">
<meta name="twitter:image" content="https://go.dev/doc/gopher/gopherbelly300.jpg">
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@golang">
</head>
<body class="Site">
  
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-W8MVQXG"
  height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
  


<header class="Site-header js-siteHeader">
  <div class="Header Header--dark">
    <nav class="Header-nav">
      <a href="https://go.dev/">
        <img class="js-headerLogo Header-logo" src="https://go.dev/images/go-logo-white.svg" alt="Go">
      </a>
</header>
<br>
<h1>Tutorial Inicial de desenvolvimento de API REST em GO Utilizando o framework GIN</h1>
<h2>Esse tutorial foi pego do site ofical do go: [Link](https://go.dev/doc/tutorial/web-service-gin)</h2>
<h2>Tutorial: Developing a RESTful API with Go and Gin</h2>
<div id="nav" class="TOC"><table class="unruled"><tbody><tr><th colspan="2">Table of Contents</th></tr><tr><td class="first"><dl><dt><a href="#prerequisites">Prerequisites</a></dt><dt><a href="#design_endpoints">Design API endpoints</a></dt><dt><a href="#create_folder">Create a folder for your code</a></dt><dt><a href="#create_data">Create the data</a></dt><dt><a href="#all_items">Write a handler to return all items</a></dt><dt><a href="#add_item">Write a handler to add a new item</a></dt><dt><a href="#specific_item">Write a handler to return a specific item</a></dt><dt><a href="#conclusion">Conclusion</a></dt><dt><a href="#completed_code">Completed code</a></dt></dl></td><td><dl></dl></td></tr></tbody></table></div>

<p>This tutorial introduces the basics of writing a RESTful web service API with Go
and the <a href="https://gin-gonic.com/docs/" rel="noreferrer" target="_blank">Gin Web Framework</a> (Gin).</p>
<p>You’ll get the most out of this tutorial if you have a basic familiarity with Go
and its tooling. If this is your first exposure to Go, please see
<a href="https://go.dev/doc/tutorial/getting-started">Tutorial: Get started with Go</a>
for a quick introduction.</p>
<p>Gin simplifies many coding tasks associated with building web applications,
including web services. In this tutorial, you’ll use Gin to route requests,
retrieve request details, and marshal JSON for responses.</p>
<p>In this tutorial, you will build a RESTful API server with two endpoints. Your
example project will be a repository of data about vintage jazz records.</p>
<p>The tutorial includes the following sections:</p>
<ol>
<li>Design API endpoints.</li>
<li>Create a folder for your code.</li>
<li>Create the data.</li>
<li>Write a handler to return all items.</li>
<li>Write a handler to add a new item.</li>
<li>Write a handler to return a specific item.</li>
</ol>
<p><strong>Note:</strong> For other tutorials, see <a href="https://go.dev/doc/tutorial/index.html">Tutorials</a>.</p>
<h2 id="prerequisites">Prerequisites</h2>
<ul>
<li><strong>An installation of Go 1.16 or later.</strong> For installation instructions, see
<a href="https://go.dev/doc/install">Installing Go</a>.</li>
<li><strong>A tool to edit your code.</strong> Any text editor you have will work fine.</li>
<li><strong>A command terminal.</strong> Go works well using any terminal on Linux and Mac,
and on PowerShell or cmd in Windows.</li>
<li><strong>The curl tool.</strong> On Linux and Mac, this should already be installed. On
Windows, it’s included on Windows 10 Insider build 17063 and later. For earlier
Windows versions, you might need to install it. For more, see
<a href="https://docs.microsoft.com/en-us/virtualization/community/team-blog/2017/20171219-tar-and-curl-come-to-windows" rel="noreferrer" target="_blank">Tar and Curl Come to Windows</a>.</li>
</ul>
<h2 id="design_endpoints">Design API endpoints</h2>
<p>You’ll build an API that provides access to a store selling vintage recordings
on vinyl. So you’ll need to provide endpoints through which a client can get
and add albums for users.</p>
<p>When developing an API, you typically begin by designing the endpoints. Your
API’s users will have more success if the endpoints are easy to understand.</p>
<p>Here are the endpoints you’ll create in this tutorial.</p>
<p>/albums</p>
<ul>
<li><code>GET</code> – Get a list of all albums, returned as JSON.</li>
<li><code>POST</code> – Add a new album from request data sent as JSON.</li>
</ul>
<p>/albums/:id</p>
<ul>
<li><code>GET</code> – Get an album by its ID, returning the album data as JSON.</li>
</ul>
<p>Next, you’ll create a folder for your code.</p>
<h2 id="create_folder">Create a folder for your code</h2>
<p>To begin, create a project for the code you’ll write.</p>
<ol>
<li>
<p>Open a command prompt and change to your home directory.</p>
<p>On Linux or Mac:</p>
<pre><code>$ cd
</code></pre>
<p>On Windows:</p>
<pre><code>C:\&gt; cd %HOMEPATH%
</code></pre>
</li>
<li>
<p>Using the command prompt, create a directory for your code called
web-service-gin.</p>
<pre><code>$ mkdir web-service-gin
$ cd web-service-gin
</code></pre>
</li>
<li>
<p>Create a module in which you can manage dependencies.</p>
<p>Run the <code>go mod init</code> command, giving it the path of the module your code
will be in.</p>
<pre><code>$ go mod init example/web-service-gin
go: creating new go.mod: module example/web-service-gin
</code></pre>
<p>This command creates a go.mod file in which dependencies you add will be
listed for tracking. For more about naming a module with a module path, see
<a href="https://go.dev/doc/modules/managing-dependencies#naming_module">Managing dependencies</a>.</p>
</li>
</ol>
<p>Next, you’ll design data structures for handling data.</p>
<h2 id="create_data">Create the data</h2>
<p>To keep things simple for the tutorial, you’ll store data in memory. A more
typical API would interact with a database.</p>
<p>Note that storing data in memory means that the set of albums will be lost each
time you stop the server, then recreated when you start it.</p>
<h4 id="write-the-code">Write the code</h4>
<ol>
<li>
<p>Using your text editor, create a file called main.go in the web-service
directory. You’ll write your Go code in this file.</p>
</li>
<li>
<p>Into main.go, at the top of the file, paste the following package declaration.</p>
<pre><code>package main
</code></pre>
<p>A standalone program (as opposed to a library) is always in package <code>main</code>.</p>
</li>
<li>
<p>Beneath the package declaration, paste the following declaration of an
<code>album</code> struct. You’ll use this to store album data in memory.</p>
<p>Struct tags such as <code>json:"artist"</code> specify what a field’s name should be
when the struct’s contents are serialized into JSON. Without them, the JSON
would use the struct’s capitalized field names – a style not as common in
JSON.</p>
<pre><code>// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
</code></pre>
</li>
<li>
<p>Beneath the struct declaration you just added, paste the following slice of
<code>album</code> structs containing data you’ll use to start.</p>
<pre><code>// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
</code></pre>
</li>
</ol>
<p>Next, you’ll write code to implement your first endpoint.</p>
<h2 id="all_items">Write a handler to return all items</h2>
<p>When the client makes a request at <code>GET /albums</code>, you want to return all the
albums as JSON.</p>
<p>To do this, you’ll write the following:</p>
<ul>
<li>Logic to prepare a response</li>
<li>Code to map the request path to your logic</li>
</ul>
<p>Note that this is the reverse of how they’ll be executed at runtime, but you’re
adding dependencies first, then the code that depends on them.</p>
<h4 id="write-the-code-1">Write the code</h4>
<ol>
<li>
<p>Beneath the struct code you added in the preceding section, paste the
following code to get the album list.</p>
<p>This <code>getAlbums</code> function creates JSON from the slice of <code>album</code> structs,
writing the JSON into the response.</p>
<pre><code>// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
</code></pre>
<p>In this code, you:</p>
<ul>
<li>
<p>Write a <code>getAlbums</code> function that takes a
<a href="https://pkg.go.dev/github.com/gin-gonic/gin#Context" rel="noreferrer" target="_blank"><code>gin.Context</code></a>
parameter. Note that you could have given this function any name – neither
Gin nor Go require a particular function name format.</p>
<p><code>gin.Context</code> is the most important part of Gin. It carries request
details, validates and serializes JSON, and more. (Despite the similar
name, this is different from Go’s built-in
<a href="https://go.dev/pkg/context/"><code>context</code></a> package.)</p>
</li>
<li>
<p>Call <a href="https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON" rel="noreferrer" target="_blank"><code>Context.IndentedJSON</code></a>
to serialize the struct into JSON and add it to the response.</p>
<p>The function’s first argument is the HTTP status code you want to send to
the client. Here, you’re passing the <a href="https://pkg.go.dev/net/http#StatusOK" rel="noreferrer" target="_blank"><code>StatusOK</code></a>
constant from the <code>net/http</code> package to indicate <code>200 OK</code>.</p>
<p>Note that you can replace <code>Context.IndentedJSON</code> with a call to
<a href="https://pkg.go.dev/github.com/gin-gonic/gin#Context.JSON" rel="noreferrer" target="_blank"><code>Context.JSON</code></a>
to send more compact JSON. In practice, the indented form is much easier to
work with when debugging and the size difference is usually small.</p>
</li>
</ul>
</li>
<li>
<p>Near the top of main.go, just beneath the <code>albums</code> slice declaration, paste
the code below to assign the handler function to an endpoint path.</p>
<p>This sets up an association in which <code>getAlbums</code> handles requests to the
<code>/albums</code> endpoint path.</p>
<pre><code>func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}
</code></pre>
<p>In this code, you:</p>
<ul>
<li>
<p>Initialize a Gin router using
<a href="https://pkg.go.dev/github.com/gin-gonic/gin#Default" rel="noreferrer" target="_blank"><code>Default</code></a>.</p>
</li>
<li>
<p>Use the <a href="https://pkg.go.dev/github.com/gin-gonic/gin#RouterGroup.GET" rel="noreferrer" target="_blank"><code>GET</code></a>
function to associate the <code>GET</code> HTTP method and <code>/albums</code> path with a handler
function.</p>
<p>Note that you’re passing the <em>name</em> of the <code>getAlbums</code> function. This is
different from passing the <em>result</em> of the function, which you would do by
passing <code>getAlbums()</code> (note the parenthesis).</p>
</li>
<li>
<p>Use the <a href="https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run" rel="noreferrer" target="_blank"><code>Run</code></a>
function to attach the router to an <code>http.Server</code> and start the server.</p>
</li>
</ul>
</li>
<li>
<p>Near the top of main.go, just beneath the package declaration, import the
packages you’ll need to support the code you’ve just written.</p>
<p>The first lines of code should look like this:</p>
<pre><code>package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
</code></pre>
</li>
<li>
<p>Save main.go.</p>
</li>
</ol>
<h4 id="run-the-code">Run the code</h4>
<ol>
<li>
<p>Begin tracking the Gin module as a dependency.</p>
<p>At the command line, use <a href="https://go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them"><code>go get</code></a>
to add the github.com/gin-gonic/gin module as a dependency for your module.
Use a dot argument to mean “get dependencies for code in the current
directory.”</p>
<pre><code>$ go get .
go get: added github.com/gin-gonic/gin v1.7.2
</code></pre>
<p>Go resolved and downloaded this dependency to satisfy the <code>import</code>
declaration you added in the previous step.</p>
</li>
<li>
<p>From the command line in the directory containing main.go, run the code.
Use a dot argument to mean “run code in the current directory.”</p>
<pre><code>$ go run .
</code></pre>
<p>Once the code is running, you have a running HTTP server to which you can
send requests.</p>
</li>
<li>
<p>From a new command line window, use <code>curl</code> to make a request to your running
web service.</p>
<pre><code>$ curl http://localhost:8080/albums
</code></pre>
<p>The command should display the data you seeded the service with.</p>
<pre><code>[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
]
</code></pre>
</li>
</ol>
<p>You’ve started an API! In the next section, you’ll create another endpoint with
code to handle a <code>POST</code> request to add an item.</p>
<h2 id="add_item">Write a handler to add a new item</h2>
<p>When the client makes a <code>POST</code> request at <code>/albums</code>, you want to add the album
described in the request body to the existing albums’ data.</p>
<p>To do this, you’ll write the following:</p>
<ul>
<li>Logic to add the new album to the existing list.</li>
<li>A bit of code to route the <code>POST</code> request to your logic.</li>
</ul>
<h4 id="write-the-code-2">Write the code</h4>
<ol>
<li>
<p>Add code to add albums data to the list of albums.</p>
<p>Somewhere after the <code>import</code> statements, paste the following code. (The end
of the file is a good place for this code, but Go doesn’t enforce the order
in which you declare functions.)</p>
<pre><code>// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&amp;newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
</code></pre>
<p>In this code, you:</p>
<ul>
<li>Use <a href="https://pkg.go.dev/github.com/gin-gonic/gin#Context.BindJSON" rel="noreferrer" target="_blank"><code>Context.BindJSON</code></a>
to bind the request body to <code>newAlbum</code>.</li>
<li>Append the <code>album</code> struct initialized from the JSON to the <code>albums</code>
slice.</li>
<li>Add a <code>201</code> status code to the response, along with JSON representing
the album you added.</li>
</ul>
</li>
<li>
<p>Change your <code>main</code> function so that it includes the <code>router.POST</code> function,
as in the following.</p>
<pre><code>func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
</code></pre>
<p>In this code, you:</p>
<ul>
<li>
<p>Associate the <code>POST</code> method at the <code>/albums</code> path with the <code>postAlbums</code>
function.</p>
<p>With Gin, you can associate a handler with an HTTP method-and-path
combination. In this way, you can separately route requests sent to a
single path based on the method the client is using.</p>
</li>
</ul>
</li>
</ol>
<h4 id="run-the-code-1">Run the code</h4>
<ol>
<li>
<p>If the server is still running from the last section, stop it.</p>
</li>
<li>
<p>From the command line in the directory containing main.go, run the code.</p>
<pre><code>$ go run .
</code></pre>
</li>
<li>
<p>From a different command line window, use <code>curl</code> to make a request to your
running web service.</p>
<pre><code>$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
</code></pre>
<p>The command should display headers and JSON for the added album.</p>
<pre><code>HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 02 Jun 2021 00:34:12 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
</code></pre>
</li>
<li>
<p>As in the previous section, use <code>curl</code> to retrieve the full list of albums,
which you can use to confirm that the new album was added.</p>
<pre><code>$ curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
</code></pre>
<p>The command should display the album list.</p>
<pre><code>[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        },
        {
                "id": "4",
                "title": "The Modern Sound of Betty Carter",
                "artist": "Betty Carter",
                "price": 49.99
        }
]
</code></pre>
</li>
</ol>
<p>In the next section, you’ll add code to handle a <code>GET</code> for a specific item.</p>
<h2 id="specific_item">Write a handler to return a specific item</h2>
<p>When the client makes a request to <code>GET /albums/[id]</code>, you want to return the
album whose ID matches the <code>id</code> path parameter.</p>
<p>To do this, you will:</p>
<ul>
<li>Add logic to retrieve the requested album.</li>
<li>Map the path to the logic.</li>
</ul>
<h4 id="write-the-code-3">Write the code</h4>
<ol>
<li>
<p>Beneath the <code>postAlbums</code> function you added in the preceding section, paste
the following code to retrieve a specific album.</p>
<p>This <code>getAlbumByID</code> function will extract the ID in the request path, then
locate an album that matches.</p>
<pre><code>// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
</code></pre>
<p>In this code, you:</p>
<ul>
<li>
<p>Use <a href="https://pkg.go.dev/github.com/gin-gonic/gin#Context.Param" rel="noreferrer" target="_blank"><code>Context.Param</code></a>
to retrieve the <code>id</code> path parameter from the URL. When you map this
handler to a path, you’ll include a placeholder for the parameter in the
path.</p>
</li>
<li>
<p>Loop over the <code>album</code> structs in the slice, looking for one whose <code>ID</code>
field value matches the <code>id</code> parameter value. If it’s found, you serialize
that <code>album</code> struct to JSON and return it as a response with a <code>200 OK</code>
HTTP code.</p>
<p>As mentioned above, a real-world service would likely use a database
query to perform this lookup.</p>
</li>
<li>
<p>Return an HTTP <code>404</code> error with <a href="https://pkg.go.dev/net/http#StatusNotFound" rel="noreferrer" target="_blank"><code>http.StatusNotFound</code></a>
if the album isn’t found.</p>
</li>
</ul>
</li>
<li>
<p>Finally, change your <code>main</code> so that it includes a new call to <code>router.GET</code>,
where the path is now <code>/albums/:id</code>, as shown in the following example.</p>
<pre><code>func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
</code></pre>
<p>In this code, you:</p>
<ul>
<li>Associate the <code>/albums/:id</code> path with the <code>getAlbumByID</code> function. In
Gin, the colon preceding an item in the path signifies that the item is
a path parameter.</li>
</ul>
</li>
</ol>
<h4 id="run-the-code-2">Run the code</h4>
<ol>
<li>
<p>If the server is still running from the last section, stop it.</p>
</li>
<li>
<p>From the command line in the directory containing main.go, run the code to
start the server.</p>
<pre><code>$ go run .
</code></pre>
</li>
<li>
<p>From a different command line window, use <code>curl</code> to make a request to your
running web service.</p>
<pre><code>$ curl http://localhost:8080/albums/2
</code></pre>
<p>The command should display JSON for the album whose ID you used. If the
album wasn’t found, you’ll get JSON with an error message.</p>
<pre><code>{
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
}
</code></pre>
</li>
</ol>
<h2 id="conclusion">Conclusion</h2>
<p>Congratulations! You’ve just used Go and Gin to write a simple RESTful web
service.</p>
<p>Suggested next topics:</p>
<ul>
<li>If you’re new to Go, you’ll find useful best practices described in
<a href="https://go.dev/doc/effective_go">Effective Go</a> and
<a href="https://go.dev/doc/code">How to write Go code</a>.</li>
<li>The <a href="https://go.dev/tour/">Go Tour</a> is a great step-by-step
introduction to Go fundamentals.</li>
<li>For more about Gin, see the <a href="https://pkg.go.dev/github.com/gin-gonic/gin" rel="noreferrer" target="_blank">Gin Web Framework package documentation</a>
or the <a href="https://gin-gonic.com/docs/" rel="noreferrer" target="_blank">Gin Web Framework docs</a>.</li>
</ul>
<h2 id="completed_code">Completed code</h2>
<p>This section contains the code for the application you build with this tutorial.</p>
<pre><code>package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&amp;newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
</code></pre>





</article>



</main>
<footer class="Site-footer">
  <div class="Footer">
    <div class="Container">
      <div class="Footer-links">
          <div class="Footer-linkColumn">
            <a href="https://go.dev/solutions/" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Why Go
            </a>
              <a href="https://go.dev/solutions/use-cases" class="Footer-link" aria-describedby="footer-description">
                Use Cases
              </a>
              <a href="https://go.dev/solutions/case-studies" class="Footer-link" aria-describedby="footer-description">
                Case Studies
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://go.dev/learn/" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Get Started
            </a>
              <a href="https://go.dev/play" class="Footer-link" aria-describedby="footer-description">
                Playground
              </a>
              <a href="https://go.dev/tour/" class="Footer-link" aria-describedby="footer-description">
                Tour
              </a>
              <a href="https://stackoverflow.com/questions/tagged/go?tab=Newest" class="Footer-link" aria-describedby="footer-description">
                Stack Overflow
              </a>
              <a href="https://go.dev/help/" class="Footer-link" aria-describedby="footer-description">
                Help
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://pkg.go.dev/" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Packages
            </a>
              <a href="https://go.dev/pkg/" class="Footer-link" aria-describedby="footer-description">
                Standard Library
              </a>
              <a href="https://pkg.go.dev/about" class="Footer-link" aria-describedby="footer-description">
                About Go Packages
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://go.dev/project" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              About
            </a>
              <a href="https://go.dev/dl/" class="Footer-link" aria-describedby="footer-description">
                Download
              </a>
              <a href="https://go.dev/blog/" class="Footer-link" aria-describedby="footer-description">
                Blog
              </a>
              <a href="https://github.com/golang/go/issues" class="Footer-link" aria-describedby="footer-description">
                Issue Tracker
              </a>
              <a href="https://go.dev/doc/devel/release" class="Footer-link" aria-describedby="footer-description">
                Release Notes
              </a>
              <a href="https://go.dev/brand" class="Footer-link" aria-describedby="footer-description">
                Brand Guidelines
              </a>
              <a href="https://go.dev/conduct" class="Footer-link" aria-describedby="footer-description">
                Code of Conduct
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://www.twitter.com/golang" class="Footer-link Footer-link--primary" aria-describedby="footer-description">
              Connect
            </a>
              <a href="https://www.twitter.com/golang" class="Footer-link" aria-describedby="footer-description">
                Twitter
              </a>
              <a href="https://github.com/golang" class="Footer-link" aria-describedby="footer-description">
                GitHub
              </a>
              <a href="https://invite.slack.golangbridge.org/" class="Footer-link" aria-describedby="footer-description">
                Slack
              </a>
              <a href="https://reddit.com/r/golang" class="Footer-link" aria-describedby="footer-description">
                r/golang
              </a>
              <a href="https://www.meetup.com/pro/go" class="Footer-link" aria-describedby="footer-description">
                Meetup
              </a>
              <a href="https://golangweekly.com/" class="Footer-link" aria-describedby="footer-description">
                Golang Weekly
              </a>
          </div>
      </div>
    </div>
  </div>
  <div class="screen-reader-only" id="footer-description" hidden="">
          Opens in new window.
  </div>
  <div class="Footer">
    <div class="Container Container--fullBleed">
      <div class="Footer-bottom">
        <img class="Footer-gopher" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/pilot-bust.svg" alt="The Go Gopher">
        <ul class="Footer-listRow">
          <li class="Footer-listItem">
            <a href="https://go.dev/copyright" aria-describedby="footer-description">Copyright</a>
          </li>
          <li class="Footer-listItem">
            <a href="https://go.dev/tos" aria-describedby="footer-description">Terms of Service</a>
          </li>
          <li class="Footer-listItem">
            <a href="http://www.google.com/intl/en/policies/privacy/" aria-describedby="footer-description" target="_blank" rel="noopener">
              Privacy Policy
            </a>
            </li>
          <li class="Footer-listItem">
            <a href="https://go.dev/s/website-issue" aria-describedby="footer-description" target="_blank" rel="noopener">
              Report an Issue
            </a>
          </li>
          <li class="Footer-listItem go-Footer-listItem">
            <button class="go-Button go-Button--text go-Footer-toggleTheme js-toggleTheme" aria-label="Toggle theme">
              <img data-value="auto" class="go-Icon go-Icon--inverted" height="24" width="24" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/brightness_6_gm_grey_24dp.svg" alt="System theme">
              <img data-value="dark" class="go-Icon go-Icon--inverted" height="24" width="24" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/brightness_2_gm_grey_24dp.svg" alt="Dark theme">
              <img data-value="light" class="go-Icon go-Icon--inverted" height="24" width="24" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/light_mode_gm_grey_24dp.svg" alt="Light theme">
            </button>
          </li>
        </ul>
        <a class="Footer-googleLogo" target="_blank" href="https://google.com/" rel="noopener">
          <img class="Footer-googleLogoImg" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/google-white.png" alt="Google logo">
        </a>
      </div>
    </div>
  </div>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/jquery.js"></script>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/carousels.js"></script>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/searchBox.js"></script>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/misc.js"></script>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/hats.js"></script>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/playground.js"></script>
  <script src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/godocs.js"></script>
  <script async="" src="Tutorial%20Developing%20a%20RESTful%20API%20with%20Go%20and%20Gin%20-%20The%20Go%20Programming%20Language_arquivos/copypaste.js"></script>
</footer>
<section class="Cookie-notice js-cookieNotice">
  <div>go.dev uses cookies from Google to deliver and enhance the quality of its services and to
  analyze traffic. <a target="_blank" href="https://policies.google.com/technologies/cookies">Learn more.</a></div>
  <div><button class="go-Button">Okay</button></div>
</section>




















</body></html>