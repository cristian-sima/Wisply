<div class="panel panel-default">
	<div class="panel-heading" style="padding-bottom:0px">
		<ul class="breadcrumb">
			<li><a href="/admin">Admin</a></li>
			<li><a href="/admin/developer">API</a></li>
			<li class="active">{{.action}}</li>
		</ul>
	</div>
	<div class="panel-body">
		<form method="POST" class="form-horizontal">
			{{ .xsrf_input }}
			{{ if eq .type "Modify" }}
			<input type="hidden" value="{{ .currentTable.ID }}" name="table-id" id="table-id" />
			<div class="form-group">
				<span class="col-lg-2 control-label"></span>
				<div class="col-lg-10">
				Table <a href="/developer/table/list"><strong>{{ .currentTable.Name }} </strong></a>
				</div>
			</div>
			{{ end }}
			<fieldset>
				{{ if eq .type "Add" }}
				<div class="form-group">
					<label for="table-name" class="col-lg-2 control-label">Table name</label>
					<div class="col-lg-10">
						<select class="form-control" name="table-name" id="table-name">
							{{ $currentTable := .currentTable }}
							{{range $index, $table := .tables}}
							<option
							{{ if $currentTable }}
							{{ if eq $currentTable.Name $table }}
							selected
							{{ end }}
							{{ end }}
							value="{{ $table }}">{{ $table }}</option>
							{{ end }}
						</select>
					</div>
				</div>
				{{ end }}
				<div class="form-group">
					<label for="table-description" class="col-lg-2 control-label">Description</label>
					<div class="col-lg-10">
						<textarea class="form-control" name="table-description" id="table-description" placeholder="Description" title="The name has 3 up to 1000 characters!">{{ .currentTable.Description }}</textarea>
					</div>
				</div>
				<div class="form-group">
					<div class="col-lg-10 col-lg-offset-2">
						<input type="submit" id="institution-submit-button" class="btn btn-primary" value="Submit" /> <a href="/admin/developer" class="btn btn-default">Cancel</a> </div>
					</div>
				</fieldset>
			</form>
		</div>
	</div>
	<script src="/static/3rd_party/product/tinymce/js/tinymce/tinymce.min.js"></script>
	<script>
	tinymce.init({
		selector: "#table-description",
		auto_focus: "table-description",
	});
	</script>
	{{ if eq .type "Modify" }}
	<script>
		$(document).ready(function(){
			$("#table-description").focus();
		});
	</script>
	{{ else }}
	<script>
	$(document).ready(function(){
		$("#table-name").focus();
	});
	</script>
	{{ end }}
