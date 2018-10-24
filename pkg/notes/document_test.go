package notes

import (
	"os"
	"testing"

	"github.com/kolide/kit/logutil"
	"github.com/stretchr/testify/require"
)

func TestDocument(t *testing.T) {
	client := githubClient(t)
	logger := logutil.NewCLILogger(true)

	notes, err := ListReleaseNotes(client, logger, v1_11_0, "e92ea04edb286efe76caa86183fc00850a936f74")
	require.NoError(t, err)
	require.Len(t, notes, 14)

	doc, err := CreateDocument(notes)
	require.NoError(t, err)

	require.NoError(t, RenderMarkdown(doc, os.Stdout))
}

func TestPrettySIG(t *testing.T) {
	cases := map[string]string{
		"scheduling":        "Scheduling",
		"cluster-lifecycle": "Cluster Lifecycle",
		"cli":               "CLI",
		"aws":               "AWS",
		"api-machinery":     "API Machinery",
		"vsphere":           "vSphere",
		"openstack":         "OpenStack",
	}

	for input, expected := range cases {
		require.Equal(t, expected, (prettySIG(input)))
	}
}
