package skyhdd

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

type demo struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

const gcsBucket = "learning-1130.appspot.com"

func handler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(req)

	u := user.Current(ctx)
	if u == nil {
		url, err := user.LoginURL(ctx, req.URL.String())
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(res, "Hello, %v\n\n", u)

	gcsClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler NewClient: ", err)
		return
	}
	defer gcsClient.Close()

	d := &demo{
		ctx:    ctx,
		res:    res,
		client: gcsClient,
		bucket: gcsClient.Bucket(gcsBucket),
	}

	d.delFiles()
	d.createFiles()
	d.listFiles()
}

func (d *demo) listFiles() {
	io.WriteString(d.res, "OBJECTS\n")

	objs, err := d.bucket.List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		fmt.Fprintf(d.res, "%v - %v - %v\n", obj.Name, obj.ACL, obj.MediaLink)
	}
}

func (d *demo) createFiles() {
	for _, n := range []string{"foo"} {
		d.createFile(n)
	}
}

func (d *demo) createFile(fileName string) {

	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"
	wc.ACL = []storage.ACLRule{
		{"user-toddmcleod@gmail.com", storage.RoleReader},
	}
	//wc.ACL = []storage.ACLRule{
	//	{"user-<some-email-here>@gmail.com", storage.RoleReader},
	//}

	if _, err := wc.Write([]byte("You accessed the file\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Errorf(d.ctx, "createFile: unable to close bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
}

func (d *demo) delFiles() {
	objs, err := d.bucket.List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		if err := d.bucket.Object(obj.Name).Delete(d.ctx); err != nil {
			log.Errorf(d.ctx, "deleteFiles: unable to delete bucket %q, file %q: %v", d.bucket, obj.Name, err)
			return
		}
	}
}
