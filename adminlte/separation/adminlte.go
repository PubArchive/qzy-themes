package separation

import (
	"os"

	"github.com/hongweikkx/qzy-admin/modules/config"
	adminTemplate "github.com/hongweikkx/qzy-admin/template"
	"github.com/hongweikkx/qzy-admin/template/components"
	"github.com/hongweikkx/qzy-admin/template/types"
	"github.com/hongweikkx/qzy-themes/adminlte/resource"
	"github.com/hongweikkx/qzy-themes/common"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Adminlte = Theme{
	ThemeName: "adminlte_sep",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: common.SepTemplateList,
			Separation:   true,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: common.SepTemplateList,
		Separation:   true,
	},
}

func init() {
	adminTemplate.Add("adminlte_sep", &Adminlte)
}

func Get() *Theme {
	return &Adminlte
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return common.SepTemplateList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	return os.ReadFile(config.GetAssetRootPath() + path)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}
