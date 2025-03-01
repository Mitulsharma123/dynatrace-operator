package secrets

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSecretFileContent = `apiUrl: apiUrl
apiToken: apiToken
`

func TestNewFromConfig(t *testing.T) {
	fs := afero.NewMemMapFs()
	workingDir, err := os.Getwd()

	require.NoError(t, err)

	secretsPath := path.Join(workingDir, "..", "testdata", "Secrets")
	require.NoError(t, fs.MkdirAll(secretsPath, 0655))

	require.NoError(t, afero.WriteFile(fs, path.Join(secretsPath, "Secrets-test.yaml"),
		[]byte(testSecretFileContent), 0755))

	tenantSecrets, err := NewFromConfig(fs, path.Join(secretsPath, "Secrets-test.yaml"))

	assert.NoError(t, err)
	assert.Equal(t, "apiUrl", tenantSecrets.ApiUrl)
	assert.Equal(t, "apiToken", tenantSecrets.ApiToken)
}
