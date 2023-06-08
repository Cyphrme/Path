# Path
-------------------------------------

Path is a library that enumerates path parts, including files and URI's.  


# URI's
```
                        Full path                          
                            |                              
/‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾\
https://cyphr.me:8081/bob/joe.txt?name=ferret#nose?name=bob
\___/   \___________/\__________/ \_________/ \___________/
  |            |           |            |           |      
scheme     authority    URI path       query      fragment  
        \______/\___/                         \__/\_______/
           |       |                            |      |   
          host     port                      anchor  fquery 
\___________________/\____________________________________/
          |                            |                   
        origin                      post origin            
\_______________________________/\________________________/
                 |                           |             
              URI base                      quag            
```

Naming for URI Paths
| Name                 | Example                     |
| -------------------- | --------------------------- |
| Full path            | https://cyphr.me:8081/bob/joe.txt?name=ferret#nose?name=bob |
| Origin               | https://cyphr.me:8081                                       |
| PostOrigin           | /bob/joe.txt?name=ferret#nose?name=bob                      |
| URIBase              | https://cyphr.me:8081/bob/joe.txt                           |
| Quag                 | ?name=ferret#nose?name=bob                                  |
| Scheme               | https                       |
| Authority            | cyphr.me:8081               |
| Host                 | cyphr.me                    |
| Port                 | :8081                       |
| URIPath              | /bob/joe.txt                |
| Query                | name=ferret                 |
| Fragment             | nose?name=bob               |
| Anchor               | nose                        |
| FragmentQuery        | ?name=bob                   |

Additionally, the file path information will be populated. 
| Name                 | Example                              |
| -------------------- | ------------------------------------ |
| Directory (dir)      | https://cyphr.me:8081/bob/           |
| File (filename)      | joe.txt                              |
| FileBase             | joe                                  |
| Extension (ext)      | .txt                                 |
| Extension base       | .txt                                 |


# File Path

```
        Full path                          
            |                              
/‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾\
/aaaaa/eeeee/app.min.js
\___________/\________/
       |         |     
  directory     file   
             \__/\___/
              |     | 
           base    ext
                    \__/
                      | 
									ext_base
```

File path parts
| Name                 | Example                              |
| -------------------- | ------------------------------------ |
| Full path            | /aaaaa/eeeee/app.min.js              |
| Directory (dir)      | /aaaaa/eeeee/                        |
| File (filename)      | app.min.js                           |
| FileBase             | app                                  |
| Extension (ext)      | .min.js                              |
| Ext Base (ext_base)  | .js                

If the starting slash is not included in input, it is not included in
output.  This also applies to relative references.  

# Query Parameters, Fragment Anchors, and Fragment Query Parameters
URIPath supports normal URL **query** parameters.  It also supports **fragment
query** parameters.  Unlike query parameters, fragments are not sent to the
server from the browser, which makes fragments **ideal for sensitive
information**.  We recommend applications use fragment query parameters over
query parameters when possible.   

    https://example.com:8042/over/there?name=ferret#nose?name=bob
    \_____________________________________________/\____________/
               |                                          |
       Sent to the server                            Kept local

For example, the browser requests only `https://example.com` when supplied the
URL `https://example.com#test` and does not include `#test` in the request.
Example.com is unaware of `#test`.

Fragment query parameters are located after the first `#`, then after the next
`?`.  It is ended by the end of the URL, by the next `?`, or other fragment
scheme like the delimiter `:~:`.  

If query parameter and fragment query parameter are set to the same key name,
the fragment query parameter takes precedence.  

Parts of the URL:

    foo://example.com:8042/over/there?name=ferret#nose?name=bob
    \_/   \______________/\_________/ \_________/ \___________/
     |           |            |            |            |
    scheme    authority      path         query      fragment


Where `nose?name=bob` is the **fragment**, `nose` is the **fragment anchor**,
and `?name=bob` is the **fragment query**. 

See [RFC 3986 for query
parameters](https://www.rfc-editor.org/rfc/rfc3986#section-3.5) and [Wikipedia](https://en.wikipedia.org/wiki/URI_fragment).

Fragment queries also have the advantage of not having size limits like normal
queries, although this is browser dependent.  Fragment query parameters are
"non-standard", but we hope it is useful enough to eventually standardize
through an RFC or other means.  


## Quag

We found it useful to name a super set of query and fragment, dubbed `quag`.
The quag includes `?` and `#`.

    foo://example.com:8042/over/there?name=ferret#nose?name=bob
    \_/   \______________/\_________/ \_________/ \___________/
     |           |            |            |            |
    scheme    authority      path         query       fragment
                                     \_________________________/
                                                  | 
                                                 quag


# See also:
- [FileVer][FileVer]
- [URLFormJS](https://github.com/Cyphrme/URLFormJS)



----------------------------------------------------------------------
# Attribution, Trademark notice, and License
Path is released under The 3-Clause BSD License. 

"Cyphr.me" is a trademark of Cypherpunk, LLC. The Cyphr.me logo is all rights
reserved Cypherpunk, LLC and may not be used without permission.

[FileVer]:https://github.com/Cyphrme/FileVer