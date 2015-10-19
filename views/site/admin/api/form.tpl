<div class="panel panel-default">
	<div class="panel-heading" style="padding-bottom:0px">
		<ul class="breadcrumb">
			<li><a href="/admin">Admin</a></li>
			<li><a href="/admin/api">API</a></li>
			<li class="active">{{.action}}</li>
		</ul>
	</div>
	<div class="panel-body">
		<form method="POST" class="form-horizontal">
			{{ .xsrf_input }}
			<fieldset>
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
					<div class="form-group">
						<div class="col-lg-10 col-lg-offset-2">
							<input type="submit" id="institution-submit-button" class="btn btn-primary" value="Add" /> <a href="/admin/api" class="btn btn-default">Cancel</a> </div>
						</div>
					</fieldset>
				</form>
			</div>
		</div>
		<script>
			$(document).ready(function(){
					{{ if .currentTable }}
						$("#table-status").val("{{ .currentTable.Status }}");
					{{ end }}
					$("#table-name").focus();
			});
		</script>
