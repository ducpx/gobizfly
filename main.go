package main

import (
	"bizflycloud/gobizfly"
	"context"
	"fmt"
	"log"
	"time"
)

const (
	host     = "https://manage.bizflycloud.vn"
	username = "vultr@vccloud.vn"
	password = "it1qVRzj9Q8CYLtN"
)

func main() {
	client, err := gobizfly.NewClient(
		gobizfly.WithAPIUrl(host),
		gobizfly.WithTenantName(username),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	tok, err := client.Token.Create(ctx, &gobizfly.TokenCreateRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	client.SetKeystoneToken(tok.KeystoneToken)
	fmt.Println(tok.KeystoneToken)
	//Create volume
	//volumeCreateReq := &gobizfly.VolumeCreateRequest{
	//	Name:             "ducpx",
	//	Size:             100,
	//	VolumeType:       "HDD",
	//	AvailabilityZone: "HN1",
	//	SnapshotID:       "",
	//}
	//volume, err := client.Volume.Create(ctx, volumeCreateReq)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Volume created")
	//fmt.Printf("Id: %s, name: %s, size: %d, volume type: %s, available zone: %s, snapshot id: %s",
	//	volume.ID, volume.Name, volume.Size, volume.VolumeType, volume.AvailabilityZone, volume.SnapshotID)

	//Volume list
	//volumes, err := client.Volume.List(ctx, &gobizfly.ListOptions{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%#v\n", volumes)

	//Volume get
	//volume, err := client.Volume.Get(ctx, "9d9b4c3f-f442-484f-844f-7418cab34e33")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%#v\n", volume)

	// List snapshots
	snapshots, err := client.Snapshot.List(ctx, &gobizfly.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, snapshot := range snapshots {

		fmt.Println(snapshot.Id)

	}

	// Snapshot create
	//scr := &gobizfly.SnapshotCreateRequest{
	//	Name:     "ducpx-snapshot",
	//	VolumeId: volume.ID,
	//	Force:    true,
	//}
	//snapshot, err := client.Snapshot.Create(ctx, scr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%#v\n", snapshot.Id)
	// snapshot id: 8fb3aeda-6ee4-4183-8783-2788e8a2bbfa

	// delete

	//fmt.Println("delete: f6fe2287-3288-4821-b768-308dbcb6ca85")
	//if err := client.Snapshot.Delete(ctx, "f6fe2287-3288-4821-b768-308dbcb6ca85"); err != nil {
	//	log.Fatal(err)
	//}
	// Get snapshot
	//snapshot, err := client.Snapshot.Get(ctx, "8fb3aeda-6ee4-4183-8783-2788e8a2bbfa")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%#v\n", snapshot)
}
