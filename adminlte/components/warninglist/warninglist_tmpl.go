package warninglist

var List = map[string]string{
	"warninglist": `{{define "warninglist"}}
<ul class="products-list product-list-in-box">
    {{range $key, $data := .Data}}
    <li class="item">       
        <div style="padding-left:10px;">
            <i class="fa {{index $data "icon"}}" style="color:{{index $data "iconcolor"}};"></i> 
            <a href="{{index $data "link"}}" class="product-title">
                {{index $data "title"}}
                <span class="label label-{{index $data "labeltype"}} pull-right">
                    {{index $data "count"}}
                </span>
            </a>
            <span class="product-description">
            latest:&nbsp;{{index $data "last"}}
            </span>
        </div>
    </li>
    {{end}}
</ul>
{{end}}`,
}
