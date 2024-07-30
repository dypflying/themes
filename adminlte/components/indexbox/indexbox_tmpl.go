package indexbox

var List = map[string]string{
	"indexbox": `{{define "indexbox"}}
    <div class="small-box" style="background-color:#FFFFFF;border-top: 2px solid rgb(200,200,200);">
        <div class="inner" style="text-align:center;color:#6D85C3;">
        <a href="{{.Url}}" >
            <p style="margin-bottom:5px;font-weight:bold;">
                {{langHtml .Title}}
            </p>
            <p style="margin-bottom:5px;">               
                {{langHtml .Value}}
            </p>
            <h4 style="margin-bottom:10px;"> 
                <i class="fa {{.Icon}}" style="color:{{.IconColor}};"></i> 
            </h4>
            </a>
        </div>

    </div>
{{end}}`,
}
