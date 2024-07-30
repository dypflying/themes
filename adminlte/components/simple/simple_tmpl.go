package simple

var List = map[string]string{
	"simple": `{{define "simple"}}
<div class="description-block border-{{.Border}}" style="text-align:center;">
    <h4 class="description-header text-{{.Color}}">
    {{if ne .Url ""}}
      <a href="{{.Url}}">&nbsp;&nbsp;{{langHtml .Number}}&nbsp;&nbsp;</a>
    {{else}}
      {{langHtml .Number}}
    {{end}}
    </h4>
    <span style="color:#999;">
      {{langHtml .Title}}
    </span>
</div>
{{end}}`,
}
