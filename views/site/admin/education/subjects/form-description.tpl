<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li><a href="/admin/education/subjects/{{ .subject.GetID }}">{{ .subject.GetName }}</a></li>
      <li><a href="/admin/education/subjects/{{ .subject.GetID }}/advance-options">Advance options</a></li>
      <li class="active">Modify static description</li>
    </ul>
  </div>
  <div class="panel-body">
    <form method="POST" class="form-horizontal">
      {{ .xsrf_input }}
      <div class="form-group">
        <span class="col-lg-2 control-label"></span>
        <div class="col-lg-10">
          Subject <a target="_blank" href="/education/subjects/{{ .subject.GetID }}"><strong>{{ .subject.GetName }} </strong></a>
        </div>
      </div>
      <fieldset>
        <div class="form-group">
          <label for="subject-description" class="col-lg-2 control-label">Description</label>
          <div class="col-lg-10">
            <textarea class="form-control" name="subject-description" id="subject-description" placeholder="Description" title="The name has 3 up to 1000 characters!">{{ .subject.GetDescription }}</textarea>
          </div>
        </div>
        <div class="form-group">
          <div class="col-lg-10 col-lg-offset-2">
            <input type="submit" id="institution-submit-button" class="btn btn-primary" value="Submit" /> <a href="/admin/education/subjects/{{ .subject.GetID }}/advance-options" class="btn btn-default">Cancel</a> </div>
          </div>
        </fieldset>
      </form>
    </div>
  </div>
  <script src="/static/3rd_party/product/tinymce/js/tinymce/tinymce.min.js"></script>
  <script>
  $(document).ready(function(){
    tinymce.init({
      selector: "#subject-description",
      auto_focus: "subject-description",
      height: "300px",
    });
    $("#subject-description").focus();
  });
  </script>
