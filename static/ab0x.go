// Code generated by fileb0x at "2018-08-15 14:23:28.543068851 +0300 MSK m=+0.026268369" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(7f4f9146f89e41154e74cc05ac788824.d47d734c826bc110696e7160789432d0)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileSwaggerJSON is "swagger.json"
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5d\xeb\x6f\xdc\x38\x92\xff\x9e\xbf\x82\xe8\x3b\x20\x33\xd8\xb4\x3d\x9b\x5d\x2c\x70\xf3\xcd\x93\xce\x24\xc6\xc6\x73\x0d\xdb\x59\x2c\x70\x0e\x0c\x5a\x62\x77\x73\x23\x91\x0a\x45\x75\xec\x09\xfc\xbf\x1f\x44\xea\x41\xea\xc9\xd6\xcb\x6a\x9b\xdf\x9c\x0e\xc5\x47\x91\xf5\xab\x07\xab\x8a\x3f\x5e\x01\xb0\x70\x28\x09\x23\x1f\x85\x8b\x5f\xc1\xff\xbd\x02\x00\x80\x05\x0c\x02\x0f\x3b\x90\x63\x4a\x4e\xff\x13\x52\xb2\x78\x05\xc0\x97\x37\x71\xdb\x80\x51\x37\x72\xcc\xda\x86\xdf\xe1\x76\x8b\xd8\xe2\x57\xb0\x78\x7b\xf2\xcb\x42\xfc\x86\xc9\x86\x2e\x7e\x05\x3f\xe4\xb7\x2e\x0a\x1d\x86\x83\xf8\xdb\xb8\xd5\x25\x0a\x69\xc4\x1c\xb4\xbc\x42\x6c\x8f\x1d\x04\x70\x08\xc2\xe4\xcf\x0d\x65\x20\xe4\x94\x61\xb2\x05\x5f\xa3\x3b\xc4\x08\xe2\x28\x04\x2c\xf9\x24\x04\x98\x80\x0b\x4a\xb6\x14\xac\x7e\x3b\x11\x63\x01\xb0\xe0\x98\x7b\xa8\xaa\xe7\xb4\xc1\x1e\xb1\x30\x19\xfc\x97\x93\x5f\x4e\xfe\x1a\x4f\xfe\x51\x2e\x14\xf2\x5d\x98\xcf\xf4\xd4\xa1\x64\x83\xb7\x3e\x0c\xf2\x1f\x01\x58\x6c\x11\x57\xfe\x19\x8f\x08\xb7\x39\x6d\x92\xdf\xde\x89\x4f\x2f\x60\xb0\xc8\x7e\xfd\xf2\x26\xff\x24\x8c\x7c\x1f\xb2\x87\x78\x0e\x1f\x10\x07\x51\x88\x18\xc8\x47\x03\x1e\x0e\x79\xba\x20\xd1\x9e\x06\x88\x09\x72\x9f\xbb\xc9\x37\x57\xc8\x43\x0e\x47\x6e\x36\x50\xf8\x09\x87\x5c\xfd\x26\x80\x0c\xfa\x88\x23\x56\x9c\xdc\x0f\xe5\xef\x78\xfa\x0f\x81\xa0\x57\xc8\x63\x42\x2b\x3d\x88\xff\xdd\x50\xe6\xc3\x78\xc1\x8b\x28\xc2\x6e\xf1\x7f\x09\xf4\xc5\xb7\xff\x5e\x7e\x0e\x11\x5b\x9e\xaf\x8a\x0d\xb0\xa0\xf3\x0e\x41\x17\xb1\xe2\xff\x31\xf4\x2d\xc2\x0c\xc5\x2b\xe2\x2c\x42\xca\x7f\x3e\xbe\xa9\x9f\x2e\x22\x91\x5f\x58\x90\xf8\x3d\xa6\x61\x61\x84\xf8\xa4\xba\x3e\x26\x0b\xed\xd7\x2f\x6f\x0e\x59\x7f\x61\x85\x97\xd4\x43\xe3\xaf\xf1\xa0\x29\xfd\x01\x7d\x14\x06\xd0\x19\x6a\x5e\x95\x07\x96\xa1\x30\xa0\x24\x44\xa1\x76\xf4\x01\x58\xbc\xfd\xe5\x97\xc2\x4f\x65\x1e\x2f\x9c\xec\xe2\x5c\x42\x67\x87\x7c\x58\xea\x05\x80\xc5\x7f\x33\xb4\x89\x3b\xf8\xaf\x53\x17\x6d\x30\xc1\x71\x87\xe1\x69\x7e\xe4\x2f\x93\x59\xe9\x1b\xfc\x58\x47\xe5\x85\x8b\x36\x30\xf2\x78\xfb\x84\x77\x88\xb1\x07\x80\x18\xa3\xac\xef\x6c\xb7\x75\xbc\xba\x92\x93\xf9\x8d\xba\x0f\x0d\xf3\x7f\x55\xb1\x92\xc5\xfd\xd2\x47\x7c\x47\xdd\xe5\x1e\x87\xf8\x0e\x7b\x98\x0b\x28\x09\xa2\x3b\x0f\x3b\x69\x67\xf2\xd3\xe4\xb3\xc5\xa9\x4b\x7d\x88\xc9\xe1\x48\xb6\x12\xdf\x19\xc1\x58\x32\x84\x09\x7c\xc9\x5e\x3b\x63\x96\x05\x81\x74\x4a\x01\xdc\xd6\xcc\xe5\x5b\x84\x98\x76\xb2\x06\x18\x0c\xb1\xdb\x43\x06\x1c\x1e\x4a\xd4\x23\xd6\x97\x33\x93\x33\x68\x41\xa4\x09\x44\xd2\x0f\x16\x01\x0d\x87\x43\x8b\x33\xd7\x4d\xd0\xa2\x09\x27\xce\x5c\x37\xe9\xcf\x42\x44\xe5\x1a\xd3\x41\xef\xe2\xed\xaf\x1c\xad\xea\x7f\xba\xf1\x89\xc9\xf1\x32\x64\xf3\xbf\x9a\xb2\xf9\xe8\x33\xb7\xda\x81\xae\x1d\x9c\xfe\x90\x7f\x3c\x8e\xaf\x26\x18\x29\x08\x96\xf1\xfb\x88\x6b\xb7\x48\x43\x65\x36\xb1\xad\xfb\xa4\x46\x81\x65\xf1\xe9\x64\xb7\x8b\x3c\xc4\xd1\xc4\xd2\x7b\x25\x06\x95\x5d\x7e\x84\xc4\xf5\xb4\xc3\x6f\xd9\xf9\x98\xd8\xf9\xad\x21\x3b\x03\x79\xd2\xdc\x85\xe5\xc0\x0a\x21\x8b\xfd\x80\x32\x5e\xe3\x56\x34\xd3\xaf\xcd\xfd\x8a\xe7\x62\x30\xc5\xab\xd8\xc4\xab\xb2\x71\x4e\x1c\xeb\x47\xb4\x7e\xc4\xa3\xb7\x5b\x0a\xae\xf1\xe1\xec\x97\xb7\x87\x78\x3c\x25\xd3\x5b\x48\x6c\x84\x44\x17\x05\x1e\x7d\xf0\x11\xe1\x1d\x30\x71\x95\x7d\x6c\x0a\x8a\xca\x70\xed\xa8\xb8\x52\xe6\x66\x61\xd1\xba\x4d\x8c\x2d\x93\xfc\xdc\x4c\x8e\x3f\xca\x01\xb7\x00\x64\x02\x40\x98\x6c\x19\x0a\x75\x6a\x1b\xc2\xcf\xb9\xfc\xd4\x14\x7b\xb2\x91\xda\x91\xe7\x3c\x9b\x94\xc5\x1d\x8b\x3b\xa6\x9c\x95\x9d\x9a\xc9\x51\x27\x3b\xda\x16\x73\x4c\x30\x27\x09\x73\xe9\x00\x39\x69\x44\x8b\x21\xe4\xa4\x03\xb5\x23\xce\x55\x3a\x25\x0b\x38\x16\x70\x4c\xd9\x2a\x3d\x34\x93\xe3\x4d\x7a\xae\x2d\xdc\x34\xc3\x4d\x95\x6e\x63\x74\xa9\x63\xaa\xda\x64\x31\x6c\xb9\x04\x38\x20\x84\x4d\x17\x58\x16\x79\xac\xe7\x69\x76\x11\x6c\xfa\xb9\x1e\x4c\x43\xb3\xa1\x27\xe6\x28\x46\xd2\x43\xa2\xc1\x98\xe9\xcd\x56\x1a\x0c\xdc\x8e\x65\xf2\xda\x0a\x40\xcf\x93\x90\x96\x45\x1b\xb7\xdf\x75\x9d\x79\x5e\x3e\x8e\xc5\xb1\xa3\xc1\xb1\xe1\x35\x93\xf2\xe9\xb1\x37\x63\x0c\xef\x21\x47\xad\xdc\x7d\xfa\x23\xfb\xfb\x71\x32\x4e\xd7\x52\x0a\xb2\xf1\x0f\xe3\xf8\x73\x52\x25\xc7\x06\x62\xfe\xf4\xf0\x93\x66\x51\xf9\xe4\xd7\xc0\xf5\xc4\xb4\x0c\xd0\x81\x01\xa6\x4b\x41\x39\x2c\xfb\xc4\x66\x9d\x58\x9d\x7d\xb2\xdb\xe2\x63\x40\x46\x9b\x04\xf3\x3c\xe2\xd7\xcd\x41\xf3\x1d\x43\x90\xa3\x1c\x37\x9b\x10\x53\xb6\xcd\xfb\xb6\x78\x69\xf1\xf2\x79\xe1\xe5\x7c\x82\x7d\x26\xcd\x53\xc8\xb8\x1f\x38\x82\xc5\xdd\xbe\xcb\x48\xad\x09\x93\xe5\x58\x54\x37\xb7\x0a\xcd\x71\x3d\xb1\x0a\x8d\x70\x3d\x33\x02\xc7\x8c\x9c\x3c\x26\xd3\x4f\x8d\x7b\xb3\xb6\xdf\x60\xb6\xdf\xe9\x8f\xec\xef\xc7\x79\xda\x81\x56\xa5\xb1\x2a\xcd\xcb\x51\x69\xcc\xa6\x93\xb1\xd2\xdc\x2d\x52\xab\xb4\xbc\x3c\xa5\xc5\x22\xb7\x45\x6e\x8b\xdc\xb3\x44\xee\x03\xd2\x4b\xac\x62\x6d\x16\x32\xa0\xe9\xd5\x35\xc9\x26\x66\x69\xee\xe6\xb9\x26\x22\xd5\x5d\x89\xc3\x37\xa9\x8a\x53\x48\x19\xb0\xd0\x6c\xa1\xd9\xde\xab\x98\xa7\xba\x0c\x52\x15\x28\xef\xcf\xde\xac\x0c\x59\x19\xc8\x1c\x39\x93\xab\x95\x7c\x67\xdb\xef\x56\x94\xde\x2d\x68\x5a\xd0\xb4\x97\x2b\x23\xa5\x32\x4e\x5b\x05\x2a\x1b\x76\xe8\xeb\x15\xb9\x20\x8b\xeb\xc3\x54\x8d\x31\x47\x76\x25\xe4\x4e\x4b\x51\x3d\x34\xe8\x6e\xd4\x9c\xec\x63\xba\x70\xa9\xa3\xa2\x35\x0c\xbb\xdc\xb8\x28\xd4\x3c\xfd\x91\xff\xe3\x71\x3a\x33\x11\x40\x87\xe3\x3d\x02\x49\x91\xea\x16\x7b\xf1\x4c\x34\xb6\xba\x8f\xd5\x7d\xac\x2f\xaf\xb6\x6a\x56\x15\x77\xcc\xd3\x80\xb5\xda\xcd\x94\x56\x6b\x34\x2c\x8e\x7f\x0e\x5c\x73\xa3\x55\x36\xb6\xc0\x6d\x81\xdb\x02\xf7\x44\xc0\xfd\x3c\x8d\xe8\x43\x4a\x01\x81\x48\xa0\x8e\x35\xa2\x9f\x87\x11\x6d\x26\x69\x92\x2a\xac\x56\xd2\x58\x49\x63\x25\xcd\x3c\x4d\x84\x83\x30\xdc\x5e\xf8\xf7\xba\xf0\xd7\xdc\x3a\xa7\xd8\x87\x5b\xa4\x95\x5f\x1a\xc7\x28\x10\xe3\x00\x51\x1e\x39\x77\xd7\x39\x94\x70\x88\x09\x62\x4d\xf0\x7d\x25\x62\x6c\x65\xbb\x73\x5f\x7f\x7a\xc5\xe2\xf7\x73\xc7\x6f\xab\x91\x8f\xaa\x91\x4b\xee\x94\x6c\x65\x55\x72\x2b\x2a\x9a\x44\x05\x43\xe2\xc9\xcb\x70\x02\x69\xa1\x0a\x89\x74\x58\xe0\xd0\xa8\x59\xd1\xbf\x42\xfc\x32\x9d\xa3\x95\x11\x56\xc7\xb7\xb2\xeb\x05\xc8\xae\x8c\xe3\xad\xf8\xb2\xe2\xab\x49\x7c\x25\xf7\xc9\xf3\x0d\x78\xfe\x57\x32\x41\x1b\xf7\x6c\xe5\x97\x95\x5f\x47\x7d\x8d\x9d\x06\xaf\xd8\x78\xec\x23\x14\x10\xa7\x3f\x92\xbf\xa6\x0c\x7a\x32\x8b\x76\x2a\x09\x0b\x2b\x27\xac\x9c\xb0\x72\xe2\xa9\xec\x1c\xb3\xf9\xec\x4b\xac\x6a\x63\xaf\x6c\xec\xd5\xc0\x8e\xb3\x4b\x14\x1f\xb7\x03\x25\x8a\xfc\xc8\x8a\x12\x2b\x4a\xac\x28\x79\x71\xa2\x64\x46\xd1\x60\x29\x02\x3d\x91\x0b\x2f\xa1\x2a\x60\x02\x0e\xad\x2b\xef\x39\x64\xce\x26\x49\x24\x87\xc9\xc3\x77\x3b\x48\xb6\xc8\xa6\x94\x58\xc1\x68\x05\xa3\xb5\xb1\x06\x4c\x53\xac\x03\x23\xe0\x08\xc4\xb1\x22\xe7\xb9\xc5\x23\x67\x1b\xfc\x13\xa1\x69\x46\xe3\xcf\x87\x04\x29\x5b\xa3\xcc\xca\x1e\x2b\x7b\xac\xec\xb1\x81\xd3\x47\x73\x5b\x74\xea\xe2\xcd\x66\xd4\xb7\xfb\xdf\x51\x3f\x80\xac\x4a\xce\x84\xe0\x3b\xe6\x3b\x10\x30\xb4\xc7\x34\x0a\x4d\x8c\x9d\x15\xde\x6c\xf2\xb1\xd7\xc9\x97\x69\x10\xc2\xb1\xd4\x94\xb0\x68\x32\xed\x6d\x41\x7c\xc4\x6d\x41\x8d\x11\x01\x24\xfb\xe7\xdb\xc7\x49\xb0\x84\x7f\xa7\x55\x78\x62\x8e\x1c\x16\x31\x5e\x86\x53\xf8\xa0\xc9\xbc\xb5\xf8\xf5\x4c\xf0\x6b\x9a\xa7\xb3\x0f\x7a\x35\x5b\x7b\x2d\xfb\x23\x24\xae\xa7\xd9\x43\xd6\x50\xb6\x86\xb2\x2d\x14\x0b\xec\x1b\xde\x47\x79\xd5\x65\x8a\x9a\xc9\x3d\x57\xb2\xa5\xed\xe5\x61\xd3\x7e\x2d\x4e\x5a\x9c\x7c\x31\x0e\xc5\x69\x02\x19\x4a\x2c\x0b\xc6\x2f\x0c\x9b\xf0\xfd\xd0\x55\x61\xdb\x97\x62\x91\xdc\xfc\x06\xc9\x14\xcb\x95\x9a\xb0\x99\x88\x36\x2a\x02\x9b\x89\xe0\x97\x5d\x02\x36\xd7\x6b\x6c\x01\xd8\xc1\xec\xbd\xd3\x1f\xc9\x9f\x8f\xb3\xb3\xfc\xac\x12\x63\x95\x18\x7b\x2b\xaa\x4f\x07\x97\x58\x63\x96\x96\xa7\x55\x54\x66\x97\x65\x60\x8a\xd6\x69\x25\xa7\x76\x8b\x33\xa9\x2a\x63\xc1\xda\x82\xb5\x05\xeb\xb1\xc1\xfa\x19\x9a\xbf\x6f\x8d\xcd\xdf\x81\xab\x6f\x58\xa9\xf2\x84\xe6\xaf\x81\x60\x91\x2d\xad\x60\xb1\x82\xc5\x0a\x96\xd9\x59\x01\xe6\xa8\x6d\x03\x11\x0f\x0f\x44\x0c\x11\xdb\x63\xa7\xc3\x35\xfc\x95\xfc\xd0\xc8\x19\x93\x0e\x62\xe2\x8b\x49\xfa\xb5\xe5\x8a\x2c\x14\xdb\xdb\x77\x13\x1f\x88\xc6\x5c\x7d\xd1\x2a\xe5\x3e\x7b\xf7\x3e\xe0\xdd\xbb\x29\x54\x26\x77\xef\xc9\x86\xb6\xdf\xbd\xa7\xfd\x5a\x94\xb4\x28\x69\xef\xde\x07\x75\x3e\x94\x58\x76\x82\xbb\xf7\x84\xef\x87\xbe\x7b\x6f\x5f\x8a\x45\x72\x73\xe7\x83\x29\x96\x2b\x77\xef\xa9\x80\x36\xba\x7a\x4f\x05\xf0\xcb\xbe\x79\xcf\x74\x1a\x7b\xf1\xde\xf7\xe2\x3d\x25\xe5\xe9\x8f\xe4\xaf\xc7\xb9\x99\x7a\x56\x7f\xb1\xfa\x8b\x75\xb8\xe9\xd3\x09\x4b\xac\x31\x43\x93\xd3\x6a\x28\xb3\xbb\x74\x37\x85\xea\xe4\xd2\xdd\xc0\xd4\x94\x2d\x2d\x54\x5b\xa8\xb6\x50\x3d\x3a\x54\x3f\x43\xbb\xf7\xad\xb1\xdd\x3b\xf0\xa5\xbb\x95\x2a\x4f\x68\xf7\x1a\x08\x16\xd9\xd2\x0a\x16\x2b\x58\xac\x60\x99\x9d\x0d\x60\x8e\xda\xf6\xd2\xbd\xc3\xa5\x3b\xf5\x22\x39\xd1\x1f\xe9\x9f\x5a\x45\x0f\xd5\x43\x33\xa6\x63\x32\x19\x5b\x7d\x5b\xc8\xcc\x49\x99\x7c\xa8\xbc\x20\xf2\x62\xca\x74\xa4\x44\x9b\xb3\xf7\xb4\x62\x5f\x2d\x9f\x76\x72\x9e\x56\x31\x6a\x55\xc8\xcc\x14\x5c\x7a\xd8\x3d\x42\xf2\xd5\xb1\xdd\x27\xbc\x28\xfe\xcc\x9c\xe5\x96\x39\x9b\x98\x93\x25\x16\xdd\xe1\x41\x6a\x97\xd9\x97\x26\x77\x17\xd9\x38\xed\x6f\xc2\x7e\x40\x3c\xeb\xfb\x5d\xdc\xd8\x1a\x30\xc7\x63\xc0\x0c\xef\x10\x2f\x9c\x9c\xbe\x9c\x57\x3a\x5d\x36\x1c\xab\x5d\xe7\x7e\x95\x7c\xba\x50\x26\x95\xcd\x7e\x71\xe6\x38\x28\x0c\x3f\xa1\x3d\xf2\x54\x14\xa9\x39\x82\x8b\xfb\xe5\x96\x2e\x03\xe8\x7c\x85\x5b\xf1\xff\x5b\xcc\x4f\x1c\x4a\x38\xc4\x04\xb1\xc8\x3f\x21\x88\x9f\x3a\xbb\x0c\x99\x96\x09\x92\x9f\xee\x11\x71\x29\x3b\xdd\x62\xbe\x8b\xee\x4e\x1c\xea\x9f\x2a\x5f\x9d\x7e\x8d\xee\xd0\xd2\xf1\x30\x22\xfc\x34\xf8\xba\x3d\xf5\xa9\x8b\xbc\x85\x06\x76\x92\x74\xbf\x63\x0f\x55\x4d\x93\xde\xfd\x07\x39\xf9\xf1\x5a\x04\x2c\xc6\x25\x8e\x0b\xa7\x76\xe1\x88\x5e\x6e\x37\x7a\x37\x6d\x4c\x27\x57\x9d\xf2\x5d\x6c\x9e\x2f\x2a\x77\x25\xe9\xa2\x5b\xb7\xd7\x71\xc3\x57\xc5\x6d\x7f\x9c\x0b\xe5\x2f\x60\xa0\xeb\x75\x1a\xbf\x64\x6d\xc0\x72\x09\x44\x1f\x60\x43\x19\x90\xe4\x06\x3e\x0c\xf2\xad\xa9\xdb\x32\x05\x97\x72\x24\x95\x60\xa7\x90\xd8\x85\x1c\xa6\x54\xfa\xd2\xbe\xdd\x32\x62\xe8\x16\x16\x39\xbe\xc4\xef\x71\x43\x61\x19\xc8\x94\x4b\x70\xf9\xfb\xbb\xbf\xfd\xed\x6f\xff\x03\x12\x61\xf3\xa6\xd3\x96\xca\x68\x44\xf7\x8c\x57\x1f\x17\xb1\x96\xc2\xbc\xaa\x21\x24\x23\xef\x4a\x59\x7e\xa1\x33\xa9\x29\xb5\x2f\xd5\x4d\x6a\xa7\x0f\xb9\x50\xa9\x5d\xd7\x2e\x34\x69\x36\x34\xbb\xe5\x0a\x76\x9f\xbe\x65\x0f\x95\x03\xd0\xef\x04\xb1\xce\x9d\xff\xaf\xf8\x7a\xd6\x1c\xbd\xd2\x4f\x60\x2d\x57\xc7\xed\xea\x38\x1b\x88\x63\xdc\xca\xde\xd0\x75\xc5\x61\x86\xde\xba\x86\x59\x0b\x84\x9d\x21\xb9\x64\xee\x85\x01\xbd\x44\xc3\x3a\x82\x15\x02\xd2\x3b\xca\xb0\xb8\xa3\xba\xa3\x09\x19\x83\xfa\x8d\xd7\x02\x73\xe4\x87\x65\xb5\xa7\x05\x6e\xea\xed\x2e\x1d\xe7\xb2\x75\xcf\xfb\xb8\xe7\xd1\xfb\x26\x7b\x98\x36\x8e\xf7\x31\xa3\x39\x48\xd5\xe1\xbe\xfb\x77\x5b\xb5\x81\xd5\xdb\x11\x1f\xa6\x71\x49\x9c\x11\x2e\x3c\xcd\x56\x5a\xa2\xa1\xec\xa5\x99\x74\xb2\x4d\xe9\xe4\x27\x3f\x63\xd5\xf3\xd6\x55\x23\xc0\x7e\xbc\xb8\x37\xf5\x2a\x82\x87\x7d\xcc\xc3\x03\x94\x04\xea\xfb\x90\xb8\x03\x70\x53\x35\x82\xb5\x32\x4f\x32\x7c\xa5\x08\x6a\x38\x2c\x03\x72\xbb\xdc\x9f\x7f\x51\x2f\xf2\x51\x1f\x9e\x57\x26\x8e\xc8\x7e\xac\x09\xbf\x27\x7b\xc3\x49\x6a\x2d\xd5\xd9\xc9\x43\xd4\x55\xb2\x9f\x8b\xaf\x2b\x3b\x4e\x0e\x9f\x11\x5f\xa7\xa6\xec\xa4\x4a\x53\x40\x19\x1f\xff\x28\xad\x29\xe3\x86\x7b\xb4\x16\x13\xaa\x9c\xeb\x5e\x1c\xc8\x5b\x3f\xb6\xf4\x67\x76\xfc\x65\xe3\x0b\x39\xb3\x99\x0a\x3d\x65\x2b\x0c\x40\x3b\x6e\xa7\x03\x77\x7c\x54\x62\xcc\xce\x46\x1e\xca\x88\x8b\x3b\xd6\xfe\xcd\x28\xa7\x0e\xf5\xcc\x31\x7b\x44\xe6\xa8\xeb\x16\x13\x8e\xb6\x05\x77\x9a\xe2\x8b\xc4\x84\xff\xe3\xef\x8b\xe6\x63\x5e\x33\x68\xba\x7a\x23\xd0\x58\x17\x68\x35\xcf\x43\x97\xf0\x92\xc1\xb1\x93\x2d\xe3\x83\x27\xb9\x1d\xfc\xa4\x29\xcb\x3f\x03\xc1\xfd\xc8\x1d\xe4\x20\x8a\xbe\x6e\xc5\x5d\x87\xf1\x51\x8b\x97\xd9\xf9\xa8\x5d\xc4\x1f\x57\xee\xba\x32\x95\xee\x9d\x47\x84\xaf\x95\xc5\x4c\x22\x40\xc2\xe8\xae\xdf\xb4\xaf\xa2\xbb\xc2\xa4\x67\x74\x80\x85\xc7\xa6\xfd\xf4\x2a\xcd\xc0\x72\xd9\xdd\x16\xf0\xe0\x1d\xf2\x3a\x53\xf2\x93\xf8\xba\x7a\x97\x38\x65\x7d\x54\x9c\xab\xe4\xfb\x6a\x3f\x27\x64\x78\xb3\xb9\xc5\x6e\x77\x67\xa7\xe8\xe1\x7c\x35\xcb\x43\xa0\xbc\xea\x52\x7f\x04\xf2\x46\xba\xd0\x74\x2b\x42\x3f\x0e\x44\xa9\x6c\xba\x61\x93\x99\xc3\x50\xe0\x61\x07\x1e\x60\xe8\xc8\x67\x11\xeb\xf6\xec\x8e\x52\x0f\x41\xd2\xb0\x69\xf2\xa1\xde\x5a\x3b\x25\x9d\xf3\xd8\x7a\x9a\xb9\x81\x92\xce\xa8\x7a\xc6\xf3\xf6\x0d\xcf\xd8\x9d\x2b\x8c\xa7\xdb\x20\xf2\xbc\xdb\x10\x39\x0c\xf1\x27\xb2\x9c\x85\x19\xb6\x8e\x3c\xef\x4a\xcc\x22\xb4\xce\xe7\x7a\xe7\x73\x15\x70\x8c\xa3\xe8\x5e\x16\x70\xa9\x20\x97\x92\x10\x8f\x0a\xe1\x51\x38\xd8\x69\xa0\x0e\x38\x5f\x81\x9f\x28\xf1\x1e\x00\xde\xa8\x2f\x67\xe1\x10\x04\x90\x71\x40\x37\x59\xdc\xc8\xcf\x1d\xcf\x7b\x3a\x94\x26\x8f\x34\x61\x0a\x79\x64\x68\xd4\xe7\x72\xe1\x4a\x7e\x55\x2d\x41\x29\x87\xde\xad\x13\x44\x2d\x44\x10\xed\xc0\xbb\xf5\x67\x10\x85\x70\x8b\xc0\xdd\x83\x08\x7d\xca\xd1\x36\xe6\x79\xbe\xc3\x21\xa8\x79\x25\xcb\x74\x57\xa3\xb6\x6d\xbd\x8e\x67\xf2\x6e\xfd\xb9\x69\x3d\x3e\xf2\xa9\x88\x1a\x69\x5f\xd2\xe5\xd9\xc5\x3c\x96\x74\x21\xe7\x5c\xed\x7f\x48\x5e\xfa\x32\xda\xf8\xf4\x05\xe1\x79\x6b\x34\x57\xc5\xa3\x5c\xaf\xd7\xc8\xa6\xb1\x76\x13\x8f\xc0\x08\xe2\x28\x04\x92\x15\x62\xa6\x3b\xc4\x93\x5b\xab\x91\xec\x21\xf6\xe0\x9d\x87\x6e\x47\x06\xa5\xb3\x74\xa0\x66\x74\x62\x08\xba\x0f\xb7\xa3\x03\x24\x74\x1f\xda\xe6\xf1\x84\x10\x1d\x91\xc9\xb6\xe5\x73\x3e\x54\xcb\x9c\x64\x0a\xd4\xe8\xf3\x91\xc3\x54\xcc\x65\x8e\xbc\xfc\xaf\x12\x3c\xd5\x33\x73\xd2\xb6\xce\x56\xc9\xde\x4f\x97\x84\xee\xce\xd1\x35\x90\x69\x2c\x8a\x8f\x03\x45\xdb\x2e\x82\x0b\x2d\x6b\x2d\xc4\x9e\x37\xc1\xd5\x99\x11\x83\x5a\x5d\x55\x2f\x9c\x36\x2a\xe5\x6a\xea\xc3\xcc\x37\xd1\xe0\x36\xb8\xa2\x75\xbc\x99\xea\x16\xf6\xbf\x10\x6e\xd8\xc6\xfa\x0b\x61\x39\xb3\x91\x2f\x83\xdd\xc2\xee\x67\x44\xa4\x3e\xc4\x8d\xc0\x23\x1a\xe8\xe7\x3e\x83\x7a\x90\xa5\x48\xc9\x56\xf1\x7f\x16\x67\x01\xdc\xbb\xae\xee\x13\xd9\xab\x16\x34\x26\x7e\xb9\xdd\x32\x1a\x05\xea\xef\x38\x30\x77\x9e\xf4\xf1\x76\xd5\xd9\x15\x6e\x91\x8c\x0d\xa4\x84\xae\x5b\xaa\x34\x7b\x80\x45\x2f\x47\x6a\x98\x45\x42\x9d\xe6\xb9\x7c\x88\xdb\x48\x10\x2b\x12\xf9\xf0\xd9\x88\xce\x6a\x3c\x0c\x81\x19\x51\x70\x90\xd2\x05\xd5\x50\x66\x6c\x17\xc4\x7a\x6c\x0e\x54\x36\x4e\xe7\x3e\x23\xf8\x92\x2d\x53\x19\x24\x3b\x1b\x04\xb2\x64\x4f\x07\xc0\x55\xf1\x00\x4e\x45\xac\xf7\x5a\x1c\x42\x91\x40\xef\xc9\x5e\x18\x38\xe8\x61\xb9\x87\x5e\x84\x40\x00\x31\x8b\xad\x1b\x44\xf6\x98\x51\x22\x75\x23\xc8\x70\x8c\x5b\x9d\x1d\xba\xa2\xeb\x92\x2f\xf7\xa9\x2f\x3b\xe5\xac\x3a\x2b\x6a\xe2\xeb\x39\x4a\xf8\xf7\xac\x29\x48\xe9\x3d\x13\xe1\x49\x21\x87\xc4\x85\xcc\x8d\x65\x11\x86\x1e\xfe\x53\x08\xa6\xb3\xf5\xb9\x8c\xf9\xbf\x21\x17\x28\x14\xee\x09\x19\xfc\x15\x37\xe7\xf2\xbf\x80\x2f\xff\xe7\xd7\x1b\xf2\x17\x70\xb3\xc0\x64\x0f\x3d\xec\x82\x28\x44\x2c\xa6\xcd\xcd\x42\xfe\xfe\x2d\xa2\x1c\x02\x74\xef\x20\xe4\x22\x37\xfd\x55\xb4\x95\xde\x64\x39\xce\xe2\x86\x9c\x9c\x9c\x20\xee\x9c\x9c\x9c\xdc\x90\xf3\x55\x3c\x5e\x44\xf0\xb7\x08\x25\xa3\x61\x17\x11\x8e\x37\xd8\x91\x5f\x39\xd4\x45\x37\x64\x85\x38\xc4\x9e\xb0\xcd\x69\x20\xa3\x2a\x85\x03\x05\xdd\x17\x26\x19\x82\xaf\x98\xb8\x50\x0e\xbe\xc1\xc8\x73\xc1\xeb\xd4\xba\x79\x0d\xfc\x28\xe4\xe0\x0e\x01\x42\xc9\xf2\x4f\xc4\x28\x10\x47\x22\x9d\x2b\xa1\x1c\x20\x42\xa3\xed\x0e\x70\xbc\xdd\xf1\x10\x70\x0a\x36\x08\xb9\x60\x4b\x83\x1d\x62\x69\xbb\x74\x07\xc1\xeb\x0f\xd4\x7d\x0d\x5c\x8a\xc2\xd7\x1c\xa0\x7b\x1c\xf2\xb8\xc9\xef\xf1\xa8\xfa\x54\x43\x24\xdc\x76\x3a\xd7\x85\x7d\x14\x28\x41\x8e\x27\x72\x43\x27\x9b\x51\xcd\x64\x82\xe6\x86\x38\x29\x29\x55\x23\x11\x5d\xb3\x3e\xde\x33\x56\xa7\x6d\x24\x47\xa2\xfb\x75\x73\xf2\x7d\x83\x8b\xf4\x76\xc7\x79\x30\x92\x6d\x2e\xdd\x51\x1f\xaf\xaf\xd7\xd3\xa3\x8e\x4c\x07\x2a\xc1\xcc\xf9\xaa\x19\x68\x24\x3b\x33\x14\x30\x14\x0a\x5b\x41\xe3\x6c\x25\x35\xed\xe0\x13\x1f\x73\xb5\xf1\x81\xf8\x67\xdc\xb8\x7a\xd7\x0e\x38\x56\x57\x4f\x71\x5b\x5b\x43\xf8\x7f\xea\xcb\xaf\x20\x7d\xdc\xa2\x40\xfc\x98\x66\x42\xb4\x6b\x49\x5d\xb5\xc7\xb3\xde\x93\x3c\xed\x5a\xaf\xda\x4e\xd9\x55\xe9\x98\xa5\x36\xd4\xf9\xaa\x61\xbd\xd3\x64\x76\x55\x2d\xea\xf7\x22\x26\x16\x17\x95\x8b\x0c\x65\x51\x79\x02\x81\x22\x38\x24\xbc\x36\xac\xf2\x58\xd2\x10\xaa\xe8\x54\x9f\xed\x58\x4f\xba\xda\x6f\x62\x6a\x2a\xc6\x78\xac\xae\x84\xa0\x2e\x3d\x73\x84\x9c\x84\x9e\xe8\xdf\x16\x6e\x6c\xe0\x0a\xeb\x39\x83\x6a\xc7\x96\x1a\xf1\x7c\xcf\x63\x15\xd0\xbb\xad\xa8\x83\x30\xe4\x44\xde\xdf\x67\x6f\x81\xd4\xe8\x0a\xd9\xe3\x8f\xe3\x4c\x20\x7f\x05\xb9\x66\xf8\x69\xe8\x70\x4e\x5a\xe8\x10\x50\x37\x1c\x2d\x9a\xd3\x1d\x3b\x11\x83\xe9\x29\xfa\x29\x28\x9c\x17\x22\xd7\x0f\x66\xd5\x3f\xda\xad\xc9\x4a\x6a\x5e\xc3\xad\xf1\x67\xf3\x31\x06\xd3\xb7\xcb\xea\x11\x33\x69\xa1\xfb\x2a\x8b\xef\x3a\xf5\x0c\x78\x66\x91\x87\xc2\x67\x93\xb6\x6a\x33\x4d\x8f\x28\xd8\x47\x9c\xbc\x91\x2e\x87\x2e\x23\xcf\x34\x5f\xe2\x52\xe5\x80\x39\x22\x04\x6a\xbb\xd3\xd3\xda\x55\xa2\x05\xea\x7b\x9f\xd7\x2a\xb8\xfb\x6e\x58\xc5\x7b\x9b\x8f\xad\x72\x7e\xde\xbb\x66\xa0\x13\x97\xda\xc6\xbb\x97\xef\x59\x7f\x6f\x78\xed\xbe\xd5\xfb\xc3\x47\xa4\xad\xa2\x43\x60\x75\x94\x94\x74\x85\x74\xd2\x7a\xba\xe9\x0d\x63\xa2\x29\xff\xaa\x3e\xe9\xfa\x01\xad\x38\x9c\x2d\xa9\x70\xa5\xcc\xe3\x29\x72\x5b\x95\xeb\x54\xd3\x78\xef\xf6\x3b\xfc\xbe\x94\xd0\xef\x77\x27\xb9\xd5\x55\xee\x69\x9a\x0f\x45\x7e\xbb\x9b\xde\x28\x0d\x4b\x03\xfd\xb2\x68\xf4\x2b\x22\x95\x1f\x9b\x17\xae\xe8\x8a\x6d\x98\xdf\x77\xff\x0b\x00\x31\x0d\x2c\xa4\x55\xf8\x9a\xa9\x90\xb4\x12\xd7\x28\x95\x6f\xec\x0d\x46\x84\x42\x55\xc0\x41\x89\x10\xaa\x7d\xa7\x44\xf8\xa3\x42\xd1\x2b\x92\x20\x6b\x13\x13\x20\x7f\x84\x27\x73\x55\x41\xad\xca\xde\x81\x96\x03\x2b\x16\x66\x33\x49\xeb\x70\xf4\x93\xdb\x40\x58\xb5\xda\xd3\x11\x66\x45\x8c\x10\x83\x51\x99\x81\x55\x58\x6b\x14\x22\x26\x2b\x6d\x79\x08\x88\x0f\x84\xda\xc7\x77\x08\x54\x17\x79\x1c\x24\x85\xcb\x87\xf7\xb7\xe8\x9e\xa7\x6e\x94\x8e\x9e\x8c\xd6\x28\xe8\x0b\x78\x9f\x7b\x94\xea\x67\x82\xc9\x14\x33\xc9\x7d\x3a\xf5\x33\xe1\x0c\x6e\x36\xd8\x19\x71\x16\xd7\xc9\x08\xd3\x9a\x6a\xa2\xe3\x5b\x8f\x6e\x31\xe9\xd7\xfd\x27\xd1\x45\x4d\x60\x71\xb9\x62\xa4\x01\x0e\x87\x13\xe6\x01\xaa\x51\xbf\xe1\x88\xf9\x64\x9f\x43\xc4\x24\x1c\x1a\x9a\x43\x9f\x43\x3d\x97\x6c\x3e\xc6\x50\x26\x91\xda\x6c\x58\xbd\xa1\x6e\xc4\xe6\x95\x7e\x7b\x5a\xb1\x79\x47\x63\x6d\x5d\x85\x33\xa5\x71\xe7\xf2\x65\xcf\x72\xfb\xd6\x97\xed\x91\xdc\x59\x1b\xed\x4a\x0c\xac\x19\xba\x44\x1e\x82\x21\x02\x69\x1f\x9d\xf7\xed\x3c\xfc\x43\xd4\x31\x6d\xcc\x0e\xad\xe4\xd3\x64\xe8\x86\xcf\xcd\x40\xb9\xa9\xf3\x2b\xce\xe6\xe0\x7e\xbe\xf3\x20\xd9\x9e\x86\xc8\xdf\xa7\x30\x9e\xed\xa2\x9e\x13\x5f\xda\x40\xc8\x77\x8a\xdd\x00\xb4\x82\xcb\x07\x2a\x89\x85\x62\xcd\xe9\x93\x03\xb7\x45\xb7\x73\xfa\x7b\xa0\xd4\x9e\x68\xd7\x27\x7b\x25\xf7\xd7\x97\x23\xd0\x26\xd9\x39\x2d\x5d\x76\xd2\x50\x9d\x40\x5d\xf2\x48\x81\x28\x72\x88\x42\x41\x8f\x19\xc1\x49\xcc\x30\x21\x47\x84\xcb\x02\x05\x52\xc8\x5d\xe8\x65\x2c\xe6\x57\xe0\x75\x4d\x9b\x42\x3a\xd6\xd4\x2d\x16\xa7\x71\x7b\x5d\x57\xdb\x6c\xf5\x41\xaf\x84\x0a\xde\xb3\x43\x2f\x75\x0a\xb9\x0c\x36\xd9\x7c\x22\x9b\xe3\x90\xcc\xe6\x35\x75\xa7\x4d\x69\x56\x59\xfc\xd8\x73\x99\xc7\x5d\x4b\x29\x89\x79\x46\xd2\x28\x3b\x35\x8d\xe0\xde\x92\x6c\xdc\x0b\xed\x83\x1d\x0c\xbb\x73\xe5\x5a\x7c\x5d\x67\x4d\x73\xc8\xf8\xad\x0c\x9e\x1a\x2b\x57\x57\x8c\x21\xdf\x52\xa8\xe3\x61\xc6\xdb\x05\x43\x40\x5d\x20\x9a\x0e\x2b\x19\xae\xe2\x2e\xcf\xf8\x5c\xcf\x5e\x9b\x45\x9c\x36\x29\x29\x17\x7d\xad\xe0\xa6\xe0\xa3\xbe\xca\x45\xac\x2a\x99\x16\x57\x74\x67\x6a\xf3\x96\x8b\xde\x95\xdf\x8e\xc8\x0c\x5d\x51\x89\x30\xad\x93\x07\x62\x32\x82\x9f\xae\xdf\xad\x01\x65\xe0\xf3\x6a\xfd\xf3\x53\x05\xb7\x1a\xac\x33\xab\xf1\x59\xbf\xce\xb4\x49\x21\xde\x35\xbf\x59\x88\x65\x25\x24\x6e\x2c\x67\x3a\x97\x94\x0a\xa2\x85\x96\x0a\xa0\xca\x0b\x83\x80\xa4\x56\xa9\x1e\xcf\x11\x13\xe0\x8f\x23\xe4\x6a\x65\xb5\x91\x94\x8e\xe5\x33\x26\xe0\x02\x8f\x33\xb9\x19\x0b\xdf\xf2\xb5\x7a\xfb\x31\xd4\x9e\x72\x50\x30\x31\xff\x7d\xe0\x7c\xdd\x7e\x4f\x3c\x8c\x91\x97\xfb\xf2\x9e\x8d\xe8\x5e\x14\x6e\x55\x7c\xa1\xcb\x86\x0a\x8e\x1b\x2a\x98\x0d\xd0\xe3\xe0\x67\x43\xd4\x71\xc0\x4c\x5f\xbe\x68\x8f\xe5\x29\x84\xcf\xb4\xc3\x9d\x6c\x58\x5f\xa1\x63\x60\xb0\x1b\xab\xb6\xe3\x18\x30\x68\xeb\x45\x1e\x5d\xbd\xc8\x17\x88\xe3\xd6\x41\x68\xa5\x8e\x2d\x79\x69\x4b\x5e\xda\x92\x97\xcf\xa9\xe4\xa5\x41\xb8\x72\x31\x52\xb6\x5d\xd7\x53\xc2\x77\xcb\xd9\x1b\x63\x9b\xb5\x07\xe6\x80\xbd\x40\xbb\xd6\x9a\xa2\x56\x29\xb0\xa9\x71\xd3\x25\x02\x15\xa3\xec\xdb\x01\x54\x89\xfc\x57\x00\x34\xc5\xca\x91\x01\x34\x50\x1f\x8a\xb2\x00\x6a\x01\x74\x92\xf8\x93\xea\x1a\x8d\x83\x14\x5e\xc4\xc1\x53\x95\xac\x3a\x5f\x5b\xf3\x74\x46\x92\x68\xd4\x37\xf9\x2a\xc3\x15\xbb\xbe\xc8\xd7\xd3\x96\x4d\x85\xc2\x84\x86\x6c\xd5\x9b\xec\xd5\x84\x1a\xeb\x01\xf6\xb6\x9c\xb3\x54\xb8\xca\xe2\xe3\x85\xf2\x1c\x75\x92\x38\x6f\x9c\x54\x28\xe4\x22\x0f\x93\xa0\xef\x99\x08\x06\x9a\x00\x7d\xe2\xa7\xa9\xe6\x77\x27\x2a\x29\xa8\x64\x81\x98\x12\x3d\xff\x44\x23\x7d\x14\x22\x06\x64\x2e\x9e\xe1\x93\xd8\x93\x24\xf4\xa5\x75\x29\x3b\xef\xe4\xe7\xb4\x83\x39\x6f\x66\x68\x52\x11\x87\x63\xee\x21\x75\x3f\xc3\xda\x90\x0b\xf9\x8c\xab\x88\xba\x10\xfb\x9a\xe5\x4d\x9d\xb4\x69\xad\x3b\xc8\x5c\x73\x25\x55\xb4\x1e\xe0\xb1\xd8\x28\x44\xdd\xfb\x99\xd1\x76\x46\x5e\x23\xf8\x45\x9e\x5a\xbb\x01\xc4\x16\x5c\x57\xeb\x62\x47\x43\xfd\x0d\xd2\x83\x1e\x85\x14\x5f\x77\xe5\xa8\x8f\xf1\xc7\xd5\x8a\x40\x43\xea\x45\xef\xa0\x31\x2d\x33\xa3\x59\x01\xa8\x4d\xe2\xe0\x5e\xd8\x72\xcb\xd2\x9e\xf2\xf7\xe9\x4a\xde\x90\xcc\xf4\x04\x92\x54\xa3\x30\xa8\x2c\x52\xd1\x5a\x42\x4a\xf2\x37\xa7\x80\x45\x24\xd3\x71\x40\x7c\x14\x51\xaf\xb7\xff\xa5\x05\x38\x56\x91\xbd\xa4\xf7\xea\xf2\x76\x8c\x51\xf6\x54\x06\xcb\x7b\x39\x78\xb5\xee\x4f\xf9\xed\xb8\x74\xf9\x83\xf2\x32\x69\xe6\x73\x66\xaf\x8a\x0c\x59\x3c\xa6\xaa\xab\x46\x34\x7d\x92\x48\xad\xe3\x8b\xaa\xaa\xa1\x8e\x59\x21\xd1\xee\x87\xbd\x35\x5a\xcb\x3a\xb1\x67\x5b\x7a\x6d\x6e\xa8\xd0\x16\xff\xae\xb4\xd2\x63\xa0\x24\x52\xf4\x8d\x82\x4f\x7a\x19\xcf\xb7\xa1\x2b\x12\x8d\x6c\x55\x8e\xcb\x98\xd3\x66\x79\xc8\xe1\xc8\xcd\x0b\x6c\xb5\xee\x5b\xd5\x07\xfa\x16\xe6\xaf\xa8\xcb\x6d\x04\x1b\x46\x7d\x71\xf3\xac\xa4\xe8\x0f\x50\x31\xb9\x25\xf8\x54\x2e\x65\x7e\xb4\x36\x2d\xfc\x57\xd9\xbe\xa9\x00\xe0\xb4\x74\xd6\x97\x31\x27\x32\xb7\xdd\x2a\x95\x2b\x96\x3f\xd1\x55\x91\xad\xb7\x6a\x2f\x5e\xec\xc5\xcb\xbc\x0a\xd7\xda\x3b\x91\x81\xef\x44\x66\x27\x19\xd6\x7a\x89\x8e\x5a\xe9\x90\x51\x48\xd4\xf4\x18\x48\x42\x70\xc8\xb6\x88\xdf\x6a\x5d\x4a\x09\x21\xd3\x12\x9f\xfa\x71\xae\x11\xeb\x97\x14\x0a\x97\x54\xad\xde\xac\x1e\x40\x81\x56\xa5\xda\x60\x19\x81\xc7\x59\xc7\xb5\x18\x61\xb6\x65\x58\x92\x53\x7e\xad\xdf\x3f\xb6\xeb\x40\x22\xbd\x75\xc6\x19\xad\xe9\x9b\x0b\xad\x5a\x73\xde\xac\x68\x59\x36\xd6\x0c\x35\x37\x2d\x9b\x5f\x97\x18\x48\x46\x18\x1b\x97\xa5\x97\x28\x66\x77\x14\x4d\x6a\x65\x17\x9b\x0e\x5d\x2a\xbb\x6e\xd7\xea\x2b\x65\x97\xeb\x41\x4e\x76\x33\x9f\xca\xd1\x26\x7a\xa5\x22\x7d\xb9\x04\x2c\x22\x04\x93\x6d\x26\xbe\xbb\xca\x2a\x8e\xfc\xc0\x53\x5f\xe1\xae\x90\x5f\xa4\xa8\x96\xb5\x4b\xaa\x3b\x06\x89\xd3\xbd\xb2\xd7\x6f\xf2\xf3\xea\xbb\x01\xed\x91\xcd\xa7\xf4\x9d\xbe\x27\xfb\xc9\x4a\xc7\x1e\xab\xce\x9e\x9d\xaf\xce\x57\x78\x69\x07\xd5\x77\xd1\xac\x7b\xd8\xc8\xe7\xcb\x4f\xf3\x44\xd0\x84\xa5\x9b\x5f\x93\x55\x5a\x89\xe2\xd8\x29\x36\x74\x7c\x4c\xb6\x8e\x91\xe7\xca\x6d\xf3\xdb\xae\x4f\x22\x8c\xc3\x60\xc7\x64\x43\x6d\xd3\xf2\xb7\xd1\x64\x30\x48\x8f\x4b\xdb\x72\xce\xd3\xe8\x85\x31\x18\xf4\xc7\x1a\xf3\xf2\xec\x62\xd6\x9b\x5e\x19\x13\x54\xb3\xef\x5a\x24\x90\xf0\xdd\x2a\xe6\x77\x7e\x04\xba\xef\x7d\x6d\xc1\xe5\xe1\x38\xb6\x6a\x3f\x2b\x0c\x29\xc3\x64\x84\x8a\x8a\xcf\xf3\xdb\xe2\xeb\xb2\x04\xab\xdb\xe1\xb4\xa9\xc6\xdb\xdf\x77\xd8\xd9\xc9\xc8\x2e\x07\x92\x58\x71\xeb\x13\xad\x37\x5e\x96\xfc\x08\x4a\x8b\x48\x97\x0e\x9f\x32\x47\xba\xc6\x8d\xe6\x15\xb1\xba\xc9\x36\xd3\xf1\x7d\x4a\xed\xec\x39\x6b\x37\xad\x36\xbd\xda\x4e\x45\xcc\xc3\x2d\x9f\x5a\xfb\x30\x1d\x62\x34\xb3\x3e\x9d\xa2\xa1\x5d\x9f\xcd\x67\xd6\x1b\x97\xa2\x9c\xf1\x0e\x6a\x1f\xa8\x5b\x09\xf7\x10\x7b\xe2\x31\xfc\x7c\x2b\xe6\xbf\x9b\x65\x8b\xe4\x88\x77\xb5\xe8\x32\x1c\xda\x05\xd8\xe6\xf6\x90\x61\xee\xa5\x87\x46\x0b\x47\x49\x69\xa5\xfb\xf7\x22\xf1\x1f\x79\x76\x3b\x10\x22\xc7\x3c\xf6\xb1\xa5\x70\x4e\xa9\xca\xb9\xe8\xfe\x80\xcb\x5e\xbd\xa3\xae\x58\x5e\x51\x5e\xa6\x24\x66\xbb\x0b\x6f\x5f\x7f\x68\x7f\x3e\x67\x33\xcf\x3b\x11\x0e\x86\x96\x6c\x95\x8a\xd6\xa5\x74\x15\x25\xe8\x5e\xe4\x8c\x74\x3c\x1f\x9e\x9a\x70\xf2\xe5\xc5\x66\xb7\x48\x8a\x5f\x96\xeb\x9d\x54\x6f\x4d\xda\xb0\xbc\x2b\x91\x7f\x87\x98\x90\xed\x69\x5f\x9d\x5f\xb4\x3a\xb4\x96\xd5\xf4\xd5\x5a\x66\xb4\x7f\x95\xf9\x48\xdd\x32\x88\x6e\x3d\x91\x16\x64\xf3\x88\xcc\xc8\xfe\x81\xd1\xa8\xa9\x40\x66\xd6\x26\x66\x96\xad\xf8\x83\x6e\xe4\x43\xfa\x33\xc9\xf5\xb2\x8f\xb7\x0d\xf6\xf6\x1a\x8a\xe1\x6f\xd4\xb7\xa7\xc4\x51\xba\x10\xe3\x18\x6a\xad\x17\xc9\xa4\x9a\x66\xdc\xab\x16\xb9\x49\xa9\xdb\x78\x90\x86\x62\xe4\x53\x3c\x62\x26\xc7\x88\x19\xaf\x4f\x8d\x07\x31\xca\x4c\x23\x64\xb2\xf3\x51\x2f\x0a\x66\x13\x0f\x50\x3c\xcb\x06\x00\x2a\x5b\xe6\x30\x2a\xcf\xee\x4c\x40\x74\x04\x98\x7a\x21\xb2\xf3\xa2\x04\x9a\x2d\x27\x40\xf3\xba\x0b\x6f\xac\x7a\x1e\x7a\x48\xd5\x23\x81\xef\x19\xee\xa1\xd1\xee\xd5\xec\x5b\x8f\x0d\xdb\x16\xc7\x1e\x67\xbf\x0c\x77\x2a\xa1\xc4\x5c\x37\xea\x23\x82\x2e\x62\x2b\x3d\x61\xad\x21\xb0\x6c\x27\xda\x8b\xa2\x04\xc2\x41\xf3\xef\x65\xdc\xcb\x32\x7f\xc7\x18\x12\x37\xfd\x51\x3e\x19\x96\x7c\x12\x82\x9f\x10\x71\xa8\x8b\xdc\x58\x47\xbc\x83\x21\xfa\xc7\xdf\x7f\xee\x6a\x07\x62\xf5\x01\x9c\x85\x6e\xeb\xe7\x18\xde\xf3\xd5\xe3\x02\x11\x28\x41\x80\x6e\x7e\x05\x37\x52\x67\xb8\x59\xbc\x01\x37\x0b\x86\xa0\x2b\xff\xfa\xce\x30\x47\xf9\x8f\x4b\x19\x40\x2f\x7f\x20\x94\xa0\x9b\x45\x47\x4d\x38\x11\xdb\x6f\x46\x11\x47\x85\x35\xee\x68\xc8\x31\xd9\x2e\x63\xdd\x8e\x11\xe8\x81\x82\x7f\x6c\x06\x8f\x24\x57\xdd\x71\x3e\x03\x5f\x4b\xfb\xab\x9d\x95\x6f\x76\xca\xc7\x23\x81\x43\xfd\x00\x72\x41\xad\x7d\xdf\xa7\x3b\x7f\x8b\xb0\xe7\x8e\x7a\xa7\x57\x79\x2c\x2e\xe0\x7f\x28\x1b\xe1\xc9\xcf\x0b\x4c\x46\xe9\x77\x0d\x79\x7d\x00\x5c\x9f\x7e\x19\x1a\xad\x6a\xc3\x65\xa9\xc8\xaa\xb6\x1b\x4f\xf2\xd4\xa9\x14\x12\x4d\xe7\x5e\x4a\x91\xe5\x12\xec\xe5\x5f\x87\x3e\x8e\x3f\x89\xa6\x9f\x78\xc9\x7c\xfd\x65\xcc\xa6\xcd\xa8\x7f\x59\xb3\xda\x1f\x03\x03\xe8\x60\xfe\x30\x5a\xf4\x4f\xda\xff\xf1\x16\x2b\xac\x98\x9d\x4d\x48\x3f\xe2\x07\xe6\xab\xeb\x86\x4f\x90\xa8\xd8\x54\x7a\x3c\xe4\x94\xc1\x6d\xdf\xa7\x87\x65\x27\xf5\x07\xc3\xbe\x86\x3f\x13\xb5\x4c\x60\x73\x5b\x24\x84\xd2\x4a\xbf\xbe\x96\x12\xab\x6f\x76\x4a\xd2\xcb\x58\xdb\x95\x48\x60\xb3\xad\x4a\x96\x3a\xcb\xcd\xda\x22\x5e\x5d\xca\x60\x85\x36\x30\xf2\xf8\x6f\xd4\x6d\x7a\xc5\xe7\x3d\x13\x6e\xc4\x90\x43\xe2\x42\xe6\x82\x10\x31\x0c\x3d\xfc\xa7\x08\x60\x39\x5b\x9f\x03\x51\xb5\xe8\x86\x5c\xa0\x30\x4c\x02\x15\x1c\x4a\xe2\xe6\x5c\xfe\x17\xf0\xe5\xff\xfc\x7a\x43\xfe\x02\x6e\x16\x98\xec\xa1\x87\x65\xd1\xb7\x98\x7c\x37\x0b\xf9\xfb\xb7\x88\x72\x08\xd0\xbd\x83\x90\x8b\xdc\xf4\x57\xd1\x56\x8a\x4e\x39\xce\xe2\x86\x9c\x9c\x9c\x20\xee\x9c\x9c\x9c\xdc\x90\xf3\x55\x3c\x5e\x44\xf0\xb7\x08\x25\xa3\x61\x17\x11\x8e\x37\xd8\x91\x5f\xc5\xb6\xf5\x0d\x59\x21\x0e\xb1\x27\x9c\x2a\x34\x90\x71\x9f\xe2\x32\x16\xdd\x17\x26\x19\x82\xaf\x98\xb8\x50\x0e\xbe\xc1\xc8\x73\xc1\xeb\xf4\x36\xf3\x35\xf0\xa3\x90\x83\x3b\x04\x08\x25\xcb\x3f\x11\xa3\x60\x0f\xbd\x28\x5b\x01\xa1\x1c\x20\x42\xa3\xed\x0e\x70\xbc\xdd\xf1\x10\x70\x0a\x36\x08\xb9\x60\x4b\x83\x1d\x62\x69\xbb\xac\x7a\xe4\xeb\x0f\xd4\x7d\x0d\x5c\x8a\xc2\xd7\x1c\xa0\x7b\x1c\xf2\xb8\xc9\xef\xf1\xa8\xfa\x54\x43\x24\xfc\x40\x5f\xd1\xc3\x52\x8c\x08\x02\x88\xfb\x38\xef\x5c\x49\x8e\x27\x0a\x51\x4c\x36\xa3\x1a\x82\x05\xcd\xdb\xa4\x5b\x4e\x22\xd5\xec\xcb\x42\x7a\x15\x42\xc9\xee\x62\xe2\x89\x6d\xae\x96\x7c\x13\x85\xf4\xeb\xcc\x5f\xc5\xcd\xce\x0e\xb1\xba\x67\x23\x5a\xbd\x14\xef\x19\x93\xec\xa0\xd0\x44\xe3\x8c\xf3\x95\xf9\xf2\x83\x86\x45\xc7\x1c\x52\xfa\xb5\x72\x3e\xff\xc4\xc4\x2d\xcc\x28\xfe\xb8\x7a\x3b\x0c\x22\xae\xdb\x74\xf7\x3e\xb4\x2e\x6d\x5a\xac\xd5\x60\xc3\x85\x5e\x95\x28\x9f\xa6\x9a\x9e\xaf\x0c\x56\x5b\xa1\xb6\xf4\x5e\xca\x98\x67\x31\x01\xcb\xce\x5a\x58\x22\x2b\x9a\xf4\xdb\xdb\x1d\xe7\xc1\x48\xe1\x29\x52\x91\xfd\x78\x7d\xbd\x6e\x15\xd8\x5b\x44\x96\x1e\x95\xb2\x44\xf0\x4a\x80\x18\x2c\x84\x6c\x76\x20\x68\x3c\x9a\x18\x69\x11\x40\x06\x7d\xc4\x55\xbd\x53\x28\x7e\xe7\x2b\xe9\x83\x36\xb9\x98\x54\x78\x22\x52\x7c\xc0\xa9\x79\xb6\x48\x1c\xcf\x0a\xff\x2f\xb0\x74\x6c\xca\x21\xaa\xdc\xca\x9c\x45\xa8\xe4\x16\xcf\xcc\x2a\xf3\xb9\x15\xe6\x90\x1b\x66\x3d\xa7\x72\x49\xbd\x8a\x59\x20\x12\xf9\xba\x4f\x3c\x56\x33\x34\x27\xb8\xeb\xe7\x86\xd5\x97\x37\x07\x4e\x3f\x1e\xb5\xd3\xcc\xb3\xfd\x4e\xf3\x7b\x95\xed\x96\xe0\x50\xaf\x81\xc9\x83\x53\xc0\x90\x45\xe8\xec\x90\x0f\x4d\xca\x17\x19\x29\x80\xaf\x54\x36\x10\xd3\x7d\xf5\xf8\xff\x01\x00\x00\xff\xff\x43\x49\xf8\xc8\x55\xc8\x01\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
