package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"purple_basic_go/3-bin/api"
	"purple_basic_go/3-bin/bins"
	"purple_basic_go/3-bin/file"
	"purple_basic_go/3-bin/storage"
)

// простой in-memory storage для мокового jsonbin
type memStore struct {
	byID map[string]map[string]any
}

func newMemStore() *memStore { return &memStore{byID: map[string]map[string]any{}} }

// поднимаем httptest.Server, эмулирующий jsonbin v3
func startMockJSONBin(t *testing.T, ms *memStore) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()

	// POST /b — create
	mux.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		if r.Header.Get("X-Master-Key") == "" {
			http.Error(w, "missing key", 401)
			return
		}
		var body struct {
			Metadata map[string]any `json:"metadata"`
			Record   map[string]any `json:"record"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		// генерируем простой id
		id := "id_" + strings.ReplaceAll(body.Metadata["name"].(string), " ", "_")
		ms.byID[id] = body.Record

		resp := map[string]any{
			"metadata": map[string]any{"id": id, "name": body.Metadata["name"]},
			"record":   body.Record,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// GET /b/{id}/latest
	mux.HandleFunc("/b/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/b/")
		if strings.HasSuffix(path, "/latest") && r.Method == http.MethodGet {
			id := strings.TrimSuffix(path, "/latest")
			rec, ok := ms.byID[id]
			if !ok {
				http.Error(w, "not found", 404)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"record": rec})
			return
		}
		// PUT /b/{id}
		if r.Method == http.MethodPut {
			id := path
			if _, ok := ms.byID[id]; !ok {
				http.Error(w, "not found", 404)
				return
			}
			var rec map[string]any
			if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			ms.byID[id] = rec
			w.WriteHeader(200)
			return
		}
		// DELETE /b/{id}
		if r.Method == http.MethodDelete {
			id := path
			if _, ok := ms.byID[id]; !ok {
				http.Error(w, "not found", 404)
				return
			}
			delete(ms.byID, id)
			w.WriteHeader(200)
			return
		}
		http.Error(w, "bad request", 400)
	})

	srv := httptest.NewServer(mux)
	return srv
}

func newTestApp(t *testing.T) (*bins.Service, func()) {
	t.Helper()

	// temp файл для локального списка
	dir := t.TempDir()
	dataFile := filepath.Join(dir, "bins_local.json")

	// x.env
	os.Setenv("JSONBIN_KEY", "test-key")
	os.Setenv("BINS_DATA_FILE", dataFile)

	// локальное хранилище как в проекте
	fm := file.LocalFileManager{}
	st := storage.JSONStorage{FileManager: fm}

	// мок jsonbin
	ms := newMemStore()
	srv := startMockJSONBin(t, ms)
	os.Setenv("JSONBIN_BASE", srv.URL)

	cl := api.NewClient()
	svc := bins.NewService(st, dataFile, cl)

	cleanup := func() {
		srv.Close()
		_ = os.Remove(dataFile)
	}
	return svc, cleanup
}

// ---- ТЕСТЫ ----

func TestCreateBin(t *testing.T) {
	svc, cleanup := newTestApp(t)
	defer cleanup()

	// подготовим временный файл с данными
	dir := t.TempDir()
	recPath := filepath.Join(dir, "rec.json")
	_ = os.WriteFile(recPath, []byte(`{"hello":"world"}`), 0644)

	// вызываем CLI
	var out bytes.Buffer
	err := RunCLI(context.Background(), svc, []string{"create", "--file", recPath, "--name", "my bin", "--private"}, &out)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	if !strings.Contains(out.String(), "created: id=id_my_bin") {
		t.Fatalf("unexpected output: %s", out.String())
	}

	// проверь, что локальный список пополнился
	list := svc.GetBins()
	if len(list) != 1 || list[0].ID != "id_my_bin" || list[0].Name != "my bin" {
		t.Fatalf("local list not updated: %#v", list)
	}

	// cleanup: удаляем удалённо
	_ = RunCLI(context.Background(), svc, []string{"delete", "--id", "id_my_bin"}, &bytes.Buffer{})
}

func TestUpdateBin(t *testing.T) {
	svc, cleanup := newTestApp(t)
	defer cleanup()

	// сначала создадим bin
	dir := t.TempDir()
	recPath := filepath.Join(dir, "rec.json")
	_ = os.WriteFile(recPath, []byte(`{"v":1}`), 0644)

	var out bytes.Buffer
	if err := RunCLI(context.Background(), svc, []string{"create", "--file", recPath, "--name", "u1"}, &out); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	id := "id_u1"

	// обновим
	rec2 := filepath.Join(dir, "rec2.json")
	_ = os.WriteFile(rec2, []byte(`{"v":2,"x":"y"}`), 0644)
	out.Reset()
	if err := RunCLI(context.Background(), svc, []string{"update", "--id", id, "--file", rec2}, &out); err != nil {
		t.Fatalf("update failed: %v", err)
	}
	if !strings.Contains(out.String(), "updated: id_u1") {
		t.Fatalf("unexpected output: %s", out.String())
	}

	// cleanup
	_ = RunCLI(context.Background(), svc, []string{"delete", "--id", id}, &bytes.Buffer{})
}

func TestGetBin(t *testing.T) {
	svc, cleanup := newTestApp(t)
	defer cleanup()

	// создаём
	dir := t.TempDir()
	recPath := filepath.Join(dir, "rec.json")
	_ = os.WriteFile(recPath, []byte(`{"a":123}`), 0644)
	_ = RunCLI(context.Background(), svc, []string{"create", "--file", recPath, "--name", "g1"}, &bytes.Buffer{})
	id := "id_g1"

	// get
	var out bytes.Buffer
	if err := RunCLI(context.Background(), svc, []string{"get", "--id", id}, &out); err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if !strings.Contains(out.String(), `"a": 123`) {
		t.Fatalf("unexpected body: %s", out.String())
	}

	// cleanup
	_ = RunCLI(context.Background(), svc, []string{"delete", "--id", id}, &bytes.Buffer{})
}

func TestDeleteBin(t *testing.T) {
	svc, cleanup := newTestApp(t)
	defer cleanup()

	// создаём
	dir := t.TempDir()
	recPath := filepath.Join(dir, "rec.json")
	_ = os.WriteFile(recPath, []byte(`{"d":true}`), 0644)
	_ = RunCLI(context.Background(), svc, []string{"create", "--file", recPath, "--name", "d1"}, &bytes.Buffer{})
	id := "id_d1"

	// удаляем
	var out bytes.Buffer
	if err := RunCLI(context.Background(), svc, []string{"delete", "--id", id}, &out); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	if !strings.Contains(out.String(), "deleted: id_d1") {
		t.Fatalf("unexpected output: %s", out.String())
	}

	// повторный get должен падать
	if err := RunCLI(context.Background(), svc, []string{"get", "--id", id}, &bytes.Buffer{}); err == nil {
		t.Fatalf("expected error on get after delete, got nil")
	}
}
