/*
Copyright © 2018, Oracle and/or its affiliates. All rights reserved.

The Universal Permissive License (UPL), Version 1.0
*/

// Sample grpc client

package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "k8s.io/apiserver/pkg/storage/value/encrypt/envelope/v1beta1"
)

func main() {
	fmt.Println("client start--")
	// nolint:staticcheck
	// TODO: fix deprecated usage
	connection, err := grpc.Dial("unix:////tmp/kms/socketfile.sock", grpc.WithInsecure(), grpc.WithTimeout(30*time.Second), grpc.WithDialer(unixDial))
	// nolint:staticcheck
	defer connection.Close()
	if err != nil {
		fmt.Println("Connection to KMS plugin failed, error", err)
	}

	kmsClient := pb.NewKeyManagementServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	plain := []byte("test data")
	request := &pb.EncryptRequest{Plain: plain, Version: "v1beta1"}
	response, err2 := kmsClient.Encrypt(ctx, request)
	if err2 != nil {
		fmt.Println("request error", err2)
	}

	cipher := response.Cipher
	decryptReq := &pb.DecryptRequest{Cipher: cipher, Version: "v1beta1"}

	fmt.Printf("Encrypt response: %s\n", response.Cipher)
	decryptRes, _ := kmsClient.Decrypt(ctx, decryptReq)
	fmt.Printf("Decrypt response: %s\n", decryptRes.Plain)
}

// This dialer explicitly ask gRPC to use unix socket as network.
func unixDial(addr string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", addr, timeout)
}
