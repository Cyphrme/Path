package path

import (
	"fmt"
)

func ExamplePopulated_file() {
	paths := []string{
		"",
		"/",
		"app",
		"/app",
		"app.js",
		"app.min.js",
		"e/app.min.js",
		"/a/e/app.min.js",
	}

	for _, v := range paths {
		p := Populated(v)
		PrintPretty(p)
	}

	// Output:
	// {
	//     "full": ""
	// }
	// {
	//     "full": "/"
	// }
	// {
	//     "full": "app",
	//     "file": "app",
	//     "file_base": "app"
	// }
	// {
	//     "full": "/app",
	//     "dir": "/",
	//     "file": "app",
	//     "file_base": "app"
	// }
	// {
	//     "full": "app.js",
	//     "file": "app.js",
	//     "file_base": "app",
	//     "ext": ".js",
	//     "ext_base": ".js"
	// }
	// {
	//     "full": "app.min.js",
	//     "file": "app.min.js",
	//     "file_base": "app",
	//     "ext": ".min.js",
	//     "ext_base": ".js"
	// }
	// {
	//     "full": "e/app.min.js",
	//     "dir": "e/",
	//     "file": "app.min.js",
	//     "file_base": "app",
	//     "ext": ".min.js",
	//     "ext_base": ".js"
	// }
	// {
	//     "full": "/a/e/app.min.js",
	//     "dir": "/a/e/",
	//     "file": "app.min.js",
	//     "file_base": "app",
	//     "ext": ".min.js",
	//     "ext_base": ".js"
	// }
}

func ExamplePopulated_uri() {
	paths := []string{
		"",
		"https://cyphr.me/coze",
		"https://cyphr.me/assets/img/cyphrme_long.png",
		"https://localhost:8081/",
		"https://localhost:8081",
		"sftp://example.com/joe/bob/file.txt",
	}

	for _, v := range paths {
		p := Populated(v)
		//fmt.Println(p)
		PrintPretty(p)
	}

	// Output:
	// {
	//     "full": ""
	// }
	// {
	//     "full": "https://cyphr.me/coze",
	//     "dir": "https://cyphr.me/",
	//     "file": "coze",
	//     "file_base": "coze",
	//     "scheme": "https",
	//     "authority": "cyphr.me",
	//     "host": "cyphr.me",
	//     "uri_path": "/coze"
	// }
	// {
	//     "full": "https://cyphr.me/assets/img/cyphrme_long.png",
	//     "dir": "https://cyphr.me/assets/img/",
	//     "file": "cyphrme_long.png",
	//     "file_base": "cyphrme_long",
	//     "ext": ".png",
	//     "ext_base": ".png",
	//     "scheme": "https",
	//     "authority": "cyphr.me",
	//     "host": "cyphr.me",
	//     "uri_path": "/assets/img/cyphrme_long.png"
	// }
	// {
	//     "full": "https://localhost:8081/",
	//     "dir": "https://",
	//     "file": "localhost:8081",
	//     "file_base": "localhost:8081",
	//     "scheme": "https",
	//     "authority": "localhost:8081",
	//     "host": "localhost",
	//     "port": ":8081"
	// }
	// {
	//     "full": "https://localhost:8081",
	//     "dir": "https://",
	//     "file": "localhost:8081",
	//     "file_base": "localhost:8081",
	//     "scheme": "https",
	//     "authority": "localhost:8081",
	//     "host": "localhost",
	//     "port": ":8081"
	// }
	// {
	//     "full": "sftp://example.com/joe/bob/file.txt",
	//     "dir": "sftp://example.com/joe/bob/",
	//     "file": "file.txt",
	//     "file_base": "file",
	//     "ext": ".txt",
	//     "ext_base": ".txt",
	//     "scheme": "sftp",
	//     "authority": "example.com",
	//     "host": "example.com",
	//     "uri_path": "/joe/bob/file.txt"
	// }

}

func ExamplePathCut() {
	paths := []string{
		"app",
		"..subdir/test_5~fv=wwjNHrIw.js",
		"sftp://example.com/joe/bob/file.txt",
		"https://cyphr.me/coze",
		"https://cyphr.me/assets/img/cyphrme_long.png",
		"https://localhost:8081/",
		"https://localhost:8081",
	}

	for _, v := range paths {
		d, f := PathCut(v)
		fmt.Printf("dir: %s file: %s\n", d, f)
	}

	// Output:
	// dir:  file: app
	// dir: ..subdir/ file: test_5~fv=wwjNHrIw.js
	// dir: sftp://example.com/joe/bob/ file: file.txt
	// dir: https://cyphr.me/ file: coze
	// dir: https://cyphr.me/assets/img/ file: cyphrme_long.png
	// dir: https:// file: localhost:8081
	// dir: https:// file: localhost:8081
}
