package featurecollection

import (
	"bufio"
	"bytes"
	"context"
	"github.com/paulmach/orb/geojson"
	"github.com/whosonfirst/go-writer/v2"
	"strings"
	"testing"
)

func TestFeatureCollectionWriter(t *testing.T) {

	features := []string{
		`{"type":"Feature", "properties":{"id":1}, "geomemetry": {"type": "Point", "geometry": [ 0.0, 0.0 ] }}`,
		`{"type":"Feature", "properties":{"id":2}, "geomemetry": {"type": "Point", "geometry": [ 0.0, 0.0 ] }}`,
		`{"type":"Feature", "properties":{"id":3}, "geomemetry": {"type": "Point", "geometry": [ 0.0, 0.0 ] }}`,
	}

	ctx := context.Background()

	var buf bytes.Buffer
	buf_wr := bufio.NewWriter(&buf)

	ctx, err := writer.SetIOWriterWithContext(ctx, buf_wr)

	if err != nil {
		t.Fatal(ctx)
	}

	writer_uri := "featurecollection://?writer=io://"

	wr, err := writer.NewWriter(ctx, writer_uri)

	if err != nil {
		t.Fatalf("Failed to create new writer for '%s', %v", writer_uri, err)
	}

	for _, f := range features {

		sr := strings.NewReader(f)

		_, err := wr.Write(ctx, "", sr)

		if err != nil {
			t.Fatalf("Failed to write feature, %v", err)
		}
	}

	err = wr.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close feature collection writer, %v", err)
	}

	buf_wr.Flush()

	_, err = geojson.UnmarshalFeatureCollection(buf.Bytes())

	if err != nil {
		t.Fatalf("Failed to unmarshal feature collection, %v", err)
	}

}
