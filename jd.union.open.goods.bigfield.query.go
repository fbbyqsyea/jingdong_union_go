package jingdong_union_go

import (
	"encoding/json"
	"errors"
)

type JdUnionOpenGoodsBigfieldQueryTopLevel struct {
	JdUnionOpenGoodsBigfieldQueryResponse JdUnionOpenGoodsBigfieldQueryResponse `json:"jd_union_open_goods_bigfield_query_responce"`
}

type JdUnionOpenGoodsBigfieldQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type VideoBigFieldInfo struct {
	ProductFeatures     string `json:"productFeatures"`
	Image               string `json:"image"`
	MaterialDescription string `json:"material_Description"`
	Comments            string `json:"comments"`
	BoxContents         string `json:"box_Contents"`
	EditerDesc          string `json:"editerDesc"`
	ContentDesc         string `json:"contentDesc"`
	Manual              string `json:"manual"`
	Catalogue           string `json:"catalogue"`
}

type BaseBigFieldInfo struct {
	Wdis       string `json:"wdis"`
	WareQD     string `json:"wareQD"`
	PropGroups string `json:"propGroups"`
	PropCode   string `json:"propCode"`
}
type BookBigFieldInfo struct {
	ProductFeatures string `json:"productFeatures"`
	Image           string `json:"image"`
	Comments        string `json:"comments"`
	RelatedProducts string `json:"relatedProducts"`
	AuthorDesc      string `json:"authorDesc"`
	BookAbstract    string `json:"bookAbstract"`
	EditerDesc      string `json:"editerDesc"`
	ContentDesc     string `json:"contentDesc"`
	Catalogue       string `json:"catalogue"`
	Introduction    string `json:"introduction"`
}
type JdUnionOpenGoodsBigfieldQueryResult struct {
	Code    int                 `json:"code"`
	Data    []BigFieldGoodsResp `json:"data"`
	Message string              `json:"message"`
}

type BigFieldGoodsResp struct {
	VideoBigFieldInfo VideoBigFieldInfo `json:"videoBigFieldInfo"`
	Owner             string            `json:"owner"`
	MainSkuId         string            `json:"mainSkuId"`
	ProductId         string            `json:"productId"`
	SkuStatus         string            `json:"skuStatus"`
	ImageInfo         ImageInfo         `json:"imageInfo"`
	CategoryInfo      CategoryInfo      `json:"categoryInfo"`
	SkuName           string            `json:"skuName"`
	ItemID            string            `json:"itemId"`
	BaseBigFieldInfo  BaseBigFieldInfo  `json:"baseBigFieldInfo"`
	BookBigFieldInfo  BookBigFieldInfo  `json:"bookBigFieldInfo"`
	DetailImages      string            `json:"detailImages"`
	SkuID             string            `json:"skuId"`
}

func (app *App) JdUnionOpenGoodsBigfieldQuery(params map[string]interface{}) (result *JdUnionOpenGoodsBigfieldQueryResult, err error) {

	body, err := app.Request("jd.union.open.goods.bigfield.query", map[string]interface{}{"goodsReq": params})
	resp := &JdUnionOpenGoodsBigfieldQueryTopLevel{}
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenGoodsBigfieldQueryResponse.Result != "" {
		result = &JdUnionOpenGoodsBigfieldQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenGoodsBigfieldQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
