<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/curriculum">Curriculum</a></li>
      <li><a href="/admin/curriculum/programs/{{ .program.GetID }}">{{ .program.GetName }}</a></li>
      <li><a href="/admin/curriculum/programs/{{ .program.GetID }}/advance-options">Advance options</a></li>
      <li class="active">Modify static description</li>
    </ul>
  </div>
  <div class="panel-body">
    <form method="POST" class="form-horizontal">
      {{ .xsrf_input }}
      <div class="form-group">
        <label for="table-name" class="col-lg-2 control-label"></label>
        <div class="col-lg-10">
        Program <a target="_blabk" href="/curriculum/{{ .program.GetID }}"><strong>{{ .program.GetName }} </strong></a>
        </div>
      </div>
      <fieldset>
        <div class="form-group">
          <label for="program-description" class="col-lg-2 control-label">Description</label>
          <div class="col-lg-10">
            <textarea type="text" class="form-control" name="program-description" id="program-description" placeholder="Description" pattern=".{3,1000}" title="The name has 3 up to 1000 characters!">{{ .program.GetDescription }}</textarea>
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            <input type="submit" id="institution-submit-button" class="btn btn-primary" value="Submit" /> <a href="/admin/api" class="btn btn-default">Cancel</a> </div>
          </div>
        </fieldset>
      </form>
  </div>
</div>
<script src="/static/3rd_party/product/tinymce/js/tinymce/tinymce.min.js"></script>
<script>
$(document).ready(function(){
	tinymce.init({
		selector: "#program-description",
		auto_focus: "program-description",
    height: "300px",
	});
  $("#program-description").focus();
});
</script>
