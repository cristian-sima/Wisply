<div class="alert alert-warning">
	<strong>Sorry...</strong> {{.messageContent}}
	{{ if .validationFailed }}
	<ul>
		{{range $index, $element := .validationErrors}}
		<li> Field <b>{{$index}}</b>
			<ul>
				{{range $index2, $element2 := $element}}
				<li>{{$element2}}</li>
				{{end}}
			</ul>
			{{end}}
		</ul>
		{{end}}
	</div>
	<div class="text-center" >
		<a href="javascript:window.location=document.referrer;" class="btn btn-info">Go back to form</a>
	</div>
