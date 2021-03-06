// pbvalidate validates pbjson files against a protobuf message
//
// Usage:
//
//    pbvalidate -f $workspace/auth/api/config.proto -I /,$workspace,$workspace/vendor/github.com/googleapis/googleapis -m auth.NamespaceConfig  /tmp/namespace_config.json
//
package main

import (
	"flag"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/juju/errors"
	"github.com/mkmik/stringlist"
)

var (
	protoFileName = flag.String("f", "", "Proto schema files")
	protoMessage  = flag.String("m", "", "Proto message")
	importPaths   = stringlist.Flag("I", "Proto include path")
)

func findMessage(fds []*desc.FileDescriptor, name string) *desc.MessageDescriptor {
	for _, fd := range fds {
		if md := fd.FindMessage(name); md != nil {
			return md
		}
	}
	return nil
}

func run(fileName string, protoMessage string, importPaths []string, src string) error {
	p := &protoparse.Parser{
		ImportPaths: importPaths,
	}
	fds, err := p.ParseFiles(fileName)
	if err != nil {
		return errors.Trace(err)
	}
	md := findMessage(fds, protoMessage)

	if md == nil {
		return errors.Errorf("cannot find message %q", protoMessage)
	}
	m := dynamic.NewMessage(md)

	b, err := ioutil.ReadFile(src)
	if err != nil {
		return errors.Trace(err)
	}
	if err := m.UnmarshalJSON(b); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func main() {
	defer glog.Flush()

	flag.Parse()
	src := flag.Arg(0)

	if err := run(*protoFileName, *protoMessage, *importPaths, src); err != nil {
		glog.Infof("%+v", err)
		glog.Exitf("%v", err)
	}
}
