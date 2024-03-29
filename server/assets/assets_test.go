/*Package assets 内嵌外部资源测试
 * go test -v
 * go test -v assets_test.go assets.go
 */
package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedFile(t *testing.T) {
	data, err := Assets.ReadFile(".gitignore")
	if err != nil {
		t.Fatalf("读取文件失败, err: %#v\n", err)
	}
	assert.Equal(t, data, []byte("!.gitignore"))
}

func TestEmbedDirPath(t *testing.T) {
	_, err := WebAssets.ReadDir("dist")
	if err != nil {
		t.Errorf("读取文件失败, err: %#v\n", err)
	}
}

func TestWebStaticAssets(t *testing.T) {
	data, err := WebAssets.ReadDir("dist/static")
	if err != nil {
		t.Errorf("读取文件失败, err: %#v\n", err)
	}
	assert.Equal(t, len(data) > 0, true)
}
