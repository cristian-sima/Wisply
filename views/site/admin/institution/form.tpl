<div class="panel panel-default">
	<div class="panel-heading" style="padding-bottom:0px">
		<ul class="breadcrumb">
			<li><a href="/admin">Admin</a></li>
			<li><a href="/admin/institutions">Institutions</a></li>
			<li class="active">{{.action}}</li>
		</ul>
	</div>
	<div class="panel-body">
		<form method="{{.actionType}}" class="form-horizontal"> {{ .xsrf_input }} {{ $safeDescription := .institutionDescription|html}}
			<fieldset>
				<input type="hidden" name="institution-wikiID" id="institution-wikiID" value="{{.wikiID}}" />
				<div class="form-group">
					<label for="institution-name" class="col-lg-2 control-label">Name</label>
					<div class="col-lg-10">
						<input type="text" value="{{.institution.Name }}" class="form-control" name="institution-name" id="institution-name" placeholder="Name" required pattern=".{3,255}" title="The name has 3 up to 255 characters!">
					</div>
				</div>
				<div class="form-group">
					<label for="institution-wikiURL" class="col-lg-2 control-label">Wiki URL</label>
					<div class="col-lg-10">
						<div class="input-group">
							<input type="url" class="form-control" value="{{.institution.WikiURL}}" name="institution-wikiURL" id="institution-wikiURL" placeholder="Wiki address" pattern=".{3,2083}" title="The wiki URL has 3 up to 2083 characters!" />
							<span class="input-group-btn">
								<button class="btn btn-default" id="button-institution-wikiURL" type="button">Get info from Wiki</button>
							</span>
						</div>
					</div>
				</div>
        {{ if eq .action "Add" }}
					<div class="form-group">
						<label for="institution-URL" class="col-lg-2 control-label">Website</label>
						<div class="col-lg-10">
							<input type="url" value="{{.institutionUrl}}" class="form-control" name="institution-URL" id="institution-URL" placeholder="URL address" required pattern=".{3,2083}" title="The URL has 3 up to 2083 characters!">
						</div>
					</div>
      	{{ end }}
			</fieldset>
				<fieldset>
					<legend>Profile</legend>
				<div id="wiki-source-div">
					<div class="form-group text-center" style="min-height:80px">
						<label for="institution-description" class="col-lg-2 control-label text-center" >
							  <div class="institution-profile">
	                <div class="insider" id="institution-logo">
										{{ if .institution }}
											{{ if eq .institution.LogoURL "" }}
												<span class="glyphicon glyphicon-education institution-logo-default"></span>
											{{ else }}
												<img src="{{ .institution.LogoURL }}" class="inlogo"  />
											{{ end }}
									 	{{ else }}
									 		<span class="glyphicon glyphicon-education institution-logo-default"></span>
									 	{{ end }}
									</div>
								</div>
	          </label>
						<div class="col-lg-10">
							<textarea class="form-control" rows="3" name="institution-description" id="institution-description" maxlength="1000">{{.institution.Description}}</textarea><span class="help-block">
	              <span class="description-modified hideMe">
									<span class="text-warning">
										<span class="glyphicon glyphicon-warning-sign"></span> Auto receiving from Wikipedia off
									</span>
									<span class="text-danger">
										<span data-toggle="tooltip" data-placement="top" title="Discard changes and activate it" class="discard-description-changes hover glyphicon glyphicon-remove">
										</span>
									</span>
								</span>
							<br /> This field may contain notes about the intitution.
						</div>
					</div>
          <div class="form-group">
						<label for="institution-logoURL" class="col-lg-2 control-label">Logo Address</label>
						<div class="col-lg-10">
							<div class="input-group">
								<input type="url" class="form-control" value="{{.institution.LogoURL}}" name="institution-logoURL" id="institution-logoURL" placeholder="Logo address" pattern=".{3,2083}" title="The wiki URL has 3 up to 2083 characters!">
								<span class="input-group-btn">
									<button disabled class="btn btn-default" id="button-get-wiki-by-address" type="button">Receive</button>
								</span>
							</div>
							<div class="text-center">
							<span class="description-modified hideMe">
								<span class="text-warning">
									<span class="glyphicon glyphicon-warning-sign"></span> Auto receiving from Wikipedia off
								</span>
								<span class="text-danger">
									<span data-toggle="tooltip" data-placement="top" title="Discard changes and activate it" class="discard-description-changes hover glyphicon glyphicon-remove">
									</span>
								</span>
							</span>
						</div>
						</div>
					</div>
				</div>
				<div class="form-group">
					<div class="col-lg-10 col-lg-offset-2">
						<input type="submit" id="institution-submit-button" class="btn btn-primary" value="Submit" /> <a href="/admin/institutions" class="btn btn-default">Cancel</a> </div>
				</div>
			</fieldset>
		</form>
	</div>
</div>
<script>
	var server = {};

	server.original = {
		description : "{{.institution.Description}}",
		wikiReceive : JSON.parse({{.wikiReceive}}),
		logoURL: {{ .institution.LogoURL }},
		{{ if .wikiID }}
		{{ if eq .wikiID "NULL" }}
			wikiID: "NULL",
		{{else }}
		wikiID: parseInt({{ .wikiID }}, 10),
		{{ end }}
		{{ else }}
		wikiID: undefined,
		{{ end }}
	};
</script>
<link href="/static/css/public/institution.css" type="text/css" rel="stylesheet" />
<script src="/static/3rd_party/others/js/jquery.elastic.source.js"></script>
<script src="/static/js/wisply/typer.js"></script>
<script src="/static/js/wisply/wikier.js"></script>
<script src="/static/js/admin/institution/functionality.js"></script>
