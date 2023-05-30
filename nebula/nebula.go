package nebula

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/synchthia/nebula-api/nebulapb"
	"google.golang.org/grpc"
)

var cConn *grpc.ClientConn
var client nebulapb.NebulaClient

func NewClient() {
	address := os.Getenv("NEBULA_ADDRESS")
	if len(address) == 0 {
		address = "localhost:17200"
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatalf("[Nebula] Failed connect to Nebula-API")
		return
	}
	logrus.WithFields(logrus.Fields{
		"address": address,
	}).Debugf("[Nebula] Connecting to Nebula-API")
	//defer conn.Close()

	cConn = conn
	client = nebulapb.NewNebulaClient(conn)
}

func Shutdown() {
	cConn.Close()
}

// IPLookup - Lookup IP
func IPLookup(ip string) (*nebulapb.IPLookupResult, error) {
	res, err := client.IPLookup(context.Background(), &nebulapb.IPLookupRequest{
		IpAddress: ip,
	})
	if err != nil {
		return nil, err
	}

	return res.GetResult(), nil
}

// SetFavicon - Set Proxy Favicon
func SetFavicon(url string) error {
	resp, err1 := http.Get(url)
	if err1 != nil {
		return err1
	}

	defer resp.Body.Close()

	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return err2
	}

	encoded := base64.StdEncoding.EncodeToString(b)

	favicon := fmt.Sprintf("data:image/png;base64,%s", encoded)
	_, err := client.SetFavicon(context.Background(), &nebulapb.SetFaviconRequest{Favicon: favicon})
	if err != nil {
		return err
	}

	return nil
}

// SetMotd - Set Proxy Motd
func SetMotd(motd string) error {
	_, err := client.SetMotd(context.Background(), &nebulapb.SetMotdRequest{Motd: motd})
	return err
}

func GetBungeeConfig() (*nebulapb.BungeeEntry, error) {
	response, err := client.GetBungeeEntry(context.Background(), &nebulapb.GetBungeeEntryRequest{})
	return response.Entry, err
}

// AddServerEntry - Add Server to Database
func AddServerEntry(name, displayName, address string, port int32, fallback bool) error {
	logrus.Printf("[Server] Request..")
	e := nebulapb.ServerEntry{
		Name:        name,
		DisplayName: displayName,
		Address:     address,
		Port:        port,
		Motd:        "",
		Fallback:    fallback,
	}

	pp.Println(e)
	_, err := client.AddServerEntry(context.Background(), &nebulapb.AddServerEntryRequest{Entry: &e})

	return err
}

func RemoveServerEntry(name string) error {
	logrus.Printf("[Server] Remove")
	_, err := client.RemoveServerEntry(context.Background(), &nebulapb.RemoveServerEntryRequest{Name: name})

	return err
}

func ListServers() ([]*nebulapb.ServerEntry, error) {
	res, err := client.GetServerEntry(context.Background(), &nebulapb.GetServerEntryRequest{})
	if err != nil {
		return nil, err
	}

	return res.Entry, nil
}
