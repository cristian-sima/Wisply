<div class="alert alert-dismissible alert-warning">
  <button type="button" class="close" data-dismiss="alert">×</button>
  <strong>Sorry...</strong> {{.messageContent}}. Go <a href="{{.messageLink}}" >back</a>.
    {{ if .validationFailed }}
        <ul>
        {{range $index, $element := .validationErrors}}
        <li> Field <b>{{$index}}</b>
            <ul>
            {{range $index2, $element2 := $element}}
            <li> Validation for <b>{{$element2}}</b> failed!</li>
            {{end}}
            </ul>
        {{end}}
        </ul> 
    {{end}}
</div>