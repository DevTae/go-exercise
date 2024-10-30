package main

import (
	"fmt"
	"github.com/devtae/go-exercise/interfaces"
)

type GuestbookReconciler struct {
	interfaces.Client
}

func main() {
	g := GuestbookReconciler{interfaces.ReconcileClient}

	// 각각 출력 진행
	g.Get("Taehyeon")
	g.Client.Get("Taehyeon") // Client 에 대한 interface 접근
}