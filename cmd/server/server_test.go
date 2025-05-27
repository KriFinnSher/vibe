package main

import (
	"context"
	tt "fullGen/tintin"
	"testing"
)

var tintin = tt.Tintin{
	ID:         "123",
	Name:       "tintin",
	Occupation: tt.OptString{Value: "journalist", Set: true},
	Age:        tt.OptInt{Value: 24, Set: true},
}

func TestTT_AllOperations(t *testing.T) {
	client := CreateClient()
	ctx := context.Background()

	// === Create ===
	tintinsPost, err := client.TintinsPost(ctx, &tintin)
	if err != nil {
		t.Fatalf("create failed: %s", err)
	}
	if tintinsPost == nil {
		t.Fatalf("create returned nil")
	}
	if *tintinsPost != tintin {
		t.Fatalf("created tintin doesn't match input")
	}

	// === Duplicate Create ===
	tintinsPostDup, err := client.TintinsPost(ctx, &tintin)
	if err == nil {
		t.Fatalf("expected error when creating duplicate tintin, got none")
	}
	if tintinsPostDup != nil {
		t.Fatalf("expected nil on duplicate create, got: %+v", tintinsPostDup)
	}

	// === Get by ID ===
	tintinGet, err := client.TintinsIDGet(ctx, tt.TintinsIDGetParams{ID: tintin.ID})
	if err != nil {
		t.Fatalf("get by id failed: %s", err)
	}
	if tintinGet == nil {
		t.Fatalf("get by id returned nil")
	}
	resp, ok := tintinGet.(*tt.Tintin)
	if !ok {
		t.Fatalf("unexpected type from get: %T", tintinGet)
	}
	if *resp != tintin {
		t.Fatalf("get by id result doesn't match original")
	}

	// === Get all ===
	all, err := client.TintinsGet(ctx)
	if err != nil {
		t.Fatalf("get all failed: %s", err)
	}
	found := false
	for _, item := range all {
		if item.ID == tintin.ID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("tintin not found in get all")
	}

	// === Update ===
	updated := *tintinsPost
	updated.Name = "Tintin Updated"
	updated.Age = tt.OptInt{Value: 25, Set: true}
	updatedRes, err := client.TintinsIDPut(ctx, &updated, tt.TintinsIDPutParams{ID: updated.ID})
	if err != nil {
		t.Fatalf("update failed: %s", err)
	}
	if updatedRes == nil {
		t.Fatalf("update returned nil")
	}
	resp, ok = updatedRes.(*tt.Tintin)
	if !ok {
		t.Fatalf("unexpected type from put: %T", tintinGet)
	}
	if resp.Name != "Tintin Updated" || resp.Age.Value != 25 {
		t.Fatalf("update didn't apply changes")
	}

	// === Delete ===
	delRes, err := client.TintinsIDDelete(ctx, tt.TintinsIDDeleteParams{ID: tintin.ID})
	if err != nil {
		t.Fatalf("delete failed: %s", err)
	}
	_, ok = delRes.(*tt.TintinsIDDeleteNoContent)
	if !ok {
		t.Fatalf("expected TintinsIDDeleteNoContent, got: %#v", delRes)
	}

	// === Get after delete ===
	deleted, err := client.TintinsIDGet(ctx, tt.TintinsIDGetParams{ID: tintin.ID})
	if err == nil {
		t.Fatalf("expected error when getting deleted tintin")
	}
	if deleted != nil {
		t.Fatalf("expected nil when getting deleted tintin, got: %+v", deleted)
	}
}
