<div class="alert alert-dismissible alert-warning">
	<button type="button" class="close" data-dismiss="alert">×</button>
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
		<a href="#" onclick="window.history.back();" class="btn btn-info">Go back to form</a>
	</div>
